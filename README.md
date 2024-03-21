# Task Manager
Task Manager - API server to manage Task

## External library used :
- gin
- sqlx
- migrate
- pq

## Design Pattern :
an MVC application that following DDD structure to achive clean code.

## How To Set Up:
 1. Clone the Repo `git clone `
 2. Download docker images for postgresDb `docker run --name task-manager -e POSTGRES_PASSWORD=task-manager -p 5432:5432 -d postgres && sleep 3 && docker exec -it task-manager psql -U postgres -d postgres -c "CREATE DATABASE task_manager;"`
 3. Change directory to cmd `cd cmd/`
 4. Run the API Server `go run .`
