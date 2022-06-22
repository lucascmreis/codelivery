# Codelivery - Study Repo (WIP)

This repository is part of my studies about new technologies for implementing reliable real-time applications.
It includes a case used in the training "Imersão Full Cycle".

## About the project

- Delivery system that allows you to view the delivery vehicle in real time
- Possibility of multiple deliverers simultaneously
- Simulator service that sends geographic positions in real time (can be replaced with real gps data or mobile app)
- Delivered data, as well as positions, will be from work on Elasticsearch for future statistical analysis

## Technologies

- Golang (simulator)
- Nest.js & MongoDB (backend and database)
- React & Material UI (fronted)
- Kafka & Kafka Connect
- Elasticsearch & Kibana
- Docker e Kubernetes
- Istio, Kiali, Prometheus & Grafana


## Simulator

- Server that emulates the gps data in real time wrote in Golang.

### Running

1. Create Dockerfile and docker-compose to run the Go environment
2. Run the command

  ```bash
  docker-compose up -d
  ```

  You can check if the container is running

  ```bash
  docker-compose ps
  ```
  
3. Get into the container and run the command

```bash
docker exec -it gps-simulator bash
```
