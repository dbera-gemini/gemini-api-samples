import axios from 'axios';
import * as dotenv from 'dotenv';

dotenv.config();

async function getTicker(symbol: string): Promise<void> {
  try {
    const response = await axios.get(`${process.env.BASE_URL}/ticker`, {
      params: { symbol },
      headers: {
        'X-API-Key': process.env.API_KEY,
      },
    });

    console.log('Ticker data:', response.data);
  } catch (error) {
    console.error('Error fetching ticker:', error);
  }
}

// Example usage
getTicker('BTCUSD');
