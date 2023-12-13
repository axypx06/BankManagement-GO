package main

import (
	"encoding/json"
	"errors"

	"gofr.dev/pkg/gofr"
)

type Customer struct {
	ID      int     `json:"id"`
	Name    string  `json:"name"`
	Balance float64 `json:"balance"`
}

func main() {
	// initialise gofr object
	app := gofr.New()
	app.GET("/greet", HelloHandler)
	app.GET("/customer", GetCustomer)
	app.GET("/customer/{id}", GetCustomerByID)
	app.POST("/customer/{name}", CreateCustomer)
	app.DELETE("/customer/{id}", DeleteCustomer)
	app.PUT("/customer/{id}", UpdateCustomer)
	app.PUT("/customer/{id}/add", AddMoney)
	app.PUT("/customer/{id}/withdraw", WithdrawMoney)

	// Starts the server, it will listen on the default port 8000.
	// it can be over-ridden through configs
	app.Start()
}
func HelloHandler(ctx *gofr.Context) (interface{}, error) {
	value, err := ctx.Redis.Get(ctx.Context, "greeting").Result()

	return value, err
}
func GetCustomer(ctx *gofr.Context) (interface{}, error) {
	var customers []Customer

	// Getting the customer from the database using SQL
	rows, err := ctx.DB().QueryContext(ctx, "SELECT * FROM customers")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var customer Customer
		if err := rows.Scan(&customer.ID, &customer.Name, &customer.Balance); err != nil {
			return nil, err
		}

		customers = append(customers, customer)
	}

	return customers, nil
}
func GetCustomerByID(ctx *gofr.Context) (interface{}, error) {
	customerID := ctx.PathParam("id")

	row := ctx.DB().QueryRowContext(ctx, "SELECT * FROM customers WHERE id = ?", customerID)

	var customer Customer
	if err := row.Scan(&customer.ID, &customer.Name, &customer.Balance); err != nil {
		return nil, err
	}

	return customer, nil
}

func CreateCustomer(ctx *gofr.Context) (interface{}, error) {
	name := ctx.PathParam("name")

	_, err := ctx.DB().ExecContext(ctx, "INSERT INTO customers (name) VALUES (?)", name)

	return nil, err
}

func DeleteCustomer(ctx *gofr.Context) (interface{}, error) {
	customerID := ctx.PathParam("id")

	_, err := ctx.DB().ExecContext(ctx, "DELETE FROM customers WHERE id = ?", customerID)

	return nil, err
}

func UpdateCustomer(ctx *gofr.Context) (interface{}, error) {

	customerID := ctx.PathParam("id")

	var updateData struct {
		Name string `json:"name"`
	}
	if err := json.NewDecoder(ctx.Request().Body).Decode(&updateData); err != nil {
		return nil, err
	}

	_, err := ctx.DB().ExecContext(ctx, "UPDATE customers SET name = ? WHERE id = ?", updateData.Name, customerID)

	return nil, err
}

func AddMoney(ctx *gofr.Context) (interface{}, error) {

	customerID := ctx.PathParam("id")

	var updateBalance struct {
		Balance float64 `json:"balance"`
	}
	if err := json.NewDecoder(ctx.Request().Body).Decode(&updateBalance); err != nil {
		return nil, err
	}
	if updateBalance.Balance < 0 {
		return nil, errors.New("Balance cannot be negative")
	}
	_, err := ctx.DB().ExecContext(ctx, "UPDATE customers SET balance = balance + ? WHERE id = ?", updateBalance.Balance, customerID)

	return nil, err
}
func WithdrawMoney(ctx *gofr.Context) (interface{}, error) {
	customerID := ctx.PathParam("id")

	var withdrawalData struct {
		Balance float64 `json:"balance"`
	}
	if err := json.NewDecoder(ctx.Request().Body).Decode(&withdrawalData); err != nil {
		return nil, err
	}

	// Check if the withdrawal amount is not negative
	if withdrawalData.Balance < 0 {
		return nil, errors.New("Withdrawal amount cannot be negative")
	}

	// Check if there is sufficient balance for withdrawal
	var currentBalance float64
	err := ctx.DB().QueryRowContext(ctx, "SELECT balance FROM customers WHERE id = ?", customerID).Scan(&currentBalance)
	if err != nil {
		return nil, err
	}

	if currentBalance < withdrawalData.Balance {
		return nil, errors.New("Insufficient balance for withdrawal")
	}

	// Update the balance after withdrawal
	_, err = ctx.DB().ExecContext(ctx, "UPDATE accounts SET balance = balance - ? WHERE id = ?", withdrawalData.Balance, customerID)
	if err != nil {
		return nil, err
	}

	return nil, nil
}
