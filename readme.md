# Category API (Golang)

RESTful API sederhana menggunakan **Golang (net/http)** dengan konsep **modular**, menyediakan **CRUD Kategori** dan **pagination**.  
Data disimpan di **PostgreSQL database**.

Project ini cocok sebagai:
- latihan REST API Golang dengan database
- boilerplate CRUD dengan PostgreSQL
- referensi struktur modular Go API

---

## 🚀 Fitur

- CRUD Kategori (Create, Read, Update, Delete)
- CRUD Produk (Create, Read, Update, Delete)
- Transaksi/Checkout untuk kasir
- Produk dengan relasi ke Kategori (foreign key)
- Validasi stok sebelum transaksi
- Auto kurangi stok setelah transaksi berhasil
- Pagination menggunakan query parameter
- Filter produk berdasarkan category_id
- Filter kategori dan produk berdasarkan nama (case-insensitive)
- Default pagination: **10 data per halaman**
- Struktur project modular
- PostgreSQL database integration
- Menggunakan standard library Go
- Environment variable configuration
- Automatic timestamp (created_at, updated_at)

---

## 🧱 Struktur Project

category-api/
├── main.go
├── go.mod
├── .env                  # Environment variables
├── database/
│   └── database.go       # PostgreSQL connection
├── models/
│   ├── categories.go     # Category data model
│   └── products.go       # Product data model
├── repositories/
│   ├── category_repository.go # Category database operations
│   └── product_repository.go   # Product database operations
├── handlers/
│   ├── category_handler.go    # Category HTTP handlers
│   └── product_handler.go     # Product HTTP handlers
├── utils/
│   └── pagination.go     # Pagination utility

---

## 📦 Model

### Category

| Field       | Type     |
|------------|----------|
| id         | int      |
| name       | string   |
| description| string   |
| created_at | time.Time|
| updated_at | time.Time|

### Product

| Field        | Type  |
|-------------|-------|
| id          | int   |
| name        | string|
| price       | int   |
| stock       | int   |
| categories_id| int   |

### Transaction

| Field       | Type     |
|------------|----------|
| id         | int      |
| total_amount| int     |
| status     | string   |
| created_at | time.Time|

### TransactionDetail

| Field        | Type  |
|-------------|-------|
| id          | int   |
| transaction_id| int |
| product_id  | int   |
| quantity    | int   |
| subtotal    | int   |

---

## 🔗 Endpoint API

### 1️⃣ Get All Categories (Pagination)

```
GET /categories
```

**Query Params (optional):**
- `page` → default `1`
- `limit` → default `10`
- `name` → filter by name (case-insensitive search)

**Contoh:**
```
GET /categories?page=2&limit=5
GET /categories?name=electronic
```

**Response:**
```json
{
  "page": 2,
  "limit": 5,
  "data": [
    {
      "id": 6,
      "name": "Category F",
      "description": "Description for category",
      "created_at": "2026-02-02T10:00:00Z",
      "updated_at": "2026-02-02T10:00:00Z"
    }
  ]
}
```

---

### 2️⃣ Get Category By ID

```
GET /categories/{id}
```

**Contoh:**
```
GET /categories/1
```

**Response:**
```json
{
  "id": 1,
  "name": "Category A",
  "description": "Description for category",
  "created_at": "2026-02-02T10:00:00Z",
  "updated_at": "2026-02-02T10:00:00Z"
}
```

---

### 3️⃣ Create Category

```
POST /categories
```

**Request Body:**

```json
{
  "name": "New Category",
  "description": "New Description"
}
```

**Response:**

```json
{
  "id": 41,
  "name": "New Category",
  "description": "New Description",
  "created_at": "2026-02-02T11:00:00Z",
  "updated_at": "2026-02-02T11:00:00Z"
}
```

---

### 4️⃣ Update Category

```
PUT /categories/{id}
```

**Request Body:**

```json
{
  "name": "Updated Category",
  "description": "Updated Description"
}
```

**Response:**
```json
{
  "id": 1,
  "name": "Updated Category",
  "description": "Updated Description",
  "created_at": "2026-02-02T10:00:00Z",
  "updated_at": "2026-02-02T12:00:00Z"
}
```

---

### 5️⃣ Delete Category

```
DELETE /categories/{id}
```

**Response:**
```
204 No Content
```

---

## 📦 Product Endpoints

### 6️⃣ Get All Products (Pagination)

```
GET /products
```

**Query Params (optional):**
- `page` → default `1`
- `limit` → default `10`
- `category_id` → filter by category
- `name` → filter by name (case-insensitive search)

**Contoh:**
```
GET /products?page=2&limit=5
GET /products?category_id=1
GET /products?name=laptop
```

**Response:**
```json
{
  "page": 2,
  "limit": 5,
  "data": [
    {
      "id": 6,
      "name": "Product F",
      "price": 100000,
      "stock": 50,
      "categories_id": 1
    }
  ]
}
```

---

### 7️⃣ Get Product By ID

```
GET /products/{id}
```

**Contoh:**
```
GET /products/1
```

**Response:**
```json
{
  "id": 1,
  "name": "Product A",
  "price": 50000,
  "stock": 100,
  "categories_id": 1
}
```

---

### 8️⃣ Create Product

```
POST /products
```

**Request Body:**

```json
{
  "name": "New Product",
  "price": 75000,
  "stock": 20,
  "categories_id": 1
}
```

**Response:**

```json
{
  "id": 1,
  "name": "New Product",
  "price": 75000,
  "stock": 20,
  "categories_id": 1
}
```

