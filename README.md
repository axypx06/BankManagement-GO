**Bank Management System**

- This is a simple Bank Management System HTTP (REST) API implemented in Go using the Gofr framework. The system provides basic functionality to manage customer accounts, including creating accounts, retrieving account information, updating customer details, and handling financial transactions such as depositing and withdrawing money.
- I have used docker to connect mysql and redis and used postman for testing API requests.
- Unit Test Coverage > 60%

![image](https://github.com/axypx06/BankManagement-GO/assets/110666919/b01c06e2-0e02-4a56-b15e-b0d302fb323a)



Prerequisites
- Golang installed
- Docker installed
- Postman for testing API requests
  
Features
- Create Customer Account: Admin can Add a new customer to the system.
- Retrieve Customer Information: Admin can Get details of all customers or a specific customer by ID.
- Update Customer Details: Admin can Modify the name of a customer.
- Deposit Money: customer add money to the account balance.
- Withdraw Money: customer withdraws the money from account balance, with checks for negative balance and sufficient funds.

API Endpoints
- GET /greet
Returns a greeting message stored in Redis.
![image](https://github.com/axypx06/BankManagement-GO/assets/110666919/39237caf-d92b-4f88-83f2-a4ff2f0f774c)


- GET /admin/viewCustomer
Returns information about all customers.
![image](https://github.com/axypx06/BankManagement-GO/assets/110666919/29e1da1f-5078-4400-918a-12a64fd446dc)


- GET /admin/viewCustomer/{id}
Returns information about a specific customer based on the provided ID.
![image](https://github.com/axypx06/BankManagement-GO/assets/110666919/dcaa14ee-c335-4a3d-aa2b-54e83a1edf55)


- POST /admin/addCustomer/{name}
Creates a new customer account with the given name.
![image](https://github.com/axypx06/BankManagement-GO/assets/110666919/babcc8d3-5912-4e12-8a52-848589184735)


- DELETE /admin/deleteCustomer/{id}
Deletes a customer account based on the provided ID.
![image](https://github.com/axypx06/BankManagement-GO/assets/110666919/4ee70804-6a99-47e1-85af-bbb060aa67fa)



- PUT /admin/update/{id}
Updates the name of a customer based on the provided ID.
![image](https://github.com/axypx06/BankManagement-GO/assets/110666919/4a7b2afe-d2f0-445b-ae62-d68a136b9787)


- PUT /customer/{id}/add
Adds money to a customer's account. Requires a JSON payload with the balance field.
![image](https://github.com/axypx06/BankManagement-GO/assets/110666919/fe5ad649-42c4-4c0f-a73c-b8f51833a7aa)


- PUT /customer/{id}/withdraw
Withdraws money from a customer's account. Requires a JSON payload with the balance field.
![image](https://github.com/axypx06/BankManagement-GO/assets/110666919/5bca2977-a417-45fb-bce6-89a56b861ada)


Usage

- Clone the repository:
git clone [https://github.com/axypx06/BankManagement-GO.git](https://github.com/axypx06/BankManagement-GO)

- Navigate to the project directory:
cd BankManagement

- Run the application:
go run main.go
Access the API at http://localhost:8000 , it can be over-ridden through configs

Dependencies
Gofr: https://gofr.dev

Database Schema : 
The application uses a MySQL database with the following schema:

sql
Copy code
CREATE TABLE customers (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    balance FLOAT DEFAULT 0
);
![image](https://github.com/axypx06/BankManagement-GO/assets/110666919/c458acbc-0421-4d58-ac2a-0705d5d12f34)


**Notes**:  

This project is a simple demonstration and may not cover all aspects of a production-grade banking system.
Future enhancements to the project may include:
- User Authentication: Implementing user authentication and middlewares for secure customer data access.
- Enhanced Transactions: Extending transaction capabilities to include more complex financial operations.
  







