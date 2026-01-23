Berikut **contoh `README.md`** yang rapi, profesional, dan siap dipakai untuk project **CRUD Kategori API (Golang, Modular, Pagination)**.

---

```md
# Category API (Golang)

RESTful API sederhana menggunakan **Golang (net/http)** dengan konsep **modular**, menyediakan **CRUD Kategori** dan **pagination**.  
Data disimpan secara **in-memory** dengan **40 dummy data awal**.

Project ini cocok sebagai:
- latihan REST API Golang
- boilerplate CRUD sederhana
- referensi struktur modular Go API

---

## 🚀 Fitur

- CRUD Kategori (Create, Read, Update, Delete)
- Pagination menggunakan query parameter
- Default pagination: **10 data per halaman**
- Struktur project modular
- Tanpa database (in-memory storage)
- Menggunakan standard library Go

---

## 🧱 Struktur Project

```

category-api/
├── main.go
├── go.mod
├── models/
│   └── category.go
├── repositories/
│   └── category_repository.go
├── handlers/
│   └── category_handler.go
├── utils/
│   └── pagination.go

```

---

## 📦 Model

### Category

| Field       | Type   |
|------------|--------|
| id         | int    |
| name       | string |
| description| string |

---

## 🔗 Endpoint API

### 1️⃣ Get All Categories (Pagination)

```

GET /categories

```

**Query Params (optional):**
- `page` → default `1`
- `limit` → default `10`

**Contoh:**
```

GET /categories?page=2&limit=5

````

**Response:**
```json
{
  "page": 2,
  "limit": 5,
  "data": [
    {
      "id": 6,
      "name": "Category F",
      "description": "Description for category"
    }
  ]
}
````

---

### 2️⃣ Get Category By ID

```
GET /categories/{id}
```

**Contoh:**

```
GET /categories/1
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
  "description": "New Description"
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

## 🧪 Dummy Data

* Saat aplikasi dijalankan, otomatis dibuat **40 data kategori dummy**
* Data bersifat **in-memory**, akan reset setiap restart server

---

## ▶️ Cara Menjalankan

### 1. Clone Repository

```bash
git clone https://github.com/username/category-api.git
cd category-api
```

### 2. Init Module (jika belum)

```bash
go mod init category-api
```

### 3. Jalankan Server

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
* JSON API

---

## 📌 Catatan Pengembangan

Project ini **belum menggunakan**:

* Database (MySQL, PostgreSQL, MongoDB)
* Framework (Gin, Echo, Fiber)
* Authentication / Authorization

---

## 🚧 Pengembangan Lanjutan (Opsional)

* Integrasi database
* Pagination metadata (`total`, `last_page`)
* Validation request body
* Middleware (logging, recovery)
* Clean Architecture / Hexagonal
* Docker support

---

## 📄 Lisensi

MIT License

