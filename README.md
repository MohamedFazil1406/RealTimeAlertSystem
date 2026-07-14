# Geofencing & Real-Time Alert System

A full-stack geofencing and vehicle tracking system built with **Go**, **React**, **PostgreSQL (PostGIS)**, **WebSockets**, and **Docker**.

## Features

### Backend

- Vehicle Management
- Geofence Management
- Location Tracking
- Point-in-Polygon Detection
- Entry/Exit Detection
- Alert Configuration
- Violation History
- WebSocket Real-Time Alerts

### Frontend

- Dashboard
- Vehicle Management
- Interactive Leaflet Map
- Geofence Drawing
- Real-Time Alert Feed
- Violation History
- Responsive UI

---

## Tech Stack

### Backend

- Go
- Gin
- GORM
- PostgreSQL
- PostGIS
- Gorilla WebSocket

### Frontend

- React
- TypeScript
- Vite
- Tailwind CSS
- Leaflet
- Axios

### Infrastructure

- Docker
- Docker Compose

---

## Project Structure

```
RealTimeAlertSystem

backend/

frontend/

docker-compose.yml

README.md

SETUP.md
```

---

## API Endpoints

### Vehicles

```
POST /vehicles

GET /vehicles

POST /vehicles/location

GET /vehicles/location/{vehicle_id}
```

### Geofences

```
POST /geofences

GET /geofences
```

### Alerts

```
POST /alerts/configure

GET /alerts

GET /ws/alerts
```

### Violations

```
GET /violations/history
```

---

## Screenshots

- Dashboard
- Vehicle Page
- Geofence Map
- Alert Feed
- Violation History

(Add screenshots here)

---

## Architecture

```
React

↓

REST API

↓

Go Backend

↓

PostgreSQL + PostGIS

↓

WebSocket Alerts
```

---

## Deployment

Backend

Docker Container

Frontend

Docker Container

Database

PostgreSQL + PostGIS

---

## Author

Mohamed Fazil
