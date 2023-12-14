package main

import (
	"encoding/json"

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
	app.GET("/admin/viewCustomer", GetCustomer)
	app.GET("/admin/viewCustomer/{id}", GetCustomerByID)
	app.POST("/admin/addCustomer/{name}", CreateCustomer)
	app.DELETE("/admin/deleteCustomer/{id}", DeleteCustomer)
	app.PUT("/admin/update/{id}", UpdateCustomer)
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

	if updateData.Name == "" {
		return map[string]string{"msg": "fail"}, nil
	}

	
	_, err := ctx.DB().ExecContext(ctx, "UPDATE customers SET name = ? WHERE id = ?", updateData.Name, customerID)
	if err != nil {
		return nil, err
	}

	// Return a success response
	return map[string]string{"msg": "success"}, nil
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
		return map[string]string{"msg": "negative balance"}, nil
	}
	_, err := ctx.DB().ExecContext(ctx, "UPDATE customers SET balance = balance + ? WHERE id = ?", updateBalance.Balance, customerID)
	if err != nil {
		return nil, err
	}
	return map[string]string{"msg": "success"}, nil
}
func WithdrawMoney(ctx *gofr.Context) (interface{}, error) {
	customerID := ctx.PathParam("id")

	var withdrawalData struct {
		Balance float64 `json:"balance"`
	}
	if err := json.NewDecoder(ctx.Request().Body).Decode(&withdrawalData); err != nil {
		return nil, err
	}
	if withdrawalData.Balance < 0 {
		return map[string]string{"msg": "negative amount"}, nil
	}

	var currentBalance float64
	err := ctx.DB().QueryRowContext(ctx, "SELECT balance FROM customers WHERE id = ?", customerID).Scan(&currentBalance)
	if err != nil {
		return nil, err
	}

	if currentBalance < withdrawalData.Balance {
		return map[string]string{"msg": "Insuffcient Balance"}, nil
	}

	_, err = ctx.DB().ExecContext(ctx, "UPDATE customers SET balance = balance - ? WHERE id = ?", withdrawalData.Balance, customerID)
	if err != nil {
		return nil, err
	}

	return map[string]string{"msg": "success"}, nil
}
