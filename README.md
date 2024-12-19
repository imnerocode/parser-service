# Service: 3D Model Parser

This project is a gRPC service designed to process 3D models in various formats. The service parses the models, converts their structure to JSON, and returns it to the client. Future stages will include storing the parsed data in a database.

## Project Structure

```
service/
├── proto/               # Protobuf definitions for services and messages
├── cmd/                 # Main application entry point
├── internal/            # Business logic implementation
├── config/              # Project configuration
├── Dockerfile           # Docker setup
├── docker-compose.yml   # Container orchestration
├── .env                 # Environment variables
├── .gitignore           # Git exclusions
├── .dockerignore        # Docker build exclusions
└── README.md            # Project documentation
```

## Technologies Used

- **Go**: Main programming language.
- **gRPC**: For efficient communication between client and server.
- **Protocol Buffers (Protobuf)**: To define services and serialize data.
- **Docker & Docker Compose**: For containerization and deployment.

## Setup Instructions

### Prerequisites

- Docker and Docker Compose installed.
- Go (v1.20 or later).

### Environment Variables

An `.env` file is used for project-specific configuration. Example:

```
GRPC_PORT=50051
PARSE_TIMEOUT=30s
```

### How to Run

1. Clone the repository:
    ```sh
    git clone <REPO_URL>
    cd service
    ```

2. Build and run the containers:
    ```sh
    docker-compose up --build
    ```

3. Access the service:
    - The gRPC server listens on the port specified in `GRPC_PORT` (default: 50051).

4. Stop the containers:
    ```sh
    docker-compose down
    ```

## Features

- Parse 3D models in common formats (e.g., OBJ, STL).
- Return the parsed structure as JSON.
- [Planned] Store the parsed data in a database.

## Contribution Guidelines

1. Fork the repository.
2. Create a branch for your changes:
    ```sh
    git checkout -b feature/new-feature
    ```
3. Commit your changes and push:
    ```sh
    git push origin feature/new-feature
    ```
4. Submit a pull request for review.
