# gomin

## Pre-requisites

1. Install [Pre-Commit](https://pre-commit.com/).  This project uses pre-commit to ensure code is all nice and tidy before others can see it.
2. Install the pre-commit hooks by running `pre-commit install`
3. Install [Atlas](https://atlasgo.io/getting-started).  Atlas is used for database migrations

## Getting Started

This project uses [Taskfile](https://taskfile.dev/) for running tasks.  The following tasks are available

- `task build` - creates the executable storing it in `out/bin` directory along with the config file
- `task clean` - cleans various build artifacts
- `task d.build` - builds the docker container
- `task d.down` - shuts down the docker containers
- `task d.up` - starts up the docker containers (postgres + gomin api)
- `task db.migrate` - useful when changing the `database/schema.hcl` file, applies changes to the local database
- `task db.up` - useful for ONLY starting the database (not the api)
- `task server.run` - runs the code locally (not using docker)
- `task test` - runs tests including coverage

## Roadmap

1. [x] - Simple REST API
2. [x] - Add initial Makefile
3. [x] - Add pre-commit
4. [x] - Initial database setup
5. [x] - Incorpoate database into REST API
6. [ ] - Integration tests for database
7. [ ] - E2E tests for REST API
8. [ ] - Add github build for golang
9. [x] - Add docker packaging
10. [ ] - Add docker image build in Github Actions on Tag
11. [ ] - Terraform plan for RDS
12. [ ] - Terraform plan for ECS
13. [ ] - Github Actions to run terraform and deploy to AWS
