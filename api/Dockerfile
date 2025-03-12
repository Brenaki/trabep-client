# Use the official Bun image as base
FROM oven/bun:1.1.26-debian

# Set working directory
WORKDIR /app

# Copy package.json and other configuration files
COPY package.json bun.lockb tsconfig.json ./

# Install dependencies
RUN bun install

# Copy source code
COPY . .

# Expose the port the app runs on
EXPOSE 3000

# Build application
RUN bun run build

# Set environment variables
ENV NODE_ENV=production

# Run the application
CMD ["bun", "run", "start"]