# Example Go Rest API with a PostgreSQL database

This repository demonstrates how to fully wire up and deploy a Golang REST API with a PostgreSQL database backend.

The application demonstrates:

- `REST API` using [Echo](https://echo.labstack.com/)
- `PostgreSQL integration` using [PGX](https://github.com/jackc/pgx)
- `Database migrations` using [Atlas](https://atlasgo.io/)

## Use as a Template

**If you want to use this as a template, here a short guide:**

1. Modify the `go.mod` file `module` to be the name of your project
1. Modify the `database/schema.hcl` file to meet your database needs
1. Modify the `docker-compose.yml` file to use the name of your project
1. Modify `pkg/common/router/router` and add additional routes as needed
1. Add your domain models to `pkg/common/models` directory
1. Add your data access (repository) to the `pkg/common/db` directory

**In general, when you introduce a new separate domain concept:**

1. Update the `database/schema.hcl` file with new tables
1. Create a new route handler `pkg/common/handler/xxx.go`
1. Create model types `pkg/common/models/xxx.go`
1. Create a new repository `pkg/common/db/xxx.go`
1. Update the main `pkg/common/router/router.go` file to use your new routes / handlers
1. Add a new `pkg/common/router/xxx_test.go` file to test your new routes / handlers

**If you want to add additional config entries:**

1. Modify the `Config` struct in `pkg/common/config.go`
2. Add a default (typically works locally) value in `cmd/config.yaml`
3. Override the default value through an environment variable at runtime (i.e. not local) as needed

## Technology Choices

The author of this repository carries opinions about how to best organize an application / repository.  Those opinions / choices are described below:

1. No ORM (Object Relational Mapper) - Golang has a few ORMs that are popular, in particular [GORM](https://gorm.io/) and [Ent](https://entgo.io/).  However, the mapping ability in tools like [Scany](https://github.com/georgysavva/scany) take care of a lot of tedious work of working with databases (mapping rows to structs).  Declaritive SQL in code (imo) is preferable to code generation and clunky SQL-esque APIs.
2. [Atlas](https://atlasgo.io/) for database migrations.  There are 1000 ways to run migrations in golang, and unfortunately there doesn't seem to be a lot of consensus.  However, what I enjoy about Atlas is that it includes Terraform support and it allows you to define your schema in an HCL file and it magically figures out how to update your target database.  You can see this in `database/schema.hcl` and just a simple `atlas schema apply` and your database is now in sync.
3. [Echo](https://echo.labstack.com/) for the REST API - There are a lot of http frameworks in the golang echosystem, including [Echo](https://echo.labstack.com/) and [Gin](https://gin-gonic.com/).  Honestly, I found using the docs with Echo simpler, but don't have a strong opinion on the web framework side.  `net/http` might be just as good as the other choices, `Gin` being pretty popular.
4. [PGX](https://github.com/jackc/pgx) for database aaccess - `database/sql` would be fine as well, but PGX is PostgreSQL optimized and is a good choice when tied to Postgres.
5. [Taskfile](https://taskfile.dev/) instead of Makefile - There is nothing inherently wrong with Makefile, Taskfile is a reasonable alternative with simple, validatable YAML syntax
6. [Pre-commit](https://pre-commit.com/) - makes sure that users cannot commit / push code that isn't up to standards.  In this project we check formatting, linting, golang critic, among others to keep the repo tidy.

## Pre-requisites

1. Install [Docker](https://docs.docker.com/get-docker/).  Used for testing
2. Install [Taskfile](https://taskfile.dev/), required to do pretty much anything
3. Install [Pre-Commit](https://pre-commit.com/).  This project uses pre-commit to ensure code is all nice and tidy before others can see it.
4. Install the pre-commit hooks by running `pre-commit install`
5. **Optionally** install [Atlas](https://atlasgo.io/getting-started).  Atlas is used for database migrations.  **Note: you can skip this step and just rely on docker, as atlas is only needed to explore its abilities**

## Quick Start

Make sure you have Docker and Taskfile installed...

1. Run `task server.run`
   1. Starts up the postgres docker database
   2. Runs migrations so the database is ready
   3. Performs a build of the REST API
   4. Starts the REST API on port 1323.  Access on http://localhost:1323

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

- `build` - Builds a local executable, outputs to out/bin/gomin
- `clean` - Cleans up build artifacts, including out, bin, and test reports
- `d.build` - Builds the docker iamge, marks it as latest
- `d.down` - Shuts down all docker containers in the docker compose file
- `d.up` - Starts up all docker containers, builds and runs the API as well
- `db.migrate` - Runs the database migration, ensures that the local postgres database is running
- `db.up` - Starts the database WITHOUT migrations
- `server.run` - Starts the database, runs migrations, builds the server, and starts the server
- `test` - Runs all of the tests in all of the go source directories

## Writing Tests

This project primarily uses _E2E_ tests hitting the API directly and using the docker postres database.

Tests live in `pkg/common/router` - create a new `xxx_test.go` file as needed

## Environment Variables and Config

This project is setup to load a default config found in `cmd/config.yaml` that is overridable via **Environment Variables**.  For example, you can override the database url by setting an environment variable named `APP_DB_URL`.

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
10. [x] - Add docker image build in Github Actions on Tag
11. [ ] - Terraform plan for RDS
12. [ ] - Terraform plan for ECS
13. [ ] - Github Actions to run terraform and deploy to AWS
