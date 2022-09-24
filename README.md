## Project description
The service allows you to create notes and attach tags to them.
The functionality of the service is very simple, there are 3 entities:
- user
- note
- tag

You can create a user, authorize him, create notes and attach tags to these notes.
JWT token authorization is used here.

The project is made using the principles of clean architecture.

The project consists of 3 microservices and each runs in a container:
1. Nginx
2. Go-service itself
3. PostgreSQL

## First start
For the initial launch of the service, you must run the command in the terminal:
```bash
make dockerup
```

After launching all the service containers (Nginx, service, Postgres), it is necessary to migrate
the Postgres database to create tables. MacOS users need to run a command in the terminal to install this software:
```bash
brew install golang-migrate
```
To migrate the database, run the command:
```bash
make migrate-up
```

To stop the service, run the command:
```bash
make dockerstop
```

## DB Migrations
Creating migration files:
```bash
make migrate-create
```
Applying migration files:
```bash
make migrate-up
```
Rollback to the previous version of the migration:
```bash
make migrate-down
```

## Unit-tests
To run the `unit tests`, run the command:
```bash
make test
```

## Linter
MacOS users need to run a command in the terminal to install this software:
```bash
brew install golangci-lint
```
All the checks can be found in the file `.golangci.yml`.

To run the check, run the command in the terminal:
```bash
make lint
```

## Er diagram
<p align="left">
    <img src="assets/er.png" width="700">
</p>

## Swagger
`Swagger` available at the link: http://localhost:8000/swagger
<p align="left">
    <img src="assets/swagger.png" width="700">
</p>