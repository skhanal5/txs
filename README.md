# txs

## About
Backend REST API for a digital bank application that supports multiple users, accounts, and transactions. Exploring ACID properties in concurrency scenarios and intermittent failures.

## Local Setup

### Environment Variables

Take a look at `docker-compose.yml` and view the different environment
variables that are passed into each service. Either export all of them or define an `.env` file in the project root like so:

```
DB_NAME=txs
DB_USER=foo
DB_PASSWORD=bar
FLYWAY_URL=jdbc:postgresql://database:5432/txs
FLYWAY_USER=foo
FLYWAY_PASSWORD=bar
PGADMIN_DEFAULT_EMAIL=foo@test.com
PGADMIN_DEFAULT_PASSWORD=bar
```

### Docker Services
Run `make rund` to start each of the services. It will spin up the following:
- postgres
- flyway (temporarily)
- pgadmin (http://localhost:8080)
- dozzle (http://localhost:9090)
