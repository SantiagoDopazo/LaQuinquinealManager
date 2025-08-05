@echo off
if "%DB_URL%"=="" (
    echo ❌ Error: DB_URL environment variable is not set
    echo 💡 Example: set DB_URL=postgresql://user:pass@localhost:5432/dbname?sslmode=disable
    exit /b 1
)

if "%1"=="" (
    echo 🚀 Migration commands:
    echo   migrate.bat up        - Apply all pending migrations
    echo   migrate.bat down      - Revert all migrations
    echo   migrate.bat up1       - Apply 1 migration
    echo   migrate.bat down1     - Revert 1 migration
    echo   migrate.bat version   - Show current migration version
    echo   migrate.bat force VERSION - Force version to VERSION
    echo.
    echo 💡 Make sure to set DB_URL environment variable:
    echo    set DB_URL=postgresql://user:pass@localhost:5432/dbname?sslmode=disable
    exit /b 0
)

if "%1"=="up" (
    echo 🔄 Running all migrations UP...
    go run cmd/migrate.go -cmd=up
) else if "%1"=="down" (
    echo ⬇️  Running all migrations DOWN...
    go run cmd/migrate.go -cmd=down
) else if "%1"=="up1" (
    echo 🔄 Running 1 migration UP...
    go run cmd/migrate.go -cmd=up -steps=1
) else if "%1"=="down1" (
    echo ⬇️  Running 1 migration DOWN...
    go run cmd/migrate.go -cmd=down -steps=1
) else if "%1"=="version" (
    echo 📊 Checking migration version...
    go run cmd/migrate.go -cmd=version
) else if "%1"=="force" (
    if "%2"=="" (
        echo ❌ Error: VERSION parameter is required
        echo 💡 Example: migrate.bat force 1
        exit /b 1
    )
    echo ⚠️  Forcing version to %2...
    go run cmd/migrate.go -cmd=force -steps=%2
) else (
    echo ❌ Unknown command: %1
    echo Available commands: up, down, up1, down1, version, force
    exit /b 1
)