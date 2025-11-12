# Atlas Service

**Atlas Service** is a production-grade web service written in Go, designed to explore distributed system design principles through a modular, service-oriented architecture. It focuses on implementing authentication, observability, and service communication patterns inspired by real-world production systems. This project follows the Ardan Labs Service [course](https://github.com/ardanlabs/service) to learn production-grade Go service design.

---

### 🔧 Tech Overview
- **Language:** Go
- **Architecture:** Modular microservice setup; Sales and Auth services
- **Auth:** JWT with RSA key rotation and middleware-based authorization, OPA for policy
- **Observability:** Prometheus, Grafana, Tempo, Loki, and Expvar monitoring
- **Deployment:** Kubernetes cluster with Kind
- **Containerization:** Docker + Kustomize

---

### ⚙️ Development Commands
Common make targets:

| Command | Description |
|----------|--------------|
| `make run` | Run the Sales API locally |
| `make build` | Build Docker image |
| `make dev-up` | Start local Kind cluster |
| `make dev-apply` | Deploy service to cluster |
| `make test` | Run tests and linters |

For full command list, see the [`Makefile`](./Makefile).

---

work-in-progress
