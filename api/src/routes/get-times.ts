import { getDatabase } from "../db/create-database";

export function getTimes(): Response {
  try {

    const db = getDatabase();

    // Query all records from the user_times table
    const records = db.query("SELECT * FROM user_times ORDER BY created_at DESC").all();
   
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