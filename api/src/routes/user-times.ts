import { calculateTimeDifference } from "../utils/calculate-time-difference";
import { saveUserTimeToDatabase } from "../db/save-user-time-to-database";
import { UserTimeSchema, type UserTimeData } from "./user-time-data.dto";



// Function to handle the user time data
export async function handleUserTimeData(req: Request) {
    try {
        const data = await req.json();
        
        // Validate the received data using Zod
        const result = UserTimeSchema.safeParse(data);
        
        if (!result.success) {
            return new Response(
                JSON.stringify({ 
                    error: "Validation failed", 
                    issues: result.error.issues 
                }),
                { status: 400, headers: { "Content-Type": "application/json" } }
            );
        }
        
        const validatedData: UserTimeData = result.data;
        
        // Calculate time difference
        const timeDifference = calculateTimeDifference(
            validatedData.startTime, 
            validatedData.endTime
        );
         
        // Save the data to the database
        const savedToDb = saveUserTimeToDatabase(validatedData, timeDifference.minutes);
        
        // Return success response with time calculation
        return new Response(
            JSON.stringify({ 
                success: true, 
                message: "Data received successfully",
                data: validatedData,
                timeSpent: {
                    minutes: timeDifference.minutes,
                    formatted: timeDifference.formatted
                },
                savedToDatabase: savedToDb
            }),
            { 
                status: 200, 
                headers: { "Content-Type": "application/json" } 
            }
        );
    } catch (error) {
        console.error("Error processing request:", error);
        return new Response(
            JSON.stringify({ error: "Invalid JSON data" }),
            { status: 400, headers: { "Content-Type": "application/json" } }
        );
    }
}