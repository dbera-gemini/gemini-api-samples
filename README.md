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
   # Edit .env and add your API_KEY and API_SECRET
   ```

3. **Run Examples:**

   **TypeScript:**
   ```bash
   cd typescript
   npm install
   npx ts-node src/placeOrder.ts
   npx ts-node src/balances.ts
   ```

   **Python:**
   ```bash
   cd python
   pip install -r requirements.txt
   python3 place_order.py
   python3 balances.py
   ```

   **Go:**
   ```bash
   cd go
   go mod download
   go run place_order.go
   go run balances.go
   ```

## API Documentation

- [Gemini API Documentation](https://docs.gemini.com/rest-api/)
- [Public Ticker API](https://docs.gemini.com/rest-api/#ticker)
- [Private Endpoints](https://docs.gemini.com/rest-api/#authenticated-api-invocation)

## Example Output

```bash
$ go run get_ticker.go
Ticker data: map[ask:65448.39 bid:65437.93 last:65422.70 volume:map[BTC:300.93753108 USD:19688145.814587516 timestamp:1.772231493e+12]]
```

## Notes

- Public endpoints (like `getTicker`) can be called without authentication
- Private endpoints require valid API credentials
- Use the sandbox environment for testing to avoid real trades
- All examples use lowercase symbols (e.g., `btcusd` not `BTCUSD`)