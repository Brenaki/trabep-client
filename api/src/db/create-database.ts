import { Database } from "bun:sqlite";
import { existsSync, mkdirSync } from "fs";
import { join, resolve } from "path";

// Database file path - using absolute path for clarity
const DB_DIR = resolve(join(import.meta.dir, "../../data"));
const DB_PATH = join(DB_DIR, "user_times.db");

/**
 * Initializes the SQLite database
 * Creates the database file and tables if they don't exist
 */
export function initializeDatabase(): Database {
    console.log(`Database directory path: ${DB_DIR}`);
    console.log(`Database file path: ${DB_PATH}`);
    
    // Create data directory if it doesn't exist
    if (!existsSync(DB_DIR)) {
        try {
            mkdirSync(DB_DIR, { recursive: true });
            console.log(`Created database directory: ${DB_DIR}`);
        } catch (err) {
            console.error(`Error creating database directory: ${err}`);
            throw new Error(`Failed to create database directory: ${err}`);
        }
    }

    // Create and connect to the database
    try {
        const db = new Database(DB_PATH);
        console.log(`Connected to database at: ${DB_PATH}`);

        // Create tables if they don't exist
        db.exec(`
            CREATE TABLE IF NOT EXISTS user_times (
                id INTEGER PRIMARY KEY AUTOINCREMENT,
                user TEXT NOT NULL,
                start_time TEXT NOT NULL,
                end_time TEXT NOT NULL,
                hours_spent INTEGER NOT NULL,
                minutes_spent INTEGER NOT NULL,
                seconds_spent INTEGER NOT NULL,
                created_at TEXT DEFAULT CURRENT_TIMESTAMP
            );
        `);
        console.log("Database tables created/verified");

        return db;
    } catch (err) {
        console.error(`Error connecting to database: ${err}`);
        throw new Error(`Failed to initialize database: ${err}`);
    }
}

// Export the database instance for use in other files
export const getDatabase = (): Database => {
    if (!existsSync(DB_PATH)) {
        console.warn("Database file not found, initializing...");
        return initializeDatabase();
    }
    return new Database(DB_PATH);
};
