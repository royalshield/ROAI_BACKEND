
# 🎯 Go Event Management Backend

This is a backend API built with **Go (Golang)** using **Gin**, **GORM**, **MySQL**, and **JWT-based authentication**. It supports user authentication via wallet addresses, event creation and registration, and notification broadcasting. The system implements **role-based access control** for users and admins.

---

## 🧰 Tech Stack

| Layer         | Technology           |
|---------------|----------------------|
| Language      | Go (Golang)          |
| Web Framework | Gin                  |
| ORM           | GORM                 |
| Database      | MySQL                |
| Auth          | JWT (Role-based)     |
| Config        | Dotenv               |

---

## 📁 Project Structure

```
go-event-backend/
├── cmd/               # Main application entry
├── config/            # Database and environment config
├── controllers/       # Business logic and route handlers
├── middlewares/       # JWT authentication and role middleware
├── models/            # GORM models
├── routes/            # All route definitions
├── utils/             # JWT utilities
├── .env               # Environment variables
└── README.md
```

---

## 📦 Setup Instructions

### 1. Clone the repository

```bash
git clone https://github.com/your-username/go-event-backend.git
cd go-event-backend
```

### 2. Create `.env` file

```env
PORT=8080
DB_USER=root
DB_PASSWORD=your_password
DB_NAME=eventdb
DB_HOST=127.0.0.1
DB_PORT=3306
JWT_SECRET=supersecretkey
```

### 3. Install dependencies

```bash
go mod tidy
```

### 4. Run the application

```bash
go run cmd/main.go
```

---

## 🔐 Authentication

Use the wallet address to login and receive a JWT token:

### ➤ `POST /api/auth/wallet`
```json
{
  "wallet": "0xabc123..."
}
```

> ✅ Response includes a `token` to use in Authorization header for protected routes.

---

## 🧑 Roles

- **User**: Can login and register for events.
- **Admin**: Can create/delete events, create notifications, and access full user/event data.

---

## 🔗 API Endpoints

### ✅ Authentication

| Method | Endpoint              | Access  | Description             |
|--------|-----------------------|---------|-------------------------|
| POST   | `/api/auth/wallet`    | Public  | Authenticate with wallet|

---

### 👥 Users

| Method | Endpoint       | Access | Description    |
|--------|----------------|--------|----------------|
| GET    | `/api/users`   | Admin  | Get all users  |

---

### 🔔 Notifications

| Method | Endpoint            | Access     | Description            |
|--------|---------------------|------------|------------------------|
| POST   | `/api/notifications`| Admin      | Create a notification  |
| GET    | `/api/notifications`| Admin/User | Get all notifications  |

---

### 📅 Events

| Method | Endpoint                     | Access     | Description                    |
|--------|------------------------------|------------|--------------------------------|
| POST   | `/api/events`                | Admin      | Create an event                |
| GET    | `/api/events`                | Admin/User | Get all events                 |
| POST   | `/api/events/:id/register`   | User       | Register user for an event     |
| POST   | `/api/events/:id/delete`     | Admin      | Delete an event                |
| POST   | `/api/events/:id/attendees`  | Admin      | Get event attendees by wallet  |

---

## 🗃️ Database Schema

### 📌 Users

| Field     | Type       |
|-----------|------------|
| id        | uint       |
| wallet    | string     |
| role      | enum       |
| created_at| timestamp  |
| updated_at| timestamp  |

### 📌 Notifications

| Field       | Type     |
|-------------|----------|
| id          | uint     |
| title       | string   |
| description | string   |
| created_at  | timestamp|
| updated_at  | timestamp|

### 📌 Events

| Field             | Type     |
|-------------------|----------|
| id                | uint     |
| picture           | string   |
| title             | string   |
| description       | text     |
| place             | string   |
| attendees_expected| int      |
| members_limit     | int      |
| date              | datetime |
| created_at        | timestamp|
| updated_at        | timestamp|

---

## 🧪 Testing

Use the provided Postman Collection to test all APIs.

📥 [Download Postman Collection](./go-event-backend.postman_collection.json)

---

## 🚀 Future Enhancements

- Docker support
- Email/Push notifications
- Admin dashboard UI
- Swagger auto-documentation

---

## 📜 License

MIT © 2025 Go Event Backend
