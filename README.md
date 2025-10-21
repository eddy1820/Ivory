# 🧱 Ivory Project

Ivory is a modular backend system written in **Golang**, designed for web server logic and API services.  
The project uses **Docker Compose** to orchestrate all dependencies such as **MySQL**, **Redis**, **Nginx**, and **Swagger UI**.

---

## 🚀 Setup Project

### 1️⃣ Build all services
```bash
docker compose build
```
### 2️⃣ Start all containers in the background
```bash
docker compose up -d
```

### Swagger API Documentation
http://localhost:7500/swagger/index.html

## 🚀 Debug Mode (optional)
```bash
docker compose -f docker-compose.debug.yml up -d
```
Debug connection (Delve)
Host:localhost
Port:7800

Ivory/
├── gate/
│   ├── cmd/server/main.go       # main entry point
│   ├── docs/                    # swagger auto-generated docs
│   ├── Debug.Dockerfile         # debug mode image
│   ├── Swag.Dockerfile          # swagger build image
│   ├── handler/                 # API handlers
│   ├── domain/                  # business domain models
│   ├── repository/              # database layer
│   └── usecase/                 # application logic
├── nginx/
│   └── default.conf             # nginx proxy config
├── docker-compose.yml           # normal mode
├── docker-compose.debug.yml     # debug mode (with Delve)
└── README.md