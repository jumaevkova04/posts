# 🛠 Project Setup & Run Guide

This guide helps you get the application up and running using Docker and Go, including environment setup and database migrations.

---

## 🚀 Steps to Run the Project

### 1. Create Configuration Files

Copy the example configuration files into working config files:

```bash
cp ./config.json.example ./config.json
cp ./dbconfig.yml.example ./dbconfig.yml
cp ./docker-compose.env.example ./docker-compose.env
```

### 2. Build and Run the Application

#### Option A: Using Docker

Use Docker Compose to build the containers and start the application in the background:

```bash
docker-compose --env-file docker-compose.env up --build -d
```

#### Option B: Without Docker

Make sure you have Go installed, then build and run the application directly:

```bash
go build -o app ./cmd/posts/main.go
./app
```

### 3. Run Database Migrations

Install the sql-migrate CLI tool if not already installed:

```bash
go install github.com/rubenv/sql-migrate/sql-migrate@latest
```

Then check migration status and apply pending migrations:

```bash
sql-migrate status
sql-migrate up
```

### 📝 Notes

Ensure Docker, Docker Compose, and Go are installed on your system.

You may need to adjust the .example files to match your environment before use.

The migration tool uses dbconfig.yml for database connection settings.

To stop services, use:

```bash
docker-compose down
```

### ✅ Helpful Commands

```bash
# Rebuild containers after making changes
docker-compose --env-file docker-compose.env up --build -d

# View logs
docker-compose logs -f

# Stop and remove containers
docker-compose down
```

---

### 🌐 Application Host and Port

By default, the application listens on:

- **Host:** `localhost` (or `0.0.0.0` if configured)
- **Port:** `8080`

You can customize these values in your configuration files (`config.json` or environment variables).

---

### 📖 API Documentation (Swagger)

Once the app is running, the Swagger UI for the API documentation is available at:

http://localhost:8080/swagger/index.html

> Adjust the host and port accordingly if you changed them in your config.

---