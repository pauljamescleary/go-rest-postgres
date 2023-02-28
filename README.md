# gomin

## Pre-requisites

1. Install [Pre-Commit](https://pre-commit.com/).  This project uses pre-commit to ensure code is all nice and tidy before others can see it.
2. Install the pre-commit hooks by running `pre-commit install`
3. Install [Atlas](https://atlasgo.io/getting-started).  Atlas is used for database migrations

## Getting Started

This project uses [Taskfile](https://taskfile.dev/) for running tasks.  The following tasks are available

* build:            Builds a local executable, outputs to out/bin/gomin
* clean:            Cleans up build artifacts, including out, bin, and test reports
* d.build:          Builds the docker iamge, marks it as latest
* d.down:           Shuts down all docker containers in the docker compose file
* d.up:             Starts up all docker containers, builds and runs the API as well
* db.migrate:       Runs the database migration, ensures that the local postgres database is running
* db.up:            Starts the database WITHOUT migrations
* server.run:       Starts the database, runs migrations, builds the server, and starts the server
* test:             Runs all of the tests in all of the go source directories

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
