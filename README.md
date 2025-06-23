## Migration

migrate -path migrations -database "postgres://postgres:postgres@localhost:5432/mydb?sslmode=disable" up


or go run cmd/migrate/main.go 