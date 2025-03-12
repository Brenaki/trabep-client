import { getDatabase } from "./create-database";

 // Query all records from the user_times table
export const records = getDatabase().query("SELECT * FROM user_times ORDER BY created_at DESC").all();
    