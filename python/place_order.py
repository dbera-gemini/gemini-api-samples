import os
import requests
from dotenv import load_dotenv
from typing import Literal

load_dotenv()

def place_order(
    symbol: str,
    side: Literal['buy', 'sell'],
    order_type: Literal['market', 'limit'],
    quantity: float,
    price: float = None
):
    """Place an order."""
    url = f"{os.getenv('BASE_URL')}/order"
    headers = {
        'X-API-Key': os.getenv('API_KEY'),
        'X-API-Secret': os.getenv('API_SECRET'),
        'Content-Type': 'application/json'
    }
    data = {
        'symbol': symbol,
        'side': side,
        'type': order_type,
        'quantity': quantity
    }

    if price:
        data['price'] = price

    try:
        response = requests.post(url, headers=headers, json=data)
        response.raise_for_status()
        print('Order placed:', response.json())
    except requests.exceptions.RequestException as e:
        print(f'Error placing order: {e}')

if __name__ == '__main__':
    place_order('BTCUSD', 'buy', 'limit', 0.01, 50000)
