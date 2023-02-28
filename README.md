# Example Go Rest API with a PostgreSQL database

This repository demonstrates how to fully wire up and deploy a Golang REST API with a PostgreSQL database backend.

The application demonstrates:

- `REST API` using [Echo](https://echo.labstack.com/)
- `PostgreSQL integration` using [PGX](https://github.com/jackc/pgx)
- `Database migrations` using [Atlas](https://atlasgo.io/)

## Technology Choices

The author of this repository carries opinions about how to best organize an application / repository.  Those opinions / choices are described below:

1. No ORM (Object Relational Mapper) - Golang has a few ORMs that are popular, in particular [GORM](https://gorm.io/) and [Ent](https://entgo.io/).
   1. ORMs hide SQL, and SQL is already super declaritive and easy to read
   2. ORMs typically require some schema setup, and code generation, and knowing what goes where
   3. It can be difficult to optimize ORM SQL
2. Echo for the REST API - There are a lot of http frameworks in the golang echosystem, including [Echo](https://echo.labstack.com/) and [Gin](https://gin-gonic.com/).  Echo seemed to provide a lot out of the box, including things like Auth.
3. PGX for the database - `database/sql` would be fine as well, but PGX is PostgreSQL optimized and includes connection pooling.
4. Taskfile instead of Makefile - There is nothing inherently wrong with Makefile, Taskfile is a reasonable alternative with simple, validatable YAML syntax
5. Pre-commit - makes sure that users cannot commit / push code that isn't up to standards

## Pre-requisites

1. Install [Docker](https://docs.docker.com/get-docker/).  Used for testing
2. Install [Pre-Commit](https://pre-commit.com/).  This project uses pre-commit to ensure code is all nice and tidy before others can see it.
3. Install the pre-commit hooks by running `pre-commit install`
4. Install [Atlas](https://atlasgo.io/getting-started).  Atlas is used for database migrations.  Note: you can skip this step and just rely on docker

## Project Structure

- `cmd` - this is where the default config and the main app lives
- `database` - this is where the **schema.hcl** file lives.  Modify this file to alter the database
- `pkg` - this is where most of the code lives
  - `common`
    - `config` - for loading the config file / incorporating environment variable overrides
    - `db` - the underlying database in PGX, and domain specific Repositories
    - `handler` - for handling each of the type of echo requests / routes
    - `models` - core domain model classes, surfaced in the API and used in the Repositories
    - `router` - where we map echo routes to handlers

## Running tasks

This project uses [Taskfile](https://taskfile.dev/) for running tasks.  The following tasks are available

- `build`:            Builds a local executable, outputs to out/bin/gomin
- `clean`:            Cleans up build artifacts, including out, bin, and test reports
- `d.build`:          Builds the docker iamge, marks it as latest
- `d.down`:           Shuts down all docker containers in the docker compose file
- `d.up`:             Starts up all docker containers, builds and runs the API as well
- `db.migrate`:       Runs the database migration, ensures that the local postgres database is running
- `db.up`:            Starts the database WITHOUT migrations
- `server.run`:       Starts the database, runs migrations, builds the server, and starts the server
- `test`:             Runs all of the tests in all of the go source directories

For some common tasks:

- `task db.migrate` to start the local postgres database and run migrations to get the database current
- `task server.run` ensures the database is running and builds and starts the rest server - it is available at http://localhost:1323
- `task test` you must start the database first using `task db.migrate` and then you can run tests

## Writing Tests

This project primarily uses _E2E_ tests hitting the API directly and using the docker postres database.

Tests live in `pkg/common/router` - create a new `xxx_test.go` file as needed

## Roadmap

1. [x] - Simple REST API
2. [x] - Add initial Makefile
3. [x] - Add pre-commit
4. [x] - Initial database setup
5. [x] - Incorpoate database into REST API
6. [x] - Integration tests for database
7. [x] - E2E tests for REST API
8. [x] - Add github build for golang
9. [x] - Add docker packaging
10. [ ] - Add docker image build in Github Actions on Tag
11. [ ] - Terraform plan for RDS
12. [ ] - Terraform plan for ECS
13. [ ] - Github Actions to run terraform and deploy to AWS
