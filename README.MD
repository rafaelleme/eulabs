[![GoVersion](https://img.shields.io/badge/Go-v1.19-blue.svg)](https://tip.golang.org/doc/go1.19)
[![ProjectVersion](https://img.shields.io/badge/Version-v1.0-green.svg)]()

# Go API
Application developed in Go language with echo framework.

The application is a simple product CRUD

- [Requirements](#requirements)
- [Configuration](#configuration)
- [Running](#running)
- [Documentation](#documentation)

## Requirements

| Tool                                               | Versão |
|----------------------------------------------------|--------|
| [docker](https://www.docker.com/)                  | latest |
| [docker compose](https://docs.docker.com/compose/) | latest |


## Configuration

Create the .env file through the base .env.example file

```shell
cp .env.example .env
```

## Running

In the root of the project, run:
```shell
make develop
```

Then you will need to run the sql script found in /db/init.sql

Note: all connection information is in the .env file

If you are unable to connect with the service name in Docker Compose, inspect the MySQL container and get the IP of the direct container.
```shell
docker inspect CONTAINER_ID
```

## Documentation
Postman documentation is exported in .json at /doc/API.postman_collection.

If you have questions about how to export a collection, consult the postman documentation

## Credits
Rafael Leme (Author)
