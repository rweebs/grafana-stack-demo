# Docker Compose Observability Example

This project demonstrates a minimal observability stack using Docker Compose, featuring:

- **Loki** for log aggregation
- **Promtail** for log shipping from Docker containers
- **Grafana** for log and metrics visualization
- **Prometheus** for metrics collection
- **Node Exporter** for host metrics
- **A sample Go app** that emits logs

## Architecture & Data Flow
- The Go app writes logs to stdout.
- Promtail scrapes logs from containers labeled `logging=promtail` and ships them to Loki.
- Loki stores and indexes logs.
- Prometheus scrapes metrics from Node Exporter and itself.
- Grafana is pre-provisioned with Loki and Prometheus as data sources for dashboards.

## Usage
1. Build and start all services:
	```sh
	docker compose up -d --build
	```
2. Access Grafana at [http://localhost:3000](http://localhost:3000) (admin, no password required).
3. View logs and metrics in the prebuilt dashboard.

## Adding New Apps
Add a new service to `docker-compose.yaml` with the label `logging=promtail` to have its logs automatically collected.

---

See configs in each service directory for more details.
