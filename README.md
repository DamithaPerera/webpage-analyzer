# Webpage Analyzer

Webpage Analyzer is a project designed to analyze webpages by extracting information such as HTML version, titles, headings, and other metadata. This project includes both backend and frontend implementations running in Docker containers.

## Project Structure

```plaintext
.
|-- backend/                # Go backend service
|-- frontend/               # React frontend (served via Nginx)
|-- docker-compose.yml      # Docker Compose configuration
|-- README.md               # Project documentation
```

---

## Prerequisites

- Docker
- Docker Compose

---

## APIs

### Base URL

- **Backend:** `http://localhost:8080`
- **Frontend:** `http://localhost:3000`

### Endpoints

#### Home Page

- **Endpoint:** `GET /`
- **Description:** Displays a welcome message.
- **Response:**
  ```json
  {
      "message": "Welcome to the Web Page Analyzer. Use POST /analyze to analyze a webpage."
  }
  ```

#### Analyze Page

- **Endpoint:** `POST /analyze`
- **Description:** Analyzes the given webpage and returns its metadata.
- **Request Body:**
  ```json
  {
      "url": "http://example.com"
  }
  ```
- **Response:**
  ```json
  {
      "html_version": "HTML5",
      "title": "Example Domain",
      "headings": {
          "h1": 1
      },
      "internal_links": 2,
      "external_links": 5,
      "inaccessible_links": 1,
      "has_login_form": false
  }
  ```
#### Prometheus Metrics (Backend Only)

- **Endpoint:** `GET /metrics`
- **Description:** Exposes metrics in Prometheus format for monitoring and analytics.
- **Important:** This endpoint is **only accessible from the backend** and is not integrated into the frontend.

---
---

## Setup

### 1. Clone the Repository

```bash
git clone <repository-url>
cd webpage-analyzer
```

### 2. Build and Run the Application Using Docker

Run the following command to build and start the containers:

```bash
docker-compose up --build
```

### 3. Access the Application

- Frontend: Open your browser and navigate to `http://localhost:3000`
- Backend: Test APIs using Postman or `curl` at `http://localhost:8080`

---

## Testing

### Running Unit Tests

Navigate to the `backend/` directory and run:

```bash
cd backend
go test ./...
```

### Generating Test Coverage

Generate and display test coverage:

```bash
go test ./... -cover
```

---

## File Descriptions

### Backend Dockerfile

Located in `backend/Dockerfile`:

- Builds the Go application.
- Uses `golang:1.23.3` for building and `debian:bookworm-slim` for a lightweight runtime.

### Root `docker-compose.yml`

Configures the frontend and backend services.

```yaml
version: "3.9"
services:
  backend:
    build:
      context: .
      dockerfile: backend/Dockerfile
    ports:
      - "8080:8080"
    networks:
      app-network:
        aliases:
          - backend
    dns:
      - 8.8.8.8
      - 8.8.4.4

  frontend:
    build:
      context: .
      dockerfile: frontend/Dockerfile
    ports:
      - "3000:80"
    networks:
      app-network:
        aliases:
          - frontend
    dns:
      - 8.8.8.8
      - 8.8.4.4

networks:
  app-network:
    driver: bridge
```

---

## Troubleshooting

- **Backend API not accessible:** Ensure the backend container is running and listening on port `8080`.
- **Frontend not accessible:** Ensure the frontend container is running and Nginx is serving content on port `3000`.
- **Docker networking issues:** Try restarting Docker or clearing stale containers with `docker-compose down`.

---

## License

This project is licensed under the MIT License. See the `LICENSE` file for details.

---

## Authors

- **Damitha Perera**

