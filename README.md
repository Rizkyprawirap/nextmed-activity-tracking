# NextMed Activity Tracking — Setup Guide

## 📌 Requirements

Make sure these are installed:

- **Go** ≥ 1.21
- **Docker & Docker Compose**
- **PostgreSQL** (via Docker)
- **Redis** (via Docker)

## 🐳 Start Docker Services

The project includes a `docker-compose.yml`.  
Run this to start PostgreSQL + Redis:

```sh
docker compose up -d or docker compose -f docker-compose.yml up
```

This will start:

- PostgreSQL on port **5432**
- Redis on port **6379**

## ⚙️ Environment Variables

Create **.env** file:

```
APP_ENV=dev
PORT=8888
DB_HOST=localhost
DB_PORT=5454
DB_USER=nextmed_user
DB_PASSWORD=nextmed_pass
DB_NAME=nextmed_db
DB_MAX_OPENCONN=20
DB_MAX_IDLECONN=10
DB_MAX_LIFETIME=300
DB_MAX_IDLETIME=60
REDIS_HOST=localhost
REDIS_PORT=6464
RATE_LIMIT=500
CORS_ORIGIN=*
JWT_SECRET_ADMIN=secret
```

---

## ▶️ Install Dependencies

```sh
go mod tidy
```

---

## ▶️ Run the Application

```sh
go run cmd/server/main.go
```

If everything is correct, you should see:

```
[GIN-debug] Listening on :8888
```

---
