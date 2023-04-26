# API Documentation

## response body structure

```bash
{
    "code": int,
    "message": string,
    "body": {}
}
```

Default Url: `http://localhost:2000/api/v1`

> **Note** Settings are interchangeable by updating the `.env` file refer to the `.env.template` file

```bash
API_VERSION="v1"
PORT=2000
HOST="localhost"
```

---

## Base Routes

```bash
    - [GET] /
    - [GET] /health
```

---

## User Routes

```bash
    - [GET] /users
    - [GET] /users/user/:username
    - [PUT] /users/user/:username
    - [DEL] /users/user/:username
```

## User Sub-Routes (Address)

```bash
    - [GET] /users/user/:username/address
    - [PUT] /users/user/:username/address
```

### Address Request Body

```bash
{
    "street_address": string,
    "city": string,
    "state": string,
    "country": string,
    "postal_code": string
}
```

### Address Response

```bash
{
    "address_id": int,
    ... Same as Address Body ...
}
```

---

## Auth Routes

```bash
    - [POST] /auth/login
    - [POST] /auth/register
```

### Login Body

```bash
{
    "username": string,
    "password": string
}
```

### Register Body

```bash
{
    "username": string,
    "email": string,
    "password": string,
    "user_type": string, <- "admin" | "supplier" | "customer"
    "first_name": string,
    "last_name": string,
    "phone_number": string,
    "street_address": string,
    "city": string,
    "state": string,
    "country": string,
    "postal_code": string
}
```
