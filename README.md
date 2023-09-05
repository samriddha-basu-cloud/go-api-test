# User Management API

![Go](https://img.shields.io/badge/Go-1.16-blue?logo=go)
![MongoDB](https://img.shields.io/badge/MongoDB-4.4-green?logo=mongodb)

A simple User Management API built in Go with MongoDB integration.
---

## Overview

This project provides a RESTful API for managing user data in a MongoDB database. It includes basic CRUD operations (Create, Read, Update, Delete) for user records. Use this as a foundation for similar projects.

## Features

- Create users with random IDs
- Retrieve user details by ID
- Get a list of all users
- Update user information
- Delete users by ID
- Delete all users

## How to Use

1. **Clone the Repository**

   ```bash
   git clone https://github.com/samriddha-basu-cloud/go-api-test.git
   ```

2. **Install Dependencies**

   ```bash
   go mod tidy
   ```

3. **Set Up MongoDB**

   - Install and run MongoDB locally or provide a connection URI in `main.go`.

4. **Run the Application**

   ```bash
   go run main.go
   ```

5. **API Endpoints**

   - Create User: `POST /users`
   - Get User by ID: `GET /users/:id`
   - Get All Users: `GET /users`
   - Update User by ID: `PUT /users/:id`
   - Delete User by ID: `DELETE /users/:id`
   - Delete All Users: `DELETE /users`

6. **Modify for Your Project**

   - Customize the `User` struct and MongoDB configuration for your data model.
   - Add more routes and functionalities as needed.
   - Adapt the project for other database systems.

## Contributing

Contributions are welcome! Feel free to open issues and submit pull requests.

## License

This project is open-source and free to use, with no specific license. Use it as you see fit for your projects.

---
