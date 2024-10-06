# frontend/Dockerfile

# Use Node.js LTS version
FROM node:20

# Set working directory inside the container
WORKDIR /app

# Install Angular CLI globally
RUN npm install -g @angular/cli

## Ensure devDependencies are installed by setting NODE_ENV to development
ENV NODE_ENV=development

# Copy package.json and package-lock.json
COPY package*.json ./

# Install dependencies
RUN npm install

# Copy the rest of the frontend code
COPY src ./src
COPY angular.json tsconfig.json ./
COPY public ./public

# Expose the port Angular uses (default 4200)
EXPOSE 4200
