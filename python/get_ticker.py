import os
import requests
from dotenv import load_dotenv

load_dotenv()

def get_ticker(symbol: str):
    """Fetch ticker data for a given symbol."""
    url = f"{os.getenv('BASE_URL')}/ticker"
    headers = {
        'X-API-Key': os.getenv('API_KEY')
    }
    params = {'symbol': symbol}

    try:
        response = requests.get(url, headers=headers, params=params)
        response.raise_for_status()
        print('Ticker data:', response.json())
    except requests.exceptions.RequestException as e:
        print(f'Error fetching ticker: {e}')

if __name__ == '__main__':
    get_ticker('BTCUSD')
