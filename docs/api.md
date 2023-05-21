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
    - [GET]   /api/v1/users?pageIdx={optional}&pageSize={optional}
    - [GET]   /users/user/:username
    - [PATCH] /users/user/:username
    - [DEL]   /users/user/:username
```

## User Sub-Routes (Address)

```bash
    - [GET]   /users/user/:username/address
    - [PATCH] /users/user/:username/address
```

### User Response Body

```bash
{
    "user_id": string,
    "username": string,
    "email": string,
    "user_type": string, <- "admin" | "supplier" | "customer"
    "first_name": string,
    "last_name": string,
    "phone_number": string,
    "address_id": int64,
    "registration_date": date,
    "street_address": string,
    "city": string,
    "state": string,
    "country": string,
    "postal_code": string
}
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

### Address Response Body

```bash
{
    "address_id": int,
    ... Same as Address Request Body ...
}
```

---

## Auth Routes

```bash
    - [POST] /auth/login
    - [POST] /auth/register
    - [POST] /auth/forgotPassword
    headers - Authorization: Bearer {token}
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
