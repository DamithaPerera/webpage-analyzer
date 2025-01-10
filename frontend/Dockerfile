# Frontend Dockerfile
FROM node:20.18.1 AS builder

WORKDIR /app

# Copy package.json and package-lock.json
COPY frontend/package*.json ./

# Install dependencies
RUN npm install

# Copy the source code
COPY frontend/ .

# Build the frontend
RUN npm run build

# Serve the app using nginx
FROM nginx:alpine
COPY --from=builder /app/build /usr/share/nginx/html
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]