---

### 9️⃣ Update Product

```
PUT /products/{id}
```

**Request Body:**

```json
{
  "name": "Updated Product",
  "price": 80000,
  "stock": 15,
  "categories_id": 2
}
```

**Response:**

```json
{
  "id": 1,
  "name": "Updated Product",
  "price": 80000,
  "stock": 15,
  "categories_id": 2
}
```

---

### 🔟 Delete Product

```
DELETE /products/{id}
```

**Response:**
```
204 No Content
```

---

## 💰 Transaction Endpoints

### 1️⃣1️⃣ Create Transaction (Checkout)

```
POST /transactions
```

**Request Body:**

```json
{
  "items": [
    {
      "product_id": 1,
      "quantity": 2
    },
    {
      "product_id": 3,
      "quantity": 1
    }
  ]
}
```

**Response:**

```json
{
  "id": 1,
  "total_amount": 225000,
  "status": "completed",
  "details": [
    {
      "id": 1,
      "transaction_id": 1,
      "product_id": 1,
      "quantity": 2,
      "subtotal": 100000
    },
    {
      "id": 2,
      "transaction_id": 1,
      "product_id": 3,
      "quantity": 1,
      "subtotal": 125000
    }
  ]
}
```

**Catatan:**
- Stok produk akan otomatis dikurangi setelah transaksi berhasil
- Transaksi akan gagal jika stok tidak mencukupi atau produk tidak ditemukan
- Jika terjadi error, response akan berisi detail product_id, requested quantity, dan available stock

**Error Response Examples:**

1. Produk tidak ditemukan:
```json
{
  "error": "product not found",
  "product_id": 999
}
```

2. Stok tidak mencukupi:
```json
{
  "error": "insufficient stock for product",
  "product_id": 1,
  "requested": 10,
  "available": 5
}
```

---

### 1️⃣2️⃣ Get All Transactions (Pagination)

```
GET /transactions
```

**Query Params (optional):**
- `page` → default `1`
- `limit` → default `10`

**Contoh:**
```
GET /transactions?page=1&limit=10
```

**Response:**
```json
{
  "page": 1,
  "limit": 10,
  "data": [
    {
      "id": 1,
      "total_amount": 225000,
      "status": "completed",
      "created_at": "2026-02-10T10:00:00Z"
    }
  ]
}
```

---

### 1️⃣3️⃣ Get Transaction By ID (With Details)

```
GET /transactions/{id}
```

**Contoh:**
```
GET /transactions/1
```

**Response:**
```json
{
  "id": 1,
  "total_amount": 225000,
  "status": "completed",
  "created_at": "2026-02-10T10:00:00Z",
  "details": [
    {
      "id": 1,
      "transaction_id": 1,
      "product_id": 1,
      "quantity": 2,
      "subtotal": 100000
    },
    {
      "id": 2,
      "transaction_id": 1,
      "product_id": 3,
      "quantity": 1,
      "subtotal": 125000
    }
  ]
}
```

---

## 🗄 Database Setup

### Prerequisites

- PostgreSQL installed and running
- Create a database for the project

### Create Tables

Run these SQL commands to create all required tables:

```sql
CREATE TABLE categories (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE products (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    price INT NOT NULL,
    stock INT NOT NULL,
    categories_id INT NOT NULL,
    FOREIGN KEY (categories_id) REFERENCES categories(id) ON DELETE CASCADE
);

CREATE TABLE transactions (
    id SERIAL PRIMARY KEY,
    total_amount INT NOT NULL,
    status VARCHAR(20) DEFAULT 'completed',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE transaction_details (
    id SERIAL PRIMARY KEY,
    transaction_id INT REFERENCES transactions(id) ON DELETE CASCADE,
    product_id INT REFERENCES products(id),
    quantity INT NOT NULL,
    subtotal INT NOT NULL
);
```

---

## 🔧 Environment Variables

Create a `.env` file in the project root:

```env
PORT=8080
DB_CONN=postgres://username:password@localhost:5432/database_name?sslmode=disable
```

**Note:** Replace `username`, `password`, and `database_name` with your actual PostgreSQL credentials.

---

## ▶️ Cara Menjalankan

### 1. Clone Repository

```bash
git clone https://github.com/username/categories-api.git
cd categories-api
```

### 2. Install Dependencies

```bash
go mod download
```

### 3. Setup Environment

```bash
cp .env.example .env
# Edit .env with your database credentials
```

### 4. Create Database Table

Run the SQL command from the "Database Setup" section above.

### 5. Jalankan Server

```bash
go run main.go
```

Server akan berjalan di:

```
http://localhost:8080
```

---

## 🛠 Teknologi

* Go (Golang)
* net/http (standard library)
* PostgreSQL
* github.com/lib/pq (PostgreSQL driver)
* github.com/spf13/viper (Configuration management)
* JSON API

---

## 📌 Catatan Pengembangan

Project ini **belum menggunakan**:

* Framework (Gin, Echo, Fiber)
* Authentication / Authorization
* ORM (GORM, sqlx)

---

## 🚧 Pengembangan Lanjutan (Opsional)

* Pagination metadata (`total`, `last_page`)
* Validation request body
* Middleware (logging, recovery, CORS)
* Clean Architecture / Hexagonal
* Docker support
* Unit testing & integration testing
* API Documentation (Swagger/OpenAPI)
* Authentication with JWT
* Rate limiting

---

## 📄 Lisensi

MIT License
