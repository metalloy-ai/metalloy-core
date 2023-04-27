# metalloy-core

## Overview

metalloy core is the backbone service for handling all users and the authentication in the metalloy ai logistics platform.

## Table of Contents

- [metalloy-core](#metalloy-core)
  - [Overview](#overview)
  - [Table of Contents](#table-of-contents)
  - [Getting Started](#getting-started)
    - [Prerequisites](#prerequisites)
    - [Running the service](#running-the-service)
    - [Running the tests](#running-the-tests)
  - [Documentation](#documentation)

---

## Getting Started

### Prerequisites

- [Golang](https://golang.org/doc/install) minimum version 1.13

- Clone the repository

```bash
git clone https://github.com/metalloy-ai/metalloy-core
```

### Running the service

To get started with metalloy-core, follow these steps:

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

- [API Doc](./docs/api.md)
- [Architecture Doc](./docs/architecture.md)
