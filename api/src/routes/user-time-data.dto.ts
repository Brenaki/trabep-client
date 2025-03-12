import { z } from "zod";

// Define the schema for validation
export const UserTimeSchema = z.object({
    user: z.string().min(1, "User name is required"),
    startTime: z.string().min(1, "Start time is required"),
    endTime: z.string().min(1, "End time is required")
});

// Type inference from the schema
export type UserTimeData = z.infer<typeof UserTimeSchema>;

