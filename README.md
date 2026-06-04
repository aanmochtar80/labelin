# Labelin

Aplikasi Web Production-Ready untuk mencetak label undangan otomatis dari data Excel.

## Fitur Utama
* Dashboard interaktif
* Master Data Tamu (CRUD)
* Import Excel (.xlsx, .xls)
* Template Label Kustom
* Generate PDF
* Multi User (Role Based)
* UI/UX Modern dengan Bootstrap 5 + HTMX

## Persyaratan
* Golang 1.24+
* Docker (Opsional)

## Cara Menjalankan (Development)
1. Install dependencies:
   ```bash
   go mod download
   ```
2. Jalankan server:
   ```bash
   go run cmd/server/main.go
   ```
3. Akses melalui browser di `http://localhost:3000`

## Build untuk Production
Gunakan script build yang disediakan:
* **Linux**: `./build_linux.sh`
* **Windows**: `build_windows.bat`

Atau gunakan Docker:
```bash
docker-compose up -d --build
```

## Struktur Proyek
- `cmd/server`: Entry point aplikasi.
- `internal/config`: Konfigurasi DB dan Env.
- `internal/handlers`: HTTP handlers.
- `internal/middleware`: Middleware Auth & RBAC.
- `internal/models`: Skema database GORM.
- `internal/routes`: Routing Fiber.
- `internal/services`: Logika bisnis (Excel, PDF).
- `public/`: Aset statis (CSS, JS).
- `views/`: HTML Templates (HTMX).

## API Endpoints
- `GET /api/guests` - List tamu
- `POST /api/guests` - Tambah tamu
- `DELETE /api/guests/:id` - Hapus tamu
- `POST /api/guests/import` - Import Excel (multipart/form-data)
- `POST /api/guests/print` - Generate PDF label
