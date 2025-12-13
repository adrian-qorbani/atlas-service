# Atlas Service

**Atlas Service** is a data-oriented domain-driven production-grade Go backend project designed to explore real-world service architecture, authentication, observability, and Kubernetes-based deployment.

The project demonstrates how modern Go services are structured, secured, observed, and deployed in a distributed environment. It is built as a learning project with Ardan Labs’ *Service* architecture and practices guidance.

## 📦 Tech Stack

- Language: Go
- Containerization: Docker
- Orchestration: Kubernetes (Kind)
- Auth: JWT, RSA, OPA
- Database: PostgreSQL
- Observability: Prometheus, Grafana, Tempo, Loki

## ✨ Key Concepts

- Modular service-oriented architecture
- Data-oriented domain-driven design (DDD)
- Authentication and authorization using JWT and RSA keys
- Policy-based authorization with Open Policy Agent (OPA)
- Structured logging and observability
- Kubernetes-based local development workflow
- Production-style Makefile-driven tooling

---

## 🧱 Architecture Overview

The system currently consists of the following components:

- **Sales Service**
  - Core API service
  - Handles domain logic and HTTP endpoints
  - Protected via authentication middleware

- **Auth Service**
  - Issues and validates JWT tokens
  - Uses RSA key pairs for signing and verification
  - Integrates OPA for authorization policy evaluation

- **Database**
  - PostgreSQL
  - Deployed as a StatefulSet in Kubernetes

Each service is containerized and deployed independently inside a local Kubernetes (Kind) cluster.

---

## 🔐 Authentication & Authorization

- **Authentication**
  - JWT-based authentication
  - Tokens signed with RSA private keys
  - Public keys used by services to verify tokens

- **Authorization**
  - Middleware-based authorization in services
  - Open Policy Agent (OPA) used for policy decisions
  - Clear separation between authentication (who you are) and authorization (what you can do)

---

## 📊 Observability

The project includes a full observability stack:

- **Metrics:** Prometheus
- **Dashboards:** Grafana
- **Tracing:** Tempo
- **Logs:** Loki + Promtail
- **Runtime Stats:** Expvar, statsviz

This setup mirrors what is commonly used in production environments.

---

## 🚀 Running the Project Locally

### Prerequisites
- Go (1.25+)
- Docker
- Kind
- kubectl
- kustomize

### Common Commands

Run the Sales API locally:

Build Docker images, create  and start a local cluster, apply kustomize and deploy all services:
```bash
make dev-run
```

View logs:
```bash
make dev-logs-sales
make dev-logs-auth
Run tests and linters:
```

For the full list of commands, see the Makefile.
