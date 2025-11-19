# Employee Management REST API (Golang)

## Prasyarat

- Docker Desktop
- Docker Compose

## Cara Menjalankan Project

### 1. Setup Environment

Buat file `.env` di root project:

```env
DB_HOST=localhost
DB_PORT=5433
DB_USER=postgres
DB_PASS=postgres
DB_NAME=juke_test_db
```

### 2. Jalankan Docker

```bash
docker-compose up --build
```

### 3. Akses API

- **Base URL:** `http://localhost:8080/api`
- **Swagger UI:** `http://localhost:8080/swagger/index.html`

## API Endpoints

| Method | Endpoint           | Deskripsi                  |
| ------ | ------------------ | -------------------------- |
| GET    | /api/employees     | Mendapatkan semua karyawan |
| GET    | /api/employees/:id | Mendapatkan karyawan by ID |
| POST   | /api/employees     | Menambah karyawan baru     |
| PUT    | /api/employees/:id | Mengupdate data karyawan   |
| DELETE | /api/employees/:id | Menghapus karyawan         |

## Contoh Request Body (Postman)

### 1. GET All Employees

```
GET http://localhost:8080/api/employees
```

**Body:** Tidak perlu

**Response:**

```json
{
  "data": [
    {
      "id": 1,
      "name": "John Doe",
      "email": "john@example.com",
      "position": "Software Engineer",
      "salary": 75000.0,
      "created_at": "2025-11-19T10:15:30Z"
    }
  ]
}
```

---

### 2. GET Employee by ID

```
GET http://localhost:8080/api/employees/1
```

**Body:** Tidak perlu

**Response:**

```json
{
  "data": {
    "id": 1,
    "name": "John Doe",
    "email": "john@example.com",
    "position": "Software Engineer",
    "salary": 75000.0,
    "created_at": "2025-11-19T10:15:30Z"
  }
}
```

---

### 3. POST Create Employee

```
POST http://localhost:8080/api/employees
Content-Type: application/json
```

**Request Body:**

```json
{
  "name": "Alice Johnson",
  "email": "alice@example.com",
  "position": "Backend Developer",
  "salary": 80000.5
}
```

**Response:**

```json
{
  "data": {
    "id": 2,
    "name": "Alice Johnson",
    "email": "alice@example.com",
    "position": "Backend Developer",
    "salary": 80000.5,
    "created_at": "2025-11-19T10:20:15Z"
  }
}
```

---

### 4. PUT Update Employee

```
PUT http://localhost:8080/api/employees/2
Content-Type: application/json
```

**Request Body:**

```json
{
  "name": "Alice Johnson Updated",
  "email": "alice.updated@example.com",
  "position": "Senior Backend Developer",
  "salary": 95000.0
}
```

**Response:**

```json
{
  "data": {
    "id": 2,
    "name": "Alice Johnson Updated",
    "email": "alice.updated@example.com",
    "position": "Senior Backend Developer",
    "salary": 95000.0,
    "created_at": "2025-11-19T10:20:15Z"
  }
}
```

---

### 5. DELETE Employee

```
DELETE http://localhost:8080/api/employees/2
```

**Body:** Tidak perlu

**Response:**

```json
{
  "message": "Employee deleted successfully"
}
```

---

## Contoh Error Response

**404 Not Found:**

```json
{
  "error": true,
  "message": "Employee not found",
  "code": 404
}
```

**400 Bad Request (Validation):**

```json
{
  "error": true,
  "message": "email is required",
  "code": 400
}
```

**400 Duplicate Email:**

```json
{
  "error": true,
  "message": "Email already exists",
  "code": 400
}
```

---

## Stop Docker

```bash
# Stop services
docker-compose down

# Stop dan hapus database
docker-compose down -v
```
