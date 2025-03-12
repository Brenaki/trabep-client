/**
 * Calculates the time difference between two timestamps
 * @param startTime - Date string in format "DD/MM/YYYY, HH:MM:SS"
 * @param endTime - Date string in format "DD/MM/YYYY, HH:MM:SS"
 * @returns Object containing hours, minutes, seconds, and formatted time
 */
export function calculateTimeDifference(startTime: string, endTime: string): { 
    hours: number,
    minutes: number, 
    seconds: number,
    formatted: string 
} {
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
    
    // Convert to total seconds
    const totalSeconds = Math.floor(diffMs / 1000);
    
    // Calculate hours, minutes and remaining seconds
    const hours = Math.floor(totalSeconds / 3600);
    const minutes = Math.floor((totalSeconds % 3600) / 60);
    const seconds = totalSeconds % 60;
    
    // Format as hours, minutes and seconds
    const formatted = `${hours}h ${minutes}m ${seconds}s`;
    
    return { hours, minutes, seconds, formatted };
}