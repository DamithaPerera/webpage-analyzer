# Webpage Analyzer

Webpage Analyzer is a project designed to analyze webpages by extracting information such as HTML version, titles, headings, accessibility metadata, and other insights. This project includes both backend and frontend implementations running in Docker containers.

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

#### Health Check

- **Endpoint:** `GET /health`
- **Description:** Provides system health information and application metadata.
- **Response:**
  ```json
  {
      "status": "OK",
      "version": "Web Page Analyzer service is running."
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
      "has_login_form": false,
      "missing_labels": 0,
      "invalid_href": 0
  }
  ```

#### Prometheus Metrics (Backend Only)

- **Endpoint:** `GET /metrics`
- **Description:** Exposes metrics in Prometheus format for monitoring and analytics.
- **Important:** This endpoint is **only accessible from the backend** and is not integrated into the frontend.

---

## Features

### Error Handling

- **503 Service Unavailable:** The system returns a proper error response when a service is temporarily unavailable.
- **504 Gateway Timeout:** The system gracefully handles and logs timeout errors when the backend doesn’t receive a timely response from an external service.

### Accessibility Checks

- Validates hyperlinks for accessibility by:
  - Checking for missing labels (e.g., `<a>` elements without `aria-label` or text content).
  - Detecting invalid or inaccessible `href` attributes.

### Concurrency with Go Routines

- Go routines are used in the backend to handle multiple tasks concurrently, improving performance for webpage analysis.

### Logging

- Logs all significant events (e.g., API access, errors, and successes) to a file named `webpage-analyzer.log`. This file is dynamically created and ignored in version control (Git).
- Logs are formatted in JSON for easier parsing and integration with monitoring tools.

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

## Challenges Faced and Solutions

### 1. **Error Handling**
   - **Challenge:** Properly handling 503 and 504 errors for better user feedback.
   - **Solution:** Enhanced the backend to detect and return appropriate HTTP status codes with descriptive error messages.

### 2. **Concurrency with Go Routines**
   - **Challenge:** Optimizing the performance of webpage analysis tasks.
   - **Solution:** Used Go routines to process tasks (e.g., analyzing links, extracting headings) concurrently, reducing response time.

### 3. **Logging System**
   - **Challenge:** Centralizing logs and ensuring they are easily traceable.
   - **Solution:** Implemented JSON-formatted logs with file-based storage, ignored by Git.

### 4. **Accessibility Checks**
   - **Challenge:** Ensuring hyperlinks are accessible.
   - **Solution:** Added validation for missing labels and invalid `href` attributes during webpage analysis.

---

## Possible Improvements

1. **Enhanced Health Check API:**
   - Include real-time dependency status (e.g., database, external APIs).
   - Add dynamic metadata injection for version, build time, and Git commit.

2. **Caching Mechanism:**
   - Use a caching layer (e.g., Redis) to store results of frequently analyzed pages.

3. **Monitoring and Alerts:**
   - Integrate with monitoring tools like Prometheus and Grafana for visualizing metrics and setting up alerts.

4. **Frontend Enhancements:**
   - Add a dashboard for visualizing analysis results and logs in real time.

5. **Improved Accessibility:**
   - Extend the accessibility checks to include ARIA roles, color contrast checks, and keyboard navigation validation.

---

## Authors

- **Damitha Perera**

