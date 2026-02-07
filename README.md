# 📚 Bookstore Management Backend

![Go Version](https://img.shields.io/badge/Go-1.25.6-00ADD8?style=for-the-flat&logo=go&logoColor=white)
![MySQL](https://img.shields.io/badge/MySQL-4479A1?style=for-the-flat&logo=mysql&logoColor=white)
![GORM](https://img.shields.io/badge/ORM-GORM-blue?style=for-the-flat)
![License](https://img.shields.io/badge/License-MIT-green?style=for-the-flat)
![Build](https://img.shields.io/badge/Build-Passing-brightgreen?style=for-the-flat)

---

A robust and lightweight RESTful API for managing a bookstore, built using **Go**, **GORM**, and **MySQL**. This project demonstrates clean architecture and efficient database management in the Go ecosystem.

## 🛠 Tech Stack
* **Language:** Go (Golang)
* **Framework:** Gorilla Mux (Router)
* **Database:** MySQL
* **ORM:** GORM

## 🚀 Features
- [x] **Create Book**: Add new book records to the database.
- [x] **Get All Books**: Retrieve a complete list of books in JSON format.
- [x] **Get Book by ID**: Fetch detailed information for a specific book.
- [x] **Update Book**: Modify existing book details.
- [x] **Delete Book**: Remove records using GORM's efficient deletion (supports Soft Delete).

## 💻 Project Setup

1.  **Database Configuration**: Create a MySQL database named `simplerest`.
2.  **Environment Setup**: Update your MySQL credentials (username and password) in `pkg/config/app.go`.
3.  **Install Dependencies**:
    ```bash
    go mod tidy
    ```
4.  **Run the Server**:
    ```bash
    go run cmd/main.go
    ```
    The server will start at `http://localhost:9010/book/`.
