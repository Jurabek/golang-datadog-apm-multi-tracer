# Golang Datadog using multiple tracers Mux -> DynamoDB

### How to reproduce it:

- replace `DD_API_KEY` in docker-compose.yml file
- run `docker-compose up --build -d`
- run `sh create-table.sh`
- `curl http://localhost:3000/getItems`

check datadog tracing list.
