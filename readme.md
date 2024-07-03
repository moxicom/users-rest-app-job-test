# Time Tracker Service Documentation

This documentation provides details on how to set up and run the Time Tracker service, including configuration, database setup, logging, and Docker integration.

## Table of Contents

1. [Introduction](#introduction)
2. [Configuration](#configuration)
3. [Database Migrations](#database-migrations)
4. [Running the Service](#running-the-service)
5. [Logging](#logging)
6. [Swagger Integration](#swagger-integration)

## Introduction

The Time Tracker service allows for managing users and tracking their tasks. It provides REST APIs for user management, task tracking, and integration with an external People Info API for enriching user data.

## Configuration

All configuration data is stored in a `.env` file. Below are the configuration parameters:

```env
POSTGRES_USER=postgres
POSTGRES_PASSWORD=postgres
POSTGRES_DB=db
DB_PORT=5432
SERVER_PORT=8080
SSL_MODE=disable
DB_HOST=postgres
API_ADDRESS=http://localhost:8080/api
```

- `POSTGRES_USER`: Username for PostgreSQL database
- `POSTGRES_PASSWORD`: Password for PostgreSQL database
- `POSTGRES_DB`: Name of the PostgreSQL database
- `DB_PORT`: Port on which PostgreSQL is running
- `SERVER_PORT`: Port on which the service will run
- `SSL_MODE`: SSL mode for database connection
- `DB_HOST`: Hostname of the PostgreSQL database
- `API_ADDRESS`: Address of the external People Info API

## Database Migrations

The database schema is created and managed using migrations. Ensure that the PostgreSQL service is running and accessible based on the configuration provided in the `.env` file.

To run the migrations, just run program. It migrates on startup:

This will set up the necessary tables and schemas required by the service.

## Running the Service

To run the Time Tracker service, follow these steps:

1. Ensure that Docker and Docker Compose are installed on your machine.
2. Create a `.env` file with the required configuration parameters.
3. Build and start the Docker containers using Docker Compose:

```sh
docker-compose up --build
```

This will build the Docker images and start the containers for the application and PostgreSQL database.

## Logging

The service uses structured logging for both debug and info levels. Logs include details about operations, warnings, and errors to aid in debugging and monitoring the service.

Example of logging usage in the code:

```go
log := h.log.With(slog.String("op", "handler.CreateUser"))
log.Info("User created successfully", slog.Uint64("user_id", uint64(userID)))
log.Error("failed to create user", err)
```

Also flag `envLog` to setup logger

## Swagger Integration

Swagger is integrated into the service for API documentation and testing. The Swagger documentation is automatically generated based on the provided annotations in the code.

To access the Swagger UI, navigate to the following URL after starting the service:

```
http://localhost:8080/swagger/index.html
```

This will provide an interactive interface for exploring and testing the available API endpoints.

## Additional Information

- **Enrichment of User Data:** When a new user is added, the service makes a request to an external People Info API to retrieve additional details about the user. This enriched data is then stored in the PostgreSQL database.
- **Task Management:** The service supports tracking the time spent on tasks by users, including starting and ending task periods.

For further details and API endpoint descriptions, please refer to the generated Swagger documentation.