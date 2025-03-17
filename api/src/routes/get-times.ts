import { getDatabase } from "../db/create-database";
import { userTimes } from "../db/schema";
import { desc } from "drizzle-orm";

export async function getTimes(): Promise<Response> {
  try {
    const db = getDatabase();

    // Query all records from the user_times table using Drizzle ORM
    const records = await db.select().from(userTimes).orderBy(desc(userTimes.createdAt));
   
    // Return the records as JSON
    return new Response(
      JSON.stringify({
        success: true,
        count: records.length,
        data: records
      }),
      { 
        status: 200, 
        headers: { "Content-Type": "application/json" } 
      }
    );
  } catch (error) {
    console.error("Error fetching time records:", error);
    
    // Return error response
    return new Response(
      JSON.stringify({ 
        success: false, 
        error: "Failed to fetch time records" 
      }),
      { 
        status: 500, 
        headers: { "Content-Type": "application/json" } 
      }
    );
  }
}