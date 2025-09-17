# Database Migration Guide

## Prerequisites

- Docker and Docker Compose installed
- PostgreSQL client tools (optional, for manual database access)

## Setup
### Install golang-migrate CLI for macOS and Windows

Follow the instructions from the [official golang-migrate repository](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate) to install the CLI tool.

### Steps to for migration
- export POSTGRES_URL=postgres://postgres:postgres@localhost:5438/learn-go?sslmode=disable
- migrate -database ${POSTGRES_URL} -path internal/db/migrations up