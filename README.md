# API Documentation and Project Setup Guide

## Overview

This API is built using the Go programming language with the Gin framework. It uses Nginx, PostgreSQL, and Docker

#### Step 1

```bash
git clone https://github.com/thitiphongD/thitiphong_agnos_backend.git
cd thitiphong_agnos_backend
```

#### Step 2

```bash
docker-compose up -d
```

## API Documentation

### Base URL

The base URL for this API is `http://localhost:80`.

### Endpoints

#### 1. Initialize Password

- **Endpoint:** `POST /api/strong_password_steps`
- **Description:** Initializes a password and returns the number of steps involved in the process.

### Example 1

- **Request:**
  ```json
  {
    "init_password": "aA1"
  }
  ```
- **Response:**
  ```json
  {
    "num_of_steps": 3
  }
  ```

### Example 2

- **Request:**

  ```json
  {
    "init_password": "1445D1cd"
  }
  ```

- **Response:**
  ```json
  {
    "num_of_steps": 0
  }
  ```
