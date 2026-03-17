---
name: code-review
description: Multi-agent code review that validates API endpoints, catches placeholder errors, checks documentation, and does a general code review. Use when the user wants a thorough code review.
---

# Multi-Agent Code Review

Perform a comprehensive code review by spawning **five** sub-agents in parallel. Use the gathered context below to inform each agent.

## Gathered Context

- Git diff (staged + unstaged): !`git diff HEAD`
- Changed files: !`git diff HEAD --name-only`
- Untracked files: !`git ls-files --others --exclude-standard`

## Agents to Spawn

Launch agents 1-4 **in parallel** using the Agent tool. After all four complete, launch agent 5 (the validator) with their combined output.

### Agent 1: API Endpoint Validator

**subagent_type:** Explore

Prompt the agent with:

> You are reviewing a Gemini cryptocurrency exchange MCP server. Your job is to validate that every API endpoint used in the codebase matches the official Gemini REST API documentation at https://docs.gemini.com/rest-api/.
>
> Steps:
> 1. Read all files in `src/datasources/` and `src/client/http.ts` to find every API endpoint path and HTTP method used.
> 2. Read `src/tools/` to understand what parameters are being passed.
> 3. Fetch the Gemini API docs at https://docs.gemini.com/rest-api/ and cross-reference.
> 4. Check that URL paths, HTTP methods (GET vs POST), request payloads, and expected response shapes are correct.
> 5. Flag any endpoints that appear incorrect, deprecated, or missing required parameters.
>
> Return a structured report with: endpoint path, file location, status (correct/suspect/error), and notes.

### Agent 2: Placeholder & Hardcoded Value Scanner

**subagent_type:** Explore

Prompt the agent with:

> You are scanning this codebase for leftover placeholders, dummy values, and hardcoded secrets that should have been replaced with real values or configuration.
>
> Search for:
> - Strings like "your_", "xxx", "TODO", "FIXME", "HACK", "placeholder", "example", "test", "changeme", "replace", "dummy", "fake", "temp"
> - Hardcoded API keys, secrets, or tokens (long hex/base64 strings that aren't in .env or config)
> - Hardcoded URLs that should be configurable (especially localhost or sandbox URLs used as defaults in production code)
> - Default values that look like examples rather than sensible defaults
> - Template literal placeholders like ${} that are empty or contain placeholder text
> - Console.log statements that look like debug leftovers
> - Commented-out code blocks that should be removed or restored
>
> Check ALL files including config, source, and any scripts. Return each finding with: file, line number, the problematic string, and why it's suspect.

### Agent 3: Documentation Validator

**subagent_type:** Explore

Prompt the agent with:

> You are reviewing whether the documentation (README.md and any other docs) accurately reflects the current state of the codebase.
>
> Steps:
> 1. Read README.md thoroughly.
> 2. Read `src/tools/` files to get the full list of available MCP tools and their parameters.
> 3. Read `src/config.ts` for all environment variables and configuration options.
> 4. Read `src/index.ts` and `src/server.ts` for setup/startup behavior.
> 5. Cross-reference and check:
>    - Are all tools listed in the README? Are any missing or extra?
>    - Are environment variables documented correctly (names, defaults, required vs optional)?
>    - Are setup instructions accurate (install, build, run commands)?
>    - Are usage examples correct and using real tool names?
>    - Is any documented feature not actually implemented?
>    - Is any implemented feature not documented?
>
> Return a list of discrepancies with: what the docs say, what the code does, and suggested fix.

### Agent 4: General Code Review

**subagent_type:** general-purpose

Prompt the agent with:

> You are performing a general code review of a TypeScript MCP server for the Gemini cryptocurrency exchange API. This is a READ-ONLY review — do not modify any files.
>
> Review all source files in `src/` for:
>
> **Correctness & Logic:**
> - Error handling: are API errors caught and surfaced properly?
> - Edge cases: null/undefined handling, empty arrays, missing fields
> - Type safety: are TypeScript types used correctly? Any `any` types that should be specific?
>
> **Security:**
> - Are API keys/secrets handled safely (not logged, not exposed in errors)?
> - Is request signing implemented correctly?
> - Any injection risks in parameter handling?
>
> **Architecture & Patterns:**
> - Is the code organized consistently across modules (market, orders, funds, etc.)?
> - Are there duplicated patterns that could be consolidated?
> - Is error handling consistent?
>
> **Robustness:**
> - Are API responses validated before use?
> - Is there proper handling for rate limits or network errors?
> - Are numeric values (prices, amounts) handled safely (string vs number)?
>
> Return findings as: severity (critical/warning/info), file:line, description, and suggested fix.

### Agent 5: Validator & Final Report

**subagent_type:** general-purpose

After agents 1-4 complete, spawn this agent with all four reports concatenated into its prompt:

> You are the final validator for a multi-agent code review. You have received reports from four specialized reviewers. Your job is to:
>
> 1. **Deduplicate**: Remove findings that appear in multiple reports.
> 2. **Validate**: Remove false positives — findings that are incorrect or not actually problems. Use your judgement and read the relevant files if needed to verify.
> 3. **Prioritize**: Rank remaining findings by impact (critical > warning > info).
> 4. **Consolidate**: Group related findings together.
> 5. **Format**: Produce a clean, actionable final report.
>
> ## Input Reports
>
> [Insert the four agent reports here]
>
> ## Output Format
>
> Produce the final report in this format:
>
> ### Critical Issues
> - **[Category]** file:line — description and suggested fix
>
> ### Warnings
> - **[Category]** file:line — description and suggested fix
>
> ### Suggestions
> - **[Category]** file:line — description and suggested fix
>
> ### Summary
> Brief overall assessment of code quality and top 3 priorities.
>
> Only include findings you believe are genuinely valuable. Quality over quantity.
