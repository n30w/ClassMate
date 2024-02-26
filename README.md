# Darkspace

Academic project management platform for SE-2024.

## Requirements

One must have installed these utilities to start development:

- Taskfile
- Docker
- Postgresql
- Go
- Node

## Software

Below is a list of software one may use to code this project:

- DataGrip (Interface with databases)
- VS Code
- GoLand
- Docker Desktop
- Podman

## Structure

This project consists of a decoupled frontend and backend. The frontend is written in Typescript, the backend is written in Go.

### Frontend

We are using ```yarn``` and NOT npm.

- Next.js
  - Next Auth
  - Jest (Unit testing)
  - Cypress (Component and E2E testing)
  - [Playwright](https://playwright.dev/)
  - [Mock Service Worker](https://mswjs.io/)

### Backend

Our database for the backend is SQL based.

- Go
  - Gin
  - pq
- Postgresql
- Nginx
- Ngrok

I'm not sure whether we are going to use Nginx or Ngrok. Ngrok requires less work.

#### Directory Structure

The backend directory is structured in this manner:

```bin``` is where the binaries for deployment are placed.

```cmd/api``` is where application specific code goes. For example, routing, HTTP writing, and authentication.

```internal``` is code that is imported into cmd/api, and therefore is not *application specific* per se, in that code in here is not specific to our exact use case. For example, it contains the database access layer or data validation.

```migrations``` is where SQL database migrations live.

```remote``` contains Docker files and anything needed for deployment purposes, like setup scripts.

## Getting Started

This project uses [Taskfile](https://taskfile.dev) as a Makefile replacement. This is used to run tests and synchronize docker containers. Unless specified otherwise, all task commands must be run in the root directory of the project.

### Before First Run

Before ever starting a development environment, please run ```task first-time```. This runs ```yarn install``` and ```go mod tidy``` in their respective directories.

### Development

To start the backend and frontend for development, run:

```task dev```

There are several types of tasks, some of which are ```dev```, ```build```, ```test``` (the full list of tasks can be found in ```Taskfile.yml```). Typing ```Task (name of task)``` runs that task. If one wanted to run a task for either the frontend or backend, simply do ```task front:(name of task)``` or ```task back:(name of task)```. Therefore:

```task front:dev``` or ```task back:dev``` runs each ```dev``` task separately.

The frontend exists at http://localhost:3000/ and the backend exists at http://localhost:6789/api/v1/

## Testing

We must implement endpoint testing.

## Additional Notes

- The front end was created with the command ```npx create-next-app@latest --no-git```, run in the root directory.

## Useful Links

- [Setting up and using postgresql on Mac](https://www.sqlshack.com/setting-up-a-postgresql-database-on-mac/)
- [Setting postgresql on Windows](https://www.prisma.io/dataguide/postgresql/setting-up-a-local-postgresql-database#setting-up-postgresql-on-windows)

### Postgresql

- [Tuning postgresql for memory](https://www.enterprisedb.com/postgres-tutorials/how-tune-postgresql-memory)
- [Postgresql tuner webapp](https://pgtune.leopard.in.ua/)

### Docker

- [Running postgresql in a Docker container](https://www.docker.com/blog/how-to-use-the-postgres-docker-official-image/)
- [Golang-Nginx-Postgres Docker Compose](https://github.com/docker/awesome-compose/tree/master/nginx-golang-postgres)
