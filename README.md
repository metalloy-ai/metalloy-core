# metalloy-core

## Overview

metalloy core is the backbone service for handling all users and the authentication in the metalloy ai logistics platform.

---

## Getting Started

### Prerequisites

- [Golang](https://golang.org/doc/install) minimum version 1.13

- Clone the repository

```bash
git clone https://github.com/LogiFlow/logiflow-core
```

### Running the service

- Install the dependencies

```bash
go mod download
```

- Run the service in development mode

```bash
make dev
```

- Run the service in production mode

```bash
make [your os]
make run-[your os]
```

### Running the tests

- Run all tests

```bash
make testAll
```

- Run unit tests

```bash
make test[Unit]
```

---

## Documentation

- [Endpoints Doc](./docs/api-endpoints.md)
- [Architecture Doc](./docs/architecture.md)
