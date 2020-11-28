# Stock Service

## Getting started

Run GRPC server with Compose on port 9090:

```shell
docker-compose up
```

## Development

> Prerequisites: `docker`, `docker-compose`, `go@1.15`, `git`, `make`.

### Code style

Consistent code style enforced by `gofmt`, `EditorConfig` tools and `golangci-lint` linter.

Format code:

```shell
make format
```

Run linter:

```shell
make lint
```

### Tests

Run tests:

```shell
make test
```

### Docker

Build docker image `stock-service:latest`:

```shell
make docker
```

Build, run linter and tests in dev docker image `stock-service-dev:latest`:

```shell
make docker-dev
```

### CI

There are configured GitHub actions for build, lint, tests.
See [.github/workflows](.github/workflows) directory.
