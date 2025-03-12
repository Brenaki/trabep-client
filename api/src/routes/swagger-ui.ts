import { readFileSync } from 'fs';
import { join } from 'path';
import { swaggerSpec } from '../swagger/swagger';

export function serveSwaggerUI(req: Request): Response {
  const url = new URL(req.url);
  const path = url.pathname.replace('/api-docs', '');

  // Serve the swagger specification as JSON
  if (path === '/swagger.json') {
    return new Response(JSON.stringify(swaggerSpec), {
      headers: { 'Content-Type': 'application/json' }
    });
  }
  
  // Serve the main HTML page with embedded Swagger UI
  const html = `
<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <title>Swagger UI</title>
  <link rel="stylesheet" type="text/css" href="https://cdn.jsdelivr.net/npm/swagger-ui-dist@5/swagger-ui.css" />
  <style>
    html { box-sizing: border-box; overflow: -moz-scrollbars-vertical; overflow-y: scroll; }
    *, *:before, *:after { box-sizing: inherit; }
    body { margin: 0; background: #fafafa; }
  </style>
</head>
<body>
  <div id="swagger-ui"></div>

  <script src="https://cdn.jsdelivr.net/npm/swagger-ui-dist@5/swagger-ui-bundle.js"></script>
  <script>
    window.onload = function() {
      window.ui = SwaggerUIBundle({
        url: "/api-docs/swagger.json",
        dom_id: '#swagger-ui',
        deepLinking: true,
        presets: [
          SwaggerUIBundle.presets.apis,
          SwaggerUIBundle.SwaggerUIStandalonePreset
        ],
        layout: "BaseLayout",
        supportedSubmitMethods: ['get', 'post', 'put', 'delete', 'patch'],
      });
    };
  </script>
</body>
</html>
  `;
  
  return new Response(html, {
    headers: { 'Content-Type': 'text/html' }
  });
}