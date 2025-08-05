# Variables
DB_URL ?= $(shell echo $$DB_URL)

# Verificar que DB_URL esté configurada
check-db-url:
	@if [ -z "$(DB_URL)" ]; then \
		echo "❌ Error: DB_URL environment variable is not set"; \
		echo "💡 Example: export DB_URL='postgresql://user:pass@localhost:5432/dbname?sslmode=disable'"; \
		exit 1; \
	fi

# Comandos principales
migrate-up: check-db-url
	@echo "🔄 Running all migrations UP..."
	@go run cmd/migrate.go -cmd=up

migrate-down: check-db-url
	@echo "⬇️  Running all migrations DOWN..."
	@go run cmd/migrate.go -cmd=down

migrate-up-1: check-db-url
	@echo "🔄 Running 1 migration UP..."
	@go run cmd/migrate.go -cmd=up -steps=1

migrate-down-1: check-db-url
	@echo "⬇️  Running 1 migration DOWN..."
	@go run cmd/migrate.go -cmd=down -steps=1

migrate-version: check-db-url
	@echo "📊 Checking migration version..."
	@go run cmd/migrate.go -cmd=version

migrate-force: check-db-url
	@echo "⚠️  Usage: make migrate-force VERSION=1"
	@if [ -z "$(VERSION)" ]; then \
		echo "❌ Error: VERSION parameter is required"; \
		echo "💡 Example: make migrate-force VERSION=1"; \
		exit 1; \
	fi
	@go run cmd/migrate.go -cmd=force -steps=$(VERSION)

# Comandos de desarrollo
migrate-status: migrate-version

migrate-reset: migrate-down migrate-up

# Ayuda
help:
	@echo "🚀 Migration commands:"
	@echo "  make migrate-up        - Apply all pending migrations"
	@echo "  make migrate-down      - Revert all migrations"
	@echo "  make migrate-up-1      - Apply 1 migration"
	@echo "  make migrate-down-1    - Revert 1 migration"
	@echo "  make migrate-version   - Show current migration version"
	@echo "  make migrate-status    - Show migration status"
	@echo "  make migrate-reset     - Reset all migrations (down + up)"
	@echo "  make migrate-force VERSION=X - Force version to X"
	@echo ""
	@echo "💡 Make sure to set DB_URL environment variable:"
	@echo "   export DB_URL='postgresql://user:pass@localhost:5432/dbname?sslmode=disable'"

.PHONY: check-db-url migrate-up migrate-down migrate-up-1 migrate-down-1 migrate-version migrate-force migrate-status migrate-reset help