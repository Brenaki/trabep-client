import { migrate } from 'drizzle-orm/postgres-js/migrator';
import { db } from './create-database';

// This will run migrations on the database, creating tables if they don't exist
async function main() {
  console.log('Running migrations...');
  
  await migrate(db, { migrationsFolder: './drizzle' });
  
  console.log('Migrations completed successfully');
  process.exit(0);
}

main().catch((err) => {
  console.error('Migration failed:', err);
  process.exit(1);
});