package api_tests

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRetrieveOrderByID(t *testing.T) {
	order, err, resp := getOrderByID(456)
	if err != nil {
		log.Fatalf("Error retrieving user: %v", err)
	}

	assert.NoError(t, err)
	assert.Equal(t, 456, order.ID)
	assert.Equal(t, 123, order.UserID)
	assert.Greater(t, order.TotalAmount, 0.0)
	assert.NotEmpty(t, order.Status)
	assert.Equal(t, 200, resp.StatusCode(), "Expected status code to be 200")
}

func TestRetrieveOrderByIDWithWrongPassword(t *testing.T) {
	err, resp := getOrderByIDWithWrongPassword(456)

	assert.Error(t, err)
	assert.Equal(t, 401, resp.StatusCode(), "Expected status code to be 401")
}

func TestPlaceOrderSuccess(t *testing.T) {
	requestBody := `{
		"userId": 123,
		"totalAmount": 576.23,
		"status": "pending"
	}`

	orderId, userOrderId, err, resp := placeOrder(requestBody)
	if err != nil {
		log.Fatalf("Error placing order: %v", err)
	}

	assert.NoError(t, err)
	assert.Greater(t, orderId, 0)
	assert.Equal(t, userOrderId, 123)
	assert.Equal(t, resp.StatusCode(), 201)
	assert.Equal(t, 201, resp.StatusCode(), "Expected status code to be 201")
}

func TestPlaceOrderWithWrongPassword(t *testing.T) {
	requestBody := `{
		"userId": 123,
		"totalAmount": 576.23,
		"status": "pending"
	}`

	err, resp := placeOrderWithWrongPassword(requestBody)

	assert.Error(t, err)
	assert.Equal(t, resp.StatusCode(), 401)
	assert.Equal(t, 401, resp.StatusCode(), "Expected status code to be 401")
}

func TestRetrieveOrderByInvalidID(t *testing.T) {
	_, err, resp := getOrderByID(0)
	assert.Error(t, err)
	assert.Equal(t, 400, resp.StatusCode(), "Expected status code to be 400")
}

func TestRetrieveOrderByNonExistentID(t *testing.T) {
	_, err, resp := getOrderByID(865424)
	assert.Error(t, err)
	assert.Equal(t, 404, resp.StatusCode(), "Expected status code to be 404")
}

func TestPlaceOrderWithInvalidUserID(t *testing.T) {
	requestBody := `{
		"userId": 0,
		"totalAmount": 576.23,
		"status": "pending"
	}`

	_, _, err, resp := placeOrder(requestBody)
	assert.Error(t, err)
	statusCode := resp.StatusCode()
	assert.Equal(t, 400, statusCode, "Expected status code to be 400")
}
func TestPlaceOrderWithInvalidStatus(t *testing.T) {
	requestBody := `{
		"userId": 123,
		"totalAmount": 576.23,
		"status": "invalid_status"
	}`

	_, _, err, resp := placeOrder(requestBody)
	assert.Error(t, err)
	assert.Equal(t, 400, resp.StatusCode(), "Expected status code to be 400")
}

func TestPlaceOrderWithoutTotalAmount(t *testing.T) {
	requestBody := `{
		"userId": 123,
		"status": "pending"
	}`

	_, _, err, resp := placeOrder(requestBody)
	assert.Error(t, err)
	assert.Equal(t, 400, resp.StatusCode(), "Expected status code to be 400")
}
