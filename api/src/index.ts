import { handleUserTimeData } from "./routes/user-times";
import { initializeDatabase } from "./db/create-database";
import { getTimes } from "./routes/get-times";
import { serveSwaggerUI } from "./routes/swagger-ui";
import { handleDeleteTime } from "./routes/delete-time";

// Initialize the database when the server starts
console.log("Initializing database...");
const db = initializeDatabase();
console.log("Database initialized successfully!");

Bun.serve({
    port: 3000,
    fetch(req) {
        const url = new URL(req.url);
        const pathname = url.pathname;
        
        // Serve Swagger UI documentation
        if (pathname === "/api-docs" || pathname.startsWith("/api-docs/")) {
            return serveSwaggerUI(req);
        }
        
        // Handle root path
        if (pathname === "/") {
            return new Response("Bun API Server - Visit /api-docs for documentation");
        }
        
        // Handle the user data routes
        if (pathname === "/user-times") {
            if (req.method === "POST") {
                return handleUserTimeData(req);
            } else if (req.method === "GET") {
                return getTimes();
            } else if (req.method === "DELETE") {
                return handleDeleteTime(req);
            } else {
                return new Response("Method Not Allowed", { status: 405 });
            }
        }

        // Default response for any other path
        return new Response("Not Found", { status: 404 });
    },
});

console.log("Server running at http://localhost:3000");
console.log("API documentation available at http://localhost:3000/api-docs");