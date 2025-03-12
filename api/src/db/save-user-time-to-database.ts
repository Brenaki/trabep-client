import type { UserTimeData } from "../routes/user-time-data.dto";
import { getDatabase } from "./create-database";

// Function to save user time data to the database
export function saveUserTimeToDatabase(
    userData: UserTimeData, 
    hoursSpent: number,
    minutesSpent: number, 
    secondsSpent: number
): boolean {
    try {
        const db = getDatabase();
        
        // Prepare and execute the insert statement
        const stmt = db.prepare(`
            INSERT INTO user_times (user, start_time, end_time, hours_spent, minutes_spent, seconds_spent)
            VALUES (?, ?, ?, ?, ?, ?)
        `);
        
        stmt.run(
            userData.user,
            userData.startTime,
            userData.endTime,
            hoursSpent,
            minutesSpent,
            secondsSpent
        );
        
        console.log(`Saved user time data to database for user: ${userData.user}`);
        return true;
    } catch (error) {
        console.error("Error saving to database:", error);
        return false;
    }
}