# Gemini API Samples

Multi-language examples demonstrating how to use the Gemini cryptocurrency exchange API.

## Project Structure

```
├── .env.example          # Environment variables template
├── typescript/           # TypeScript examples
│   ├── src/
│   │   ├── getTicker.ts  # Get ticker data (public endpoint)
│   │   ├── placeOrder.ts # Place an order (requires auth)
│   │   └── balances.ts   # Get account balances (requires auth)
├── python/               # Python examples
│   ├── get_ticker.py     # Get ticker data (public endpoint)
│   ├── place_order.py    # Place an order (requires auth)
│   └── balances.py       # Get account balances (requires auth)
└── go/                   # Go examples
    ├── get_ticker.go     # Get ticker data (public endpoint)
    ├── place_order.go    # Place an order (requires auth)
    └── balances.go       # Get account balances (requires auth)
```

## Quick Start

### Get Ticker (No API Key Required)
```bash
# Python
python3 python/get_ticker.py btcusd

# TypeScript
npx ts-node typescript/src/getTicker.ts ethusd

# Go
go run go/get_ticker.go dogeusd
```

### Get Balances (API Key Required)

1. **Add your API credentials to `.env`:**
   ```bash
   cp .env.example .env
   # Edit .env and add your real GEMINI_API_KEY and GEMINI_API_SECRET
   ```

2. **Run the balances script:**
   ```bash
   # Python
   cd python && pip install -r requirements.txt
   python3 balances.py

   # TypeScript
   cd typescript && npm install
   npx ts-node src/balances.ts

   # Go
   cd go && go mod download
   go run balances.go
   ```

## Getting Started

### Configuration

The base URL is configured via environment variables. Copy the example file:

```bash
cp .env.example .env
```

The default configuration uses the production Gemini API. For testing, you can change `GEMINI_BASE_URL` to the sandbox environment:

```bash
GEMINI_BASE_URL=https://api.sandbox.gemini.com/v1
```

### Public Endpoints (No Authentication Required)

The `getTicker` examples use public API endpoints and don't require any API keys.

#### TypeScript

```bash
cd typescript
npm install
# Default symbol (btcusd)
npx ts-node src/getTicker.ts

# Custom symbol
npx ts-node src/getTicker.ts ethusd

# Show help
npx ts-node src/getTicker.ts --help
```

#### Python

```bash
cd python
pip install requests
# Default symbol (btcusd)
python3 get_ticker.py

# Custom symbol
python3 get_ticker.py ethusd

# Show help
python3 get_ticker.py --help
```

#### Go

```bash
cd go
# Default symbol (btcusd)
go run get_ticker.go

# Custom symbol
go run get_ticker.go ethusd

# Show help
go run get_ticker.go --help
```

### Private Endpoints (Authentication Required)

The `placeOrder` and `balances` examples require Gemini API credentials.

1. **Get API Credentials:**
   - Sign up at [Gemini](https://www.gemini.com/)
   - Create API keys in your account settings
   - Use the sandbox environment for testing: https://exchange.sandbox.gemini.com/

2. **Configure Environment:**
   ```bash
   cp .env.example .env
   # Edit .env and add your GEMINI_API_KEY and GEMINI_API_SECRET
   ```

   Your `.env` file should look like:
   ```
   GEMINI_BASE_URL=https://api.gemini.com/v1
   GEMINI_API_KEY=account-xxxxxxxxxxxxxx
   GEMINI_API_SECRET=xxxxxxxxxxxxxx
   ```

   **For testing, use the sandbox:**
   ```
   GEMINI_BASE_URL=https://api.sandbox.gemini.com/v1
   GEMINI_API_KEY=your-sandbox-api-key
   GEMINI_API_SECRET=your-sandbox-api-secret
   ```

3. **Run Examples:**

   **Get Balances:**

   TypeScript:
   ```bash
   cd typescript
   npm install
   npx ts-node src/balances.ts
   ```

   Python:
   ```bash
   cd python
   pip install -r requirements.txt
   python3 balances.py
   ```

   Go:
   ```bash
   cd go
   go mod download
   go run balances.go
   ```

   **Example Output:**
   ```json
   Account balances:
   [
     {
       "type": "exchange",
       "currency": "BTC",
       "amount": "5.0",
       "available": "4.5",
       "availableForWithdrawal": "4.5",
       "pendingWithdrawal": "0.25",
       "pendingDeposit": "0.25"
     },
     {
       "type": "exchange",
       "currency": "USD",
       "amount": "15000.00",
       "available": "5000.00",
       "availableForWithdrawal": "5000.00"
     },
     {
       "type": "exchange",
       "currency": "ETH",
       "amount": "10.0",
       "available": "10.0",
       "availableForWithdrawal": "10.0"
     }
   ]

   Total currencies: 3
   BTC: 4.5 available (5.0 total)
   USD: 5000.00 available (15000.00 total)
   ETH: 10.0 available (10.0 total)
   ```

   **Response Fields:**
   - `type`: Account type (e.g., "exchange")
   - `currency`: Currency code (BTC, ETH, USD, etc.)
   - `amount`: Total balance including pending
   - `available`: Amount available for trading
   - `availableForWithdrawal`: Amount that can be withdrawn
   - `pendingWithdrawal`: Amount in pending withdrawals (optional)
   - `pendingDeposit`: Amount in pending deposits (optional)

   **Place Orders:**
   ```bash
   # Similar to balances, run the respective place_order files
   npx ts-node src/placeOrder.ts
   python3 place_order.py
   go run place_order.go
   ```

## API Documentation

- [Gemini API Documentation](https://docs.gemini.com/rest-api/)
- [Public Ticker API](https://docs.gemini.com/rest-api/#ticker)
- [Private Endpoints](https://docs.gemini.com/rest-api/#authenticated-api-invocation)

## Example Outputs

### Get Ticker (Public)

```bash
$ go run get_ticker.go btcusd
Ticker data: map[ask:65448.39 bid:65437.93 last:65422.70 volume:map[BTC:300.93753108 USD:19688145.814587516 timestamp:1.772231493e+12]]
```

### Get Balances (Private)

```bash
$ python3 balances.py
Account balances:
[
  {
    "type": "exchange",
    "currency": "BTC",
    "amount": "5.0",
    "available": "4.5",
    "availableForWithdrawal": "4.5"
  }
]

Total currencies: 3
BTC: 4.5 available (5.0 total)
USD: 5000.00 available (15000.00 total)
ETH: 10.0 available (10.0 total)
```

## Authentication

The Gemini API uses a specific authentication method for private endpoints:

1. **Payload**: Create a JSON object with `request` path and `nonce` (current timestamp in milliseconds)
2. **Base64 Encode**: Encode the payload to base64
3. **HMAC Signature**: Create an HMAC-SHA384 signature of the base64 payload using your API secret
4. **Headers**: Send the request with:
   - `X-GEMINI-APIKEY`: Your API key
   - `X-GEMINI-PAYLOAD`: Base64 encoded payload
   - `X-GEMINI-SIGNATURE`: Hex-encoded HMAC signature

All implementations handle this authentication automatically.

## Notes

- Public endpoints (like `getTicker`) can be called without authentication
- Private endpoints require valid API credentials
- Use the sandbox environment for testing to avoid real trades
- All examples use lowercase symbols (e.g., `btcusd` not `BTCUSD`)
- The balances endpoint returns all your account balances across different currencies