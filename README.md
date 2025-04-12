# Fullstack_Golang
Fullstack App Utilizing HTMX FE + GO CRUD API


### Install Dependencies
go mod init go-htmx-crud
go get github.com/go-chi/chi/v5
go get github.com/joho/godotenv
go mod tidy

### Start Server
cd backend
go run main.go

### Using Docker
docker compose up --build

- Cleanup
docker compose down -v

### Open in Browser
Visit http://localhost:8080

### (Optional)
DATABASE_URL=postgres://user:password@db:5432/taskmanager?sslmode=disable

### Why this works

Isolation: PostgreSQL runs in its own container.

Reproducibility: Same environment everywhere (Dev/Prod).

Efficiency: Alpine images keep containers lightweight.

Dev UX: Hot-reload for frontend files.
### Project Structure
Fullstack_Golang/
├── backend/
│   ├── main.go
│   ├── db/
│   │   └── database.go
│   ├── handlers/
│   │   └── tasks.go
│   ├── models/
│   │   └── task.go
│   ├── templates/
│   │   ├── render.go    
│   │   ├── base.html
│   │   ├── index.html
│   │   ├── task.html
│   │   └── form.html
│   └── go.mod
├── frontend/
│   └── static/
│       └── styles.css
├── Dockerfile           # For Go backend
├── docker-compose.yml   # Orchestrates Go + PostgreSQL
└── .env                 # Environment variables