import { records } from "../db/get-all-datas";

export function getTimes(): Response {
  try {
   
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