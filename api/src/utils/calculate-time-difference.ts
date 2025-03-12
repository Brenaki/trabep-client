/**
 * Calculates the time difference between two timestamps
 * @param startTime - Date string in format "DD/MM/YYYY, HH:MM:SS"
 * @param endTime - Date string in format "DD/MM/YYYY, HH:MM:SS"
 * @returns Object containing minutes (total minutes) and formatted (human-readable format)
 */
export function calculateTimeDifference(startTime: string, endTime: string): { minutes: number, formatted: string } {
    // Parse date strings to Date objects
    const parseDate = (dateStr: string): Date => {
        // Split the date string into date and time parts
        const [datePart, timePart] = dateStr.split(', ');
        
        // Split date part into day, month, year
        const [day, month, year] = datePart.split('/');
        
        // Create a date object (month is 0-indexed in JavaScript Date)
        return new Date(`${year}-${month}-${day}T${timePart}`);
    };
    
    // Parse the dates
    const start = parseDate(startTime);
    const end = parseDate(endTime);
    
    // Calculate difference in milliseconds
    const diffMs = end.getTime() - start.getTime();
    
    // Convert to minutes
    const minutes = Math.floor(diffMs / 60000);
    
    // Format as hours and minutes
    const hours = Math.floor(minutes / 60);
    const remainingMinutes = minutes % 60;
    const formatted = `${hours}h ${remainingMinutes}m`;
    
    return { minutes, formatted };
}