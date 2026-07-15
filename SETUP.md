# Setup Guide

This guide explains how to run the **Geofencing & Real-Time Alert System** locally using Go, React, PostgreSQL (PostGIS), and Docker.

---

# Prerequisites

Before running the project, install the following software:

- Go 1.24 or later
- Node.js 20 or later
- Docker
- Docker Compose
- Git

Verify the installations:

```bash
go version
node -v
docker --version
docker compose version
```

---

# Clone the Repository

```bash
git clone https://github.com/MohamedFazil1406/RealTimeAlertSystem.git

cd RealTimeAlertSystem
```

---

# Database Setup

Start PostgreSQL and PostGIS using Docker:

```bash
docker compose up -d
```

Verify that the containers are running:

```bash
docker ps
```

---

# Backend Setup

Navigate to the backend directory:

```bash
cd backend
```

Install Go dependencies:

```bash
go mod tidy
```

Create a `.env` file inside the `backend` directory.

Example:

```env
PORT=8080

DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=geofence
DB_SSLMODE=disable
```

Start the backend server:

```bash
go run ./cmd
```

The backend will be available at:

```text
http://localhost:8080
```

---

# Frontend Setup

Open a new terminal.

Navigate to the frontend directory:

```bash
cd frontend
```

Install dependencies:

```bash
npm install
```

Start the development server:

```bash
npm run dev
```

The frontend will be available at:

```text
http://localhost:5173
```

---

# Run with Docker

To start the complete application:

```bash
docker compose up --build
```

Run in detached mode:

```bash
docker compose up -d --build
```

Stop the application:

```bash
docker compose down
```

View container logs:

```bash
docker compose logs -f
```

---

# Application URLs

| Service     | URL                           |
| ----------- | ----------------------------- |
| Frontend    | http://localhost:5173         |
| Backend API | http://localhost:8080         |
| WebSocket   | ws://localhost:8080/ws/alerts |

---

# API Testing

## Create Vehicle

```bash
curl -X POST http://localhost:8080/vehicles \
-H "Content-Type: application/json" \
-d '{
  "vehicle_number":"TN01AB1234",
  "driver_name":"John Doe",
  "vehicle_type":"Car",
  "phone":"9999999999"
}'
```

---

## Get Vehicles

```bash
curl http://localhost:8080/vehicles
```

---

## Create Geofence

```bash
curl -X POST http://localhost:8080/geofences \
-H "Content-Type: application/json" \
-d '{
  "name":"Office Zone",
  "description":"Main Office Area",
  "category":"restricted_zone",
  "coordinates":[
    {
      "latitude":12.9716,
      "longitude":77.5946
    },
    {
      "latitude":12.9720,
      "longitude":77.5956
    },
    {
      "latitude":12.9730,
      "longitude":77.5947
    },
    {
      "latitude":12.9716,
      "longitude":77.5946
    }
  ]
}'
```

---

## Update Vehicle Location

Replace `<VEHICLE_ID>` with a valid vehicle UUID.

```bash
curl -X POST http://localhost:8080/vehicles/location \
-H "Content-Type: application/json" \
-d '{
  "vehicle_id":"<VEHICLE_ID>",
  "latitude":12.9716,
  "longitude":77.5946,
  "timestamp":"2026-07-15T10:30:00Z"
}'
```

---

## Get Vehicle Location

```bash
curl http://localhost:8080/vehicles/location/<VEHICLE_ID>
```

---

## Get Dashboard Statistics

```bash
curl http://localhost:8080/dashboard
```

---

## Get Violation History

```bash
curl http://localhost:8080/violations/history
```

---

# Docker Commands

Start containers:

```bash
docker compose up -d
```

Stop containers:

```bash
docker compose down
```

Restart containers:

```bash
docker compose restart
```

Rebuild containers:

```bash
docker compose up --build
```

View running containers:

```bash
docker ps
```

---

# Notes

- PostgreSQL data is persisted using Docker volumes.
- The frontend communicates with the backend using Axios.
- Real-time alerts are delivered through WebSockets.
- Geofence detection is implemented using a Point-in-Polygon algorithm.
- Vehicle markers are displayed on an interactive Leaflet map.
- Dashboard statistics update based on backend data.
- Ensure Docker is running before starting the backend if PostgreSQL is containerized.

---

# Troubleshooting

## Backend cannot connect to PostgreSQL

Verify that PostgreSQL is running:

```bash
docker ps
```

Check backend environment variables in the `.env` file.

---

## Frontend cannot reach backend

Verify that the backend is running on:

```text
http://localhost:8080
```

Check your frontend API base URL configuration.

---

## Vehicle markers do not appear

- Verify that `GET /vehicles` returns `latitude` and `longitude`.
- Ensure at least one location has been created using `POST /vehicles/location`.
- Refresh the frontend after updating vehicle locations if WebSocket updates are not enabled.

---

## WebSocket connection fails

Verify that the backend is running and the WebSocket endpoint is available:

```text
ws://localhost:8080/ws/alerts
```
