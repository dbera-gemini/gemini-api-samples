import os
import requests
from dotenv import load_dotenv

load_dotenv()

def get_balances():
    """Fetch account balances."""
    url = f"{os.getenv('BASE_URL')}/balances"
    headers = {
        'X-API-Key': os.getenv('API_KEY'),
        'X-API-Secret': os.getenv('API_SECRET')
    }

    try:
        response = requests.get(url, headers=headers)
        response.raise_for_status()
        print('Account balances:', response.json())
    except requests.exceptions.RequestException as e:
        print(f'Error fetching balances: {e}')

if __name__ == '__main__':
    get_balances()
