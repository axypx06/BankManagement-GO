package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestHelloHandler(t *testing.T) {
	const host = "http://localhost:9000"
	go main()
	time.Sleep(time.Second * 3)

	testcases := []struct {
		path       string
		statusCode int
		body       string
	}{
		{"/", 404, ""},
		{"/greet", 200, ""},
	}

	for _, tc := range testcases {
		req, _ := http.NewRequest("GET", host+tc.path, nil)
		c := http.Client{}
		resp, err := c.Do(req)
		if err != nil {
			t.Error("Could not get response", err)
		}

		if resp != nil && resp.StatusCode != tc.statusCode {
			t.Errorf("Failed. \t Expected %v\t Got %v", tc.statusCode, resp.StatusCode)
		}
	}
}
func TestGetCustomer(t *testing.T) {
	const host = "http://localhost:9000"
	go main()
	time.Sleep(time.Second * 3)

	testcases := []struct {
		path       string
		statusCode int
		body       string
	}{
		{"/", 404, ""},
		{"/admin/viewCustomer", 200, ""},
	}

	for _, tc := range testcases {
		req, _ := http.NewRequest("GET", host+tc.path, nil)
		c := http.Client{}
		resp, err := c.Do(req)
		if err != nil {
			t.Error("Could not get response", err)
		}

		if resp != nil && resp.StatusCode != tc.statusCode {
			t.Errorf("Failed. \t Expected %v\t Got %v", tc.statusCode, resp.StatusCode)
		}
	}
}
func TestGetCustomerByID(t *testing.T) {
	const host = "http://localhost:9000"
	go main()
	time.Sleep(time.Second * 3)

	expectedResponse := `{"data":{"id":3,"name":"sahil","balance":200}}`

	req, _ := http.NewRequest("GET", host+"/admin/viewCustomer/3", nil)
	c := http.Client{}
	resp, err := c.Do(req)
	if err != nil {
		t.Error("Could not get response", err)
	}

	assert.Equal(t, http.StatusOK, resp.StatusCode, "Expected status code 200")
	body, _ := ioutil.ReadAll(resp.Body)
	var actualResponse map[string]interface{}
	err = json.Unmarshal(body, &actualResponse)
	if err != nil {
		t.Error("Error decoding JSON response", err)
	}

	assert.JSONEq(t, expectedResponse, string(body), "Expected and actual responses do not match")
}
func TestCreateCustomer(t *testing.T) {
	const host = "http://localhost:9000"
	go main()
	time.Sleep(time.Second * 3)

	name := "John"

	req, _ := http.NewRequest("POST", host+"/admin/addCustomer/"+name, nil)
	c := http.Client{}

	resp, err := c.Do(req)
	if err != nil {
		t.Error("Could not get response", err)
	}

	assert.Equal(t, http.StatusCreated, resp.StatusCode, "Expected status code 201")

	body, _ := ioutil.ReadAll(resp.Body)
	var response map[string]interface{}
	err = json.Unmarshal(body, &response)
	if err != nil {
		t.Error("Error decoding JSON response", err)
	}

	expectedResponse := map[string]interface{}{
		"data": nil,
	}

	assert.Equal(t, expectedResponse, response, "Expected and actual responses do not match")
}

func TestDeleteCustomer(t *testing.T) {
	const host = "http://localhost:9000"
	go main()
	time.Sleep(time.Second * 3)

	customerID := 1

	req, _ := http.NewRequest("DELETE", fmt.Sprintf("%s/admin/deleteCustomer/%d", host, customerID), nil)
	c := http.Client{}

	resp, err := c.Do(req)
	if err != nil {
		t.Error("Could not get response", err)
	}

	assert.Equal(t, http.StatusNoContent, resp.StatusCode, "Expected status code 204 for successful deletion")

	if resp.ContentLength != 0 {
		t.Error("Expected empty response body for status code 204")
	}
}
func TestAddMoney(t *testing.T) {

	const host = "http://localhost:9000"
	go main()
	time.Sleep(time.Second * 3)
	testBalance := 100.00
	requestBody := map[string]interface{}{"balance": testBalance}
	body, _ := json.Marshal(requestBody)
	req, _ := http.NewRequest("PUT", host+"/customer/3/add", bytes.NewBuffer(body))
	c := http.Client{}

	resp, err := c.Do(req)
	if err != nil {
		t.Error("Could not get response", err)
	}

	assert.Equal(t, http.StatusOK, resp.StatusCode, "Expected status code 200")
	body, _ = ioutil.ReadAll(resp.Body)
	var response map[string]interface{}
	err = json.Unmarshal(body, &response)
	if err != nil {
		t.Error("Error decoding JSON response", err)
	}

	expectedResponse := map[string]interface{}(map[string]interface{}{"data": map[string]interface{}{"msg": "success"}})

	assert.Equal(t, expectedResponse, response, "Expected and actual responses do not match")

}
func TestWithdrawMoney(t *testing.T) {

	const host = "http://localhost:9000"
	go main()
	time.Sleep(time.Second * 3)
	testBalance := 100.00
	requestBody := map[string]interface{}{"balance": testBalance}
	body, _ := json.Marshal(requestBody)
	req, _ := http.NewRequest("PUT", host+"/customer/3/withdraw", bytes.NewBuffer(body))
	c := http.Client{}

	resp, err := c.Do(req)
	if err != nil {
		t.Error("Could not get response", err)
	}

	assert.Equal(t, http.StatusOK, resp.StatusCode, "Expected status code 200")
	body, _ = ioutil.ReadAll(resp.Body)
	var response map[string]interface{}
	err = json.Unmarshal(body, &response)
	if err != nil {
		t.Error("Error decoding JSON response", err)
	}

	expectedResponse := map[string]interface{}(map[string]interface{}{"data": map[string]interface{}{"msg": "success"}})

	assert.Equal(t, expectedResponse, response, "Expected and actual responses do not match")

}
func TestUpdateCustomer(t *testing.T) {
	const host = "http://localhost:9000"
	go main()
	time.Sleep(time.Second * 3)
	newName := ""
	requestBody := map[string]interface{}{"name": newName}
	body, _ := json.Marshal(requestBody)
	req, _ := http.NewRequest("PUT", host+"/admin/update/3", bytes.NewBuffer(body))
	c := http.Client{}

	resp, err := c.Do(req)
	if err != nil {
		t.Error("Could not get response", err)
	}

	assert.Equal(t, http.StatusOK, resp.StatusCode, "Expected status code 200")
	body, _ = ioutil.ReadAll(resp.Body)
	var response map[string]interface{}
	err = json.Unmarshal(body, &response)
	if err != nil {
		t.Error("Error decoding JSON response", err)
	}

	expectedResponse := map[string]interface{}(map[string]interface{}{"data": map[string]interface{}{"msg": "fail"}})

	assert.Equal(t, expectedResponse, response, "Expected and actual responses do not match")

}
