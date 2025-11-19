# juke-test-restapi

# Kandidat Developer Juke

- Membuat REST API berbasis **Java** + Spring Boot (MVC) | **Golang** boleh
- Mengelola data CRUD sederhana
- Menjalankan aplikasi di Docker

# 🧩 Employee Management REST API

## 📘 Deskripsi Kasus

Proyek ini adalah REST API sederhana untuk mengelola data karyawan di perusahaan **Juke**.
Dapat dibuat menggunakan **Java + Spring Boot (MVC)** atau **Golang**, dan dijalankan di atas **Docker container** (nilai plus).

## 🎯 Tujuan Utama

- Membuat REST API sesuai prinsip MVC.
- Melakukan operasi CRUD pada entity **Employee**.
- Menjalankan aplikasi di Docker.
- (Opsional): Menambahkan validasi, exception handling, dan dokumentasi Swagger.

## 🧱 Spesifikasi Teknis

**Bahasa & Framework:**

- Java 17+ (Spring Boot) atau Golang

**Struktur minimal (MVC):**

```
src/
 ├── controller/
 ├── service/
 ├── repository/
 └── model/
```

**Entity: Employee**
| Field | Type | Keterangan |
|--------|-------|------------|
| id | Long | auto increment |
| name | String | nama lengkap |
| email | String | harus unik |
| position | String | jabatan karyawan |
| salary | Double | gaji karyawan |
| createdAt | LocalDateTime | waktu data dibuat |

## 🔌 Endpoint REST API

| Method | Endpoint            | Deskripsi                      |
| ------ | ------------------- | ------------------------------ |
| GET    | /api/employees      | Menampilkan semua karyawan     |
| GET    | /api/employees/{id} | Menampilkan detail 1 karyawan  |
| POST   | /api/employees      | Menambahkan data karyawan baru |
| PUT    | /api/employees/{id} | Mengubah data karyawan         |
| DELETE | /api/employees/{id} | Menghapus data karyawan        |

## ⚙️ Fungsi Tambahan (Opsional)

- Validasi input (`@Valid`) seperti email wajib dan salary > 0.
- Global error handling (`@ControllerAdvice`).
- Dokumentasi API dengan Swagger (`springdoc-openapi-ui`).
- Logging sederhana (`@Slf4j`).

## 🐳 Deploy di Docker

Aplikasi ini dijalankan menggunakan Docker dengan arsitektur multi-container:

1. **Database (PostgreSQL)** - Menyimpan data karyawan
2. **Migration Service** - Menjalankan migrasi database sekali saat startup
3. **App Service** - Menjalankan REST API dengan hot-reload

### Cara Menjalankan

```bash
# Build dan jalankan semua services
docker-compose up --build

# Atau jalankan di background
docker-compose up -d --build

# Melihat logs
docker-compose logs -f app

# Stop semua services
docker-compose down

# Stop dan hapus volumes (reset database)
docker-compose down -v
```

### Urutan Startup

1. Database service dimulai dan health check
2. Migration service menjalankan migrasi database
3. App service dimulai setelah migrasi selesai

---

## 🧠 Penilaian (Total 100 poin)

| Kriteria                                                | Bobot |
| ------------------------------------------------------- | ----- |
| Struktur kode rapi & sesuai MVC                         | 20    |
| Endpoint CRUD berfungsi dengan benar                    | 30    |
| Validasi input & error handling                         | 5     |
| Dokumentasi Swagger                                     | 5     |
| Dockerfile berfungsi (aplikasi bisa jalan di container) | 20    |
| Bisa menjelaskan apa yang dibuat                        | 20    |

**Total: 100 poin**

---

## 📦 Output yang Diminta

- Source code project (GitHub repo)
- File `Dockerfile` dan `docker-compose.yml`
- Petunjuk cara menjalankan project di `README.md`
