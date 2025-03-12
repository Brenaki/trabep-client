import { getDatabase } from "../db/create-database";

/**
 * Deletes a time record by ID
 * @param id The ID of the record to delete
 * @returns Response with the result of the operation
 */
export function deleteTimeById(id: number): Response {
  try {
    const db = getDatabase();
    
    // Check if the record exists
    const record = db.query("SELECT id FROM user_times WHERE id = ?").get(id);
    
    if (!record) {
      return new Response(
        JSON.stringify({
          success: false,
          error: "Record not found"
        }),
        {
          status: 404,
          headers: { "Content-Type": "application/json" }
        }
      );
    }
    
    // Delete the record
    const stmt = db.prepare("DELETE FROM user_times WHERE id = ?");
    stmt.run(id);
    
    return new Response(
      JSON.stringify({
        success: true,
        message: `Record with ID ${id} deleted successfully`
      }),
      {
        status: 200,
        headers: { "Content-Type": "application/json" }
      }
    );
  } catch (error) {
    console.error("Error deleting time record:", error);
    
    return new Response(
      JSON.stringify({
        success: false,
        error: "Failed to delete time record"
      }),
      {
        status: 500,
        headers: { "Content-Type": "application/json" }
      }
    );
  }
}

/**
 * Handler for delete requests
 * @param req The request object
 * @returns Response with the result of the operation
 */
export function handleDeleteTime(req: Request): Response {
  const url = new URL(req.url);
  const idParam = url.searchParams.get("id");
  
  if (!idParam) {
    return new Response(
      JSON.stringify({
        success: false,
        error: "Missing required parameter: id"
      }),
      {
        status: 400,
        headers: { "Content-Type": "application/json" }
      }
    );
  }
  
  const id = parseInt(idParam);
  
  if (isNaN(id)) {
    return new Response(
      JSON.stringify({
        success: false,
        error: "Invalid ID format"
      }),
      {
        status: 400,
        headers: { "Content-Type": "application/json" }
      }
    );
  }
  
  return deleteTimeById(id);
}