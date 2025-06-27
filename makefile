# DSN строка для подключения
DSN = "postgres://postgres:postgres@localhost:5432/frodo?sslmode=disable"

# Директория с миграциями
MIGRATIONS_DIR = migrations

# Установка goose
install-goose:
	go install github.com/pressly/goose/v3/cmd/goose@latest

# Создание новой миграции
migrate-create:
	goose -s -dir $(MIGRATIONS_DIR) postgres $(DSN) create $(name) sql

# Применение всех новых миграций
migrate-up:
	goose -dir $(MIGRATIONS_DIR) postgres $(DSN) up

# Откат последней миграции
migrate-down:
	goose -dir $(MIGRATIONS_DIR) postgres $(DSN) down

# Просмотр статуса миграций
migrate-status:
	goose -dir $(MIGRATIONS_DIR) postgres $(DSN) status

migrate-reset:
	goose -dir $(MIGRATIONS_DIR) postgres $(DSN) reset