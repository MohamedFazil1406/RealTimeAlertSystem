# Setup Guide

## Prerequisites

Install

- Go 1.24+
- Node.js 20+
- Docker
- Docker Compose

---

# Clone Repository

```bash
git clone <repository-url>

cd RealTimeAlertSystem
```

---

# Backend Setup

```bash
cd backend

go mod tidy
```

Create

```
.env
```

Example

```env
PORT=8080

DB_HOST=localhost

DB_PORT=5432

DB_USER=postgres

DB_PASSWORD=postgres

DB_NAME=geofence

DB_SSLMODE=disable
```

Run

```bash
go run ./cmd
```

---

# Frontend Setup

```bash
cd frontend

npm install

npm run dev
```

Runs on

```
http://localhost:5173
```

---

# Docker Setup

From project root

```bash
docker compose up --build
```

---

# Backend URL

```
http://localhost:8080
```

---

# Frontend URL

```
http://localhost:5173
```

---

# WebSocket

```
ws://localhost:8080/ws/alerts
```

---

# API Testing

## Create Vehicle

```bash
curl -X POST http://localhost:8080/vehicles \
-H "Content-Type: application/json" \
-d '{
"vehicle_number":"TN01AB1234",
"driver_name":"John",
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
"name":"Office",
"description":"Main Office",
"category":"restricted_zone",
"coordinates":[
{"latitude":12.9716,"longitude":77.5946},
{"latitude":12.9720,"longitude":77.5956},
{"latitude":12.9730,"longitude":77.5947},
{"latitude":12.9716,"longitude":77.5946}
]
}'
```

---

## Update Vehicle Location

```bash
curl -X POST http://localhost:8080/vehicles/location \
-H "Content-Type: application/json" \
-d '{
"vehicle_id":"<UUID>",
"latitude":12.9716,
"longitude":77.5946,
"timestamp":"2026-07-15T10:30:00Z"
}'
```

---

## Get Violations

```bash
curl http://localhost:8080/violations/history
```

---

## Docker Commands

Start

```bash
docker compose up --build
```

Stop

```bash
docker compose down
```

Rebuild

```bash
docker compose up --build
```

---

# Notes

- PostgreSQL data is stored in a Docker volume.
- WebSocket endpoint: `ws://localhost:8080/ws/alerts`
- The frontend connects to the backend through Axios.
- Geofence detection uses a point-in-polygon algorithm.
