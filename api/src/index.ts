import { handleUserTimeData } from "./routes/user-times";
import { getTimes } from "./routes/get-times";
import { serveSwaggerUI } from "./routes/swagger-ui";
import { handleDeleteTime } from "./routes/delete-time";
import { getDatabase } from "./db/create-database";

// Initialize the database when the server starts
console.log("Initializing database...");
const db = getDatabase();
console.log("Database initialized successfully!");

Bun.serve({
    port: 3000,
    async fetch(req) {
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
                return await handleUserTimeData(req);
            } else if (req.method === "GET") {
                return await getTimes();
            } else if (req.method === "DELETE") {
                return await handleDeleteTime(req);
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