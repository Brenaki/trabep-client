import { Database } from "bun:sqlite";
import { existsSync, mkdir } from "fs";
import { join } from "path";

// Database file path
const DB_PATH = join(import.meta.dir, "../../data/user_times.db");
const DB_DIR = join(import.meta.dir, "../../data");

/**
 * Initializes the SQLite database
 * Creates the database file and tables if they don't exist
 */
export function initializeDatabase(): Database {
    // Create data directory if it doesn't exist
    if (!existsSync(DB_DIR)) {
        mkdir(DB_DIR, { recursive: true }, (err) => {
            if (err) {
                console.error(`Error creating database directory: ${err}`);
                return;
            }
            console.log(`Created database directory: ${DB_DIR}`);
        });
    }

    // Create and connect to the database
    const db = new Database(DB_PATH);
    console.log(`Connected to database at: ${DB_PATH}`);

    // Create tables if they don't exist
    db.exec(`
        CREATE TABLE IF NOT EXISTS user_times (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            user TEXT NOT NULL,
            start_time TEXT NOT NULL,
            end_time TEXT NOT NULL,
            minutes_spent INTEGER NOT NULL,
            created_at TEXT DEFAULT CURRENT_TIMESTAMP
        );
    `);
    console.log("Database tables created/verified");

    return db;
}

// Export the database instance for use in other files
export const getDatabase = (): Database => {
    return new Database(DB_PATH);
};
