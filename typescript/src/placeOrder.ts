import axios from 'axios';
import * as dotenv from 'dotenv';

dotenv.config();

interface OrderParams {
  symbol: string;
  side: 'buy' | 'sell';
  type: 'market' | 'limit';
  quantity: number;
  price?: number;
}

async function placeOrder(params: OrderParams): Promise<void> {
  try {
    const response = await axios.post(
      `${process.env.BASE_URL}/order`,
      params,
      {
        headers: {
          'X-API-Key': process.env.API_KEY,
          'X-API-Secret': process.env.API_SECRET,
          'Content-Type': 'application/json',
        },
      }
    );

    console.log('Order placed:', response.data);
  } catch (error) {
    console.error('Error placing order:', error);
  }
}

// Example usage
placeOrder({
  symbol: 'BTCUSD',
  side: 'buy',
  type: 'limit',
  quantity: 0.01,
  price: 50000,
});
