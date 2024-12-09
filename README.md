# **Product Management System**

## **Overview**
The **Product Management System** is a backend service built using Go for managing products with asynchronous image processing, caching, and scalable architecture. This repository provides REST APIs for adding and retrieving product details while efficiently handling image processing and caching to enhance performance.

---

## **Features**
1. **RESTful APIs**:
   - **POST /products**: Create a new product.
   - **GET /products/{id}**: Retrieve product details by ID.
   - **GET /products**: Get products by user with filters.
2. **Asynchronous Image Processing**:
   - Uses RabbitMQ to queue image processing tasks for scalability.
3. **Caching**:
   - Redis caching is used to speed up frequently accessed data.
4. **Database**:
   - PostgreSQL stores product data in a relational format.
5. **Error Handling**:
   - Implements robust error handling for database, queue, and cache operations.

---

## **Setup Instructions**

### **Prerequisites**
1. **Dependencies**:
   - Go version 1.18 or later.
   - PostgreSQL.
   - RabbitMQ.
   - Redis.

2. **Clone the Repository**:
   git clone https://github.com/rp5454/Product-Management-System.git
   cd Product-Management-System

Set Environment Variables: Create a .env file in the root directory:
DB_USER=your_db_user
DB_PASSWORD=your_db_password
DB_NAME=your_db_name
DB_HOST=localhost
DB_PORT=5432

REDIS_HOST=localhost:6379

RABBITMQ_URL=amqp://guest:guest@localhost:5672/

Set Up the Database:
Use the schema provided in config/schema.sql:
psql -U your_db_user -d your_db_name -f config/schema.sql

Start RabbitMQ and Redis: Ensure both RabbitMQ and Redis are running.

Install Dependencies:
go mod tidy

Run the Application:
go run main.go

Project Structure:
Product-Management-System/
├── main.go                     # Main application entry point

├── router/                     # API routing logic

│   ├── router.go               # Routes definitions

├── handlers/                   # API handlers

│   ├── product.go              # Handlers for product APIs

├── services/                   # Core business logic and integrations

│   ├── redis.go                # Redis integration

│   ├── rabbitmq.go             # RabbitMQ integration

├── config/                     # Configuration files

│   ├── schema.sql              # Database schema

│   ├── env.go                  # Environment variable loader

├── models/                     # Data models

│   ├── product.go              # Product model

├── utils/                      # Utility functions

│   ├── logger.go               # Centralized logging utility

├── go.mod                      # Go module definition

├── README.md                   # Project documentation


API Documentation
POST /products
Description: Creates a new product and queues its images for processing.
Request Body:
{
  "name": "Product Name",
  "description": "Product Description",
  "price": 100.50,
  "image_url": ["https://example.com/image1.jpg", "https://example.com/image2.jpg"]
}

GET /products/{id}
Description: Retrieves product details by ID. Attempts to fetch from Redis cache first; falls back to PostgreSQL if not cached.
{
  "id": 1,
  "name": "Product Name",
  "description": "Product Description",
  "price": 100.50,
  "image_url": ["https://example.com/image1.jpg"]
}


Design Choices
Modular Structure:
The codebase is organized into separate modules (handlers, services, config) for better maintainability and scalability.
Asynchronous Processing:
RabbitMQ ensures non-blocking operations for image processing tasks.
Caching:
Redis improves performance by caching frequently accessed data.
Error Handling:
Comprehensive error handling in each component ensures system stability.

