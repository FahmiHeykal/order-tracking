A real-time order tracking system built with Go, PostgreSQL, and WebSocket for instant status updates.

📌 Features
User Management

Registration & JWT Authentication

Role-based access (user, admin, driver)

Order Tracking

Full CRUD operations

Status flow: Pending → Processing → Shipped → Completed

Real-time Updates

WebSocket notifications

Offline status recovery via REST API

Order History

Complete audit trail of status changes

🚀 Getting Started
Prerequisites
Go 1.21+

PostgreSQL 14+

Git (optional)

Installation
Clone the repository:

bash
git clone https://github.com/yourusername/order-tracking.git
cd order-tracking
Set up environment variables:

bash
cp .env.example .env
Install dependencies:

bash
go mod tidy
Database Setup
sql
CREATE DATABASE order_tracking;
\c order_tracking
\i migrations/001_init.sql
⚙️ Configuration
Edit .env file:

env
DB_DSN=host=localhost user=postgres password=yourpassword dbname=order_tracking port=5432 sslmode=disable
JWT_SECRET=verysecretkey
🏗️ Project Structure
text
order-tracking/
├── config/
├── internal/
│   ├── handler/
│   ├── middleware/
│   ├── model/
│   ├── repository/
│   ├── service/
│   └── websocket/
├── migrations/
├── pkg/
│   ├── response/
│   └── utils/
├── go.mod
├── go.sum
└── main.go
🌐 API Documentation
Authentication
Method	Endpoint	Body	Description
POST	/api/register	{name, email, password}	Register new user
POST	/api/login	{email, password}	Login and get JWT
Orders
Method	Endpoint	Role	Body	Description
POST	/api/orders	user	{description}	Create new order
GET	/api/orders	user	-	Get user's orders
GET	/api/orders/:id	user	-	Get order details
PUT	/api/orders/:id/status	admin/driver	{status}	Update order status
GET	/api/admin/orders	admin/driver	-	Get all orders
WebSocket
text
ws://localhost:8080/ws/orders/:id
Headers:

text
Authorization: Bearer <JWT_TOKEN>
Message Format:

json
{
  "order_id": "1",
  "status": "Shipped",
  "updated_at": "2025-07-28T10:00:00Z"
}
🧪 Testing
Manual Testing
Register a user:

bash
curl -X POST http://localhost:8080/api/register \
  -H "Content-Type: application/json" \
  -d '{"name":"Test User","email":"test@example.com","password":"password123"}'
Login to get JWT:

bash
curl -X POST http://localhost:8080/api/login \
  -H "Content-Type: application/json" \
  -d '{"email":"test@example.com","password":"password123"}'
Create an order:

bash
curl -X POST http://localhost:8080/api/orders \
  -H "Authorization: Bearer <JWT_TOKEN>" \
  -H "Content-Type: application/json" \
  -d '{"description":"Nasi Goreng Special"}'
Automated Tests
bash
go test ./...
🐛 Troubleshooting
Issue	Solution
Invalid JWT Token	Check token expiration and secret key
WebSocket not updating	Verify Hub broadcast logs
Database connection failed	Check PostgreSQL service status
🚢 Deployment
Local Build
bash
go build -o order-tracking
./order-tracking
