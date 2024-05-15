package api_tests

import (
	"axiomzadatak/config"
	"axiomzadatak/dto"
	"encoding/json"
	"fmt"
	"log"

	"github.com/go-resty/resty/v2"
)

var properties *config.Properties

func init() {
	properties, _ = config.ReadPropertiesFile()
}

func getClient(location string) *resty.Client {
	if properties == nil {
		log.Fatal("Properties are not loaded")
	}
	client := resty.New()
	if location == "user" {
		client.SetBaseURL(properties.UserBaseUrl)
	} else if location == "order" {
		client.SetBaseURL(properties.OrderBaseUrl)
	}

	client.SetBasicAuth(properties.Username, properties.Password)
	return client
}

func getClientWrongPassword(location string) *resty.Client {
	if properties == nil {
		log.Fatal("Properties are not loaded")
	}
	client := resty.New()
	if location == "user" {
		client.SetBaseURL(properties.UserBaseUrl)
	} else if location == "order" {
		client.SetBaseURL(properties.OrderBaseUrl)
	}

	client.SetBasicAuth(properties.Username, properties.IncorrectPassword)
	return client
}

func getUserByIDWithWrongPassword(id int) (*resty.Response, error) {
	client := getClientWrongPassword("user")
	resp, err := client.R().
		SetPathParam("id", fmt.Sprintf("%d", id)).
		SetResult(&dto.UserDTO{}).
		Get("/users/{id}")

	return resp, err
}

func getUserByID(id int) (*dto.UserDTO, *resty.Response, error) {
	client := getClient("user")
	resp, err := client.R().
		SetPathParam("id", fmt.Sprintf("%d", id)).
		SetResult(&dto.UserDTO{}).
		Get("/users/{id}")

	if err != nil {
		return nil, resp, err
	}

	return resp.Result().(*dto.UserDTO), resp, nil
}

func createUser(requestBody string) (int, error, *resty.Response) {
	client := getClient("user")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(requestBody).
		Post("/users")

	if err != nil {
		return 0, err, resp
	}

	// Extracting ID from response
	var result map[string]interface{}
	if err := json.Unmarshal(resp.Body(), &result); err != nil {
		return 0, err, resp
	}

	// checking if ID is existing
	id, ok := result["id"].(float64)
	if !ok {
		return 0, fmt.Errorf("ID not found in response"), resp
	}

	// Converting from float to int
	userID := int(id)

	return userID, nil, resp
}

func createUserWithWrongPassword(requestBody string) (error, *resty.Response) {
	client := getClientWrongPassword("user")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(requestBody).
		Post("/users")

	return err, resp
}

func getOrderByID(id int) (*dto.OrderDTO, error, *resty.Response) {
	client := getClient("order")
	resp, err := client.R().
		SetPathParam("id", fmt.Sprintf("%d", id)).
		SetResult(&dto.OrderDTO{}).
		Get("/orders/{id}")

	if err != nil {
		return nil, err, resp
	}

	return resp.Result().(*dto.OrderDTO), nil, resp
}

func getOrderByIDWithWrongPassword(id int) (error, *resty.Response) {
	client := getClientWrongPassword("order")
	resp, err := client.R().
		SetPathParam("id", fmt.Sprintf("%d", id)).
		SetResult(&dto.OrderDTO{}).
		Get("/orders/{id}")

	return err, resp
}

func placeOrder(requestBody string) (int, int, error, *resty.Response) {
	client := getClient("order")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(requestBody).
		Post("/orders")

	if err != nil {
		return 0, 0, err, resp
	}

	// Extracting from response
	var result map[string]interface{}
	if err := json.Unmarshal(resp.Body(), &result); err != nil {
		return 0, 0, err, resp
	}

	id, ok := result["id"].(float64)
	if !ok {
		return 0, 0, fmt.Errorf("ID not found in response"), resp
	}

	userId, ok := result["userId"].(float64)
	if !ok {
		return 0, 0, fmt.Errorf("User ID not found in response"), resp
	}

	//Converting
	orderID := int(id)
	orderUserID := int(userId)

	return orderID, orderUserID, nil, resp
}

func placeOrderWithWrongPassword(requestBody string) (error, *resty.Response) {
	client := getClientWrongPassword("order")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(requestBody).
		Post("/orders")

	return err, resp
}
