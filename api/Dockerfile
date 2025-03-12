# Use the official Bun image as base
FROM oven/bun:1.0 as base

# Set working directory
WORKDIR /app

# Copy package.json and other configuration files
COPY package.json bun.lockb tsconfig.json ./

# Install dependencies
RUN bun install --frozen-lockfile

# Copy source code
COPY src/ ./src/

# Create data directory for SQLite database
RUN mkdir -p data

# Expose the port the app runs on
EXPOSE 3000

# Set environment variables
ENV NODE_ENV=production

# Run the application
CMD ["bun", "run", "start"]