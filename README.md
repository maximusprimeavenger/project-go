# 🚀 Auth Service (Gin + MongoDB)

---

### 📌 Overview
This is an **Authentication Service** built using **Go (Gin)** and **MongoDB**. It provides user registration, login, and JWT-based authentication.

---

### 🔥 Features
- 👤 **User Authentication**: Register, login, and logout functionality.
- 🔐 **JWT Tokens**: Secure authentication with JWT.
- 📂 **MongoDB Integration**: Stores user data securely.
- 🛠 **REST API**: Fully functional REST endpoints using Gin.

---

### 🛠 Tech Stack
- **Backend**: Go (Gin framework)
- **Database**: MongoDB
- **Authentication**: JWT
- **Web Server**: Gin

---

### 📂 Project Structure
```
auth-service/
│── controllers/
│   ├── controllUser.go
│── database/
│   ├── dbConnection.go
├── helpers/
│   ├── authHelper.go
│   ├── tokenHelper.go
├── middleware/
│   ├── middleware.go
├── models/
│   ├── models.go
│── routes/
│   ├── auth.go
│   ├── user.go
│── .gitignore
│── README.md
│── go.mod
│── go.sum
│── main.go
```

---

### ⚡ Installation & Setup
#### Prerequisites
- Install **Go 1.20+**
- Install **MongoDB**

#### Steps
1. **Clone the repository:**
   ```sh
   git clone https://github.com/maximusprimeavenger/project-go.git
   cd project-go
   ```
2. **Set up the database:**
   - Run MongoDB locally or use a cloud database.
3. **Configure database credentials:**
   - Create a `.env` file and add:
   ```sh
   MONGO_URI=mongodb://localhost:27017
   AUTH_PORT=8080
   KEY=your_secret_key
   IP=any ip
   ```
4. **Build and deploy the project:**
   ```sh
   go mod tidy
   go run main.go
   ```
5. **Start the server and access the API:**
   - URL: `http://localhost:8080`

---

### 📌 API Endpoints
#### Register User
**POST /signup**
```sh
curl -X POST http://localhost:8080/users/signup \
  -H "Content-Type: application/json" \
  -d '{"email": "user@example.com", "password": "securepassword"}'
```
Response:
```json
{
    "InsertedID": "67cafb394827968cbfb325ca"
}
```
#### Login
**POST /login**
```sh
curl -X POST http://localhost:8080/users/login \
  -H "Content-Type: application/json" \
  -d '{"email": "user@example.com", "password": "securepassword"}'
```
Response:
```json
{
    "id": "object_id_example",
    "user_id": "user_id_example",
    "name": "John Doe",
    "email": "example@email.com",
    "phone": "1234567890",
    "password": "hashed_password_example",
    "type": "USER",
    "token": "jwt_access_token_example",
    "refresh_token": "jwt_refresh_token_example",
    "created_at": "yyyy-MM-ddTHH:mm:ssZ",
    "updated_at": "yyyy-MM-ddTHH:mm:ssZ"
}

```

### 📜 License
This project is licensed under the MIT License.

