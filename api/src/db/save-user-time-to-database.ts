import type { UserTimeData } from "../routes/user-time-data.dto";
import { getDatabase } from "./create-database";
import { userTimes } from "./schema";

// Function to save user time data to the database
export async function saveUserTimeToDatabase(
    userData: UserTimeData, 
    hoursSpent: number,
    minutesSpent: number, 
    secondsSpent: number
): Promise<boolean> {
    try {
        const db = getDatabase();
        
        // Insert data using Drizzle ORM
        await db.insert(userTimes).values({
            user: userData.user,
            startTime: userData.startTime,
            endTime: userData.endTime,
            hoursSpent: hoursSpent,
            minutesSpent: minutesSpent,
            secondsSpent: secondsSpent
        });
        
        console.log(`Saved user time data to database for user: ${userData.user}`);
        return true;
    } catch (error) {
        console.error("Error saving to database:", error);
        return false;
    }
}