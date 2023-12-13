**Bank Management System**

This is a simple Bank Management System implemented in Go using the Gofr framework. The system provides basic functionality to manage customer accounts, including creating accounts, retrieving account information, updating customer details, and handling financial transactions such as depositing and withdrawing money.

Prerequisites
- Golang installed
- Docker installed
- Postman for testing API requests
  
Features
- Create Customer Account: Add a new customer to the system.
- Retrieve Customer Information: Get details of all customers or a specific customer by ID.
- Update Customer Details: Modify the name of a customer.
- Deposit Money: Increase the account balance.
- Withdraw Money: Decrease the account balance, with checks for negative balance and sufficient funds.

API Endpoints
- GET /greet
Returns a greeting message stored in Redis.

- GET /customer
Returns information about all customers.

- GET /customer/{id}
Returns information about a specific customer based on the provided ID.

- POST /customer/{name}
Creates a new customer account with the given name.

- DELETE /customer/{id}
Deletes a customer account based on the provided ID.

- PUT /customer/{id}
Updates the name of a customer based on the provided ID.

- PUT /customer/{id}/add
Adds money to a customer's account. Requires a JSON payload with the balance field.

- PUT /customer/{id}/withdraw
Withdraws money from a customer's account. Requires a JSON payload with the balance field.

Usage

- Clone the repository:
git clone https://github.com/your-username/BankManagement.git

- Navigate to the project directory:
cd BankManagement

- Run the application:
go run main.go
Access the API at http://localhost:8000

Dependencies
Gofr: https://gofr.dev
Database Schema
The application uses a MySQL database with the following schema:

sql
Copy code
CREATE TABLE customers (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    balance FLOAT DEFAULT 0
);

Notes
This project is a simple demonstration and may not cover all aspects of a production-grade banking system.
I have used docker to connect mysql and redis and used postman for testing API requests.
Feel free to contribute and enhance the features of this Bank Management System. If you encounter any issues or have suggestions, please open an issue on GitHub.





