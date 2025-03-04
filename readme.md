## CPF-CNPJ-API

This project is a RESTful API designed for user authentication and CPF/CNPJ validation. It follows Clean Architecture principles, ensuring modularity, scalability, and maintainability.

### Features

- ✅ User authentication (registration and login)
- ✅ CPF/CNPJ validation and verification
- ✅ Secure password hashing and authentication middleware
- ✅ Message queue integration using RabbitMQ for asynchronous processing
- ✅ Prometheus integration for real-time metrics and monitoring
- ✅ Clean and scalable architecture with separation of concerns

### Tech Stack

- Golang – Core API development
- PostgreSQL – Relational database
- RabbitMQ – Message queue for background processing
- Prometheus – Monitoring and performance metrics
- Docker – Containerized environment
- JWT – Secure authentication and session management

## How to run the project?

### 📌 Prerequisites

Before running the project, ensure you have the following installed:

- [Docker](https://docs.docker.com/get-docker/)
- [Docker Compose](https://docs.docker.com/compose/install/)

### Setup

1️⃣ **Clone the repository**

```bash
git clone https://github.com/devjunioralves/cpf-cnpj-api
git clone https://github.com/devjunioralves/cpf-cnpj-client
cd cpf-cnpj-api
```

2️⃣ **Duplicate the .env.example file and rename it to .env**

```bash
cp .env.example .env
```

3️⃣ **Running the Project**

- Ensure that cpf-cnpj-api and cpf-cnpj-client are located at the same level in your file structure. For example:

```bash
~/Projects/cpf-cnpj-api/
~/Projects/cpf-cnpj-client/
```

```bash
docker compose up -d
```

### What would I improve on this project?

- Unit Tests: Implement more unit tests to cover edge cases and ensure robustness of individual components.
- Integration Tests: Increase the coverage of integration tests to ensure seamless communication between services and databases.
- Observability & Log Storage: Improve observability by setting up a centralized logging system (e.g., ELK stack) to store and analyze application logs.
- Crypto Key Management: The current storage method for the encryption key (CRYPTO_KEY) could be improved for better security and permission control. Consider using secure methods like:
  - AWS Secrets Manager: Securely store and control access to the key.
  - HashiCorp Vault: Provides detailed access control and auditing.

```
.
├── cmd
│ ├── consumer
│ │ └── worker.go
│ └── main.go
├── internal
│ ├── app
│ │ ├── routes.go
│ │ └── server.go
│ ├── domain
│ │ ├── models
│ │ │ ├── cpf_cnpj.go
│ │ │ └── user.go
│ │ ├── repositories
│ │ │ ├── cpf_cnpj_repository.go
│ │ │ └── user_repository.go
│ │ ├── services
│ │ │ ├── auth_service.go
│ │ │ └── cpf_cnpj_service.go
│ │ ├── validators
│ │ │ └── validators.go
│ ├── infrastructure
│ │ ├── repositories
│ │ │ ├── cpf_cnpj_repository.go
│ │ │ └── user_repository.go
│ │ ├── database.go
│ │ └── rabbitmq.go
│ ├── mappers
│ │ └── cpf_cnpj_dto.go
│ ├── presentation
│ │ ├── middlewares
│ │ │ ├── auth_middleware.go
│ │ │ ├── cors_middleware.go
│ │ │ └── metrics_middleware.go
│ │ ├── auth_handler.go
│ │ ├── cpfcnpj_handler.go
│ │ └── request.go
```

## Project Structure Overview

- **`cmd/`**: Contains the application entry points.
  - **`consumer/`**: A worker process handling asynchronous tasks via RabbitMQ.
- **`internal/`**: The core application logic, divided into multiple layers.

  - **`app/`**: Responsible for application setup and route definitions.
  - **`domain/`**: Contains business logic and domain models.

    - **`models/`**: Defines core entities, such as:
    - **`repositories/`**: Defines repository interfaces for data access.
    - **`services/`**: Contains business logic.
    - **`validators/`**: Business-related validations.

  - **`infrastructure/`**: Implements data access and external integrations.
    - **`repositories/`**: Implements the domain repository interfaces.
  - **`mappers/`**: Responsible for transforming domain objects into data transfer objects (DTOs).
  - **`presentation/`**: Manages API endpoints and middleware.
    - **`middlewares/`**: Implements request middleware.
