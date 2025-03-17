import { drizzle } from 'drizzle-orm/postgres-js';
import postgres from 'postgres';
import * as schema from './schema';
import 'dotenv/config';

// Check if DATABASE_URL environment variable exists
if (!process.env.DATABASE_URL) {
  throw new Error('DATABASE_URL environment variable is not set');
}

// Create postgres client
const client = postgres(process.env.DATABASE_URL);

// Create drizzle database instance
export const db = drizzle(client, { schema });

// No need for separate initialization function with PostgreSQL
// as tables are typically managed through migrations
export const getDatabase = () => {
  return db;
};
