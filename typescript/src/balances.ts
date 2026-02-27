import axios from 'axios';
import * as dotenv from 'dotenv';

dotenv.config();

async function getBalances(): Promise<void> {
  try {
    const response = await axios.get(`${process.env.BASE_URL}/balances`, {
      headers: {
        'X-API-Key': process.env.API_KEY,
        'X-API-Secret': process.env.API_SECRET,
      },
    });

    console.log('Account balances:', response.data);
  } catch (error) {
    console.error('Error fetching balances:', error);
  }
}

// Example usage
getBalances();
