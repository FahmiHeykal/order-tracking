# Real-time Order Tracking System
Backend sistem pelacakan status pesanan real-time seperti food delivery. Dibangun dengan Go, Gin, GORM, WebSocket Gorilla, dan PostgreSQL.

## Fitur Utama
### User Management
- Register
- Login (JWT)
- Role: user, admin/driver
- Middleware JWT

### CRUD Pesanan
- User buat pesanan
- User lihat daftar pesanan
- Admin/driver lihat semua pesanan
- Admin/driver update status

### Status Pesanan
- Flow: Pending → Diproses → Dikirim → Selesai
- Status update otomatis broadcast via WebSocket

### WebSocket Tracking
- Client subscribe ke order_id
- Server broadcast update status ke semua subscriber
- Jika offline, ambil status terbaru via REST

### Riwayat Pesanan
- Endpoint REST untuk histori status

---

## Tech Stack
- Go (Gin)
- PostgreSQL
- GORM
- Gorilla/WebSocket
- JWT (golang-jwt)
- godotenv

---

```
Struktur Folder
order-tracking
├── internal/
│   ├── model/
│   ├── dto/
│   ├── repository/
│   ├── service/
│   ├── handler/
│   ├── websocket/
│   └── middleware/
├── pkg/
│   ├── utils/
│   └── response/
├── migrations/
├── config.go
├── main.go
├── .env
├── go.mod
└── README.md
```
---

### Setup Database
1. Buat database PostgreSQL:
psql -U postgres -c "CREATE DATABASE order_tracking;"
2. Jalankan file SQL migrasi di migrations/001_init.sql.

### Cara Menjalankan
Clone project:
`git clone https://github.com/username/real-time-order-tracking.git`
cd real-time-order-tracking

### Buat file .env:
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=your_password
DB_NAME=order_tracking
JWT_SECRET=your_secret
APP_PORT=8080

### Install dependency:
`go mod tidy`

### Jalankan server:
`go run main.go`

### Server jalan di:
`http://localhost:8080`

---

## Endpoint REST API
#### Auth
Register
POST /api/register
Request:
{
"name": "John",
"email": "john@example.com",
"password": "secret"
}

#### Login
POST /api/login
Request:
{
"email": "john@example.com",
"password": "secret"
}
Response:
{
"status": "success",
"data": { "token": "jwt_token" }
}

---

### Orders
#### Buat Pesanan
POST /api/orders
Authorization: Bearer <token>
Body:
{ "items": "Nasi Goreng, Es Teh" }

#### Lihat Pesanan User
GET /api/orders
Authorization: Bearer <token>

#### Detail Pesanan
GET /api/orders/:id
Authorization: Bearer <token>

#### Update Status Pesanan (Admin/Driver)
PUT /api/orders/:id/status
Authorization: Bearer <token>
Body:
{ "status": "Dikirim" }

#### Histori Pesanan
GET /api/v1/orders/history
Authorization: Bearer <token>

---

## WebSocket
### Endpoint:
GET /ws/orders/:id
Contoh:
ws://localhost:8080/ws/orders/1
Jika status berubah, server broadcast:

{
"order_id": "1",
"status": "Dikirim",
"updated_at": "2025-07-28T10:00:00Z"
}

### Format Response
Sukses:

{
"status": "success",
"data": { ... }
}

Error:

{
"status": "error",
"error": "Unauthorized"
}

---

## Testing
1. Register & login → dapatkan token JWT
2. Buat pesanan
3. Update status sebagai admin/driver
4. Pastikan client WebSocket menerima update real-time
5. Jika client offline → login lagi → ambil status terbaru via REST

