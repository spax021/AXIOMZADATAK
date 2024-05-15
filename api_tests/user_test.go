package api_tests

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAccessUserByIDSuccess(t *testing.T) {
	user, resp, err := getUserByID(123)
	if err != nil {
		log.Fatalf("Error retrieving user: %v", err)
	}

	assert.NoError(t, err)
	assert.Equal(t, 123, user.ID)
	assert.Equal(t, "John Doe", user.Name)
	assert.Equal(t, "john.doe@example.com", user.Email)
	assert.Equal(t, "active", user.Status)
	assert.Equal(t, resp.StatusCode(), 200)
}

func TestAccessUserByIDWrongPassword(t *testing.T) {
	resp, err := getUserByIDWithWrongPassword(123)

	assert.Error(t, err)
	assert.Equal(t, 401, resp.StatusCode(), "Expected status code to be 401")
}

func TestCreateUserWrongPassword(t *testing.T) {
	requestBody := `{
		"name": "John Doe",
		"email": "john.doe@example.com",
		"status": "active"
	}`

	err, resp := createUserWithWrongPassword(requestBody)

	assert.Error(t, err)
	assert.Equal(t, 401, resp.StatusCode(), "Expected status code to be 401")
}

func TestAccessUserByInvalidID(t *testing.T) {
	_, resp, err := getUserByID(0)
	assert.Error(t, err)
	assert.Equal(t, 401, resp.StatusCode(), "Expected status code to be 401")
}

func TestAccessUserByNonExistentID(t *testing.T) {
	_, resp, err := getUserByID(5899475)
	assert.Error(t, err)
	assert.Equal(t, 404, resp.StatusCode(), "Expected status code to be 404")
}

func TestCreateUserWithInvalidEmail(t *testing.T) {
	requestBody := `{
		"name": "John Doe",
		"email": "john.doe@example",
		"status": "active"
	}`

	_, err, resp := createUser(requestBody)
	assert.Error(t, err)
	assert.Equal(t, 400, resp.StatusCode(), "Expected status code to be 400")
}

func TestCreateUserWithoutName(t *testing.T) {
	requestBody := `{
		"email": "john.doe@example.com",
		"status": "active"
	}`

	_, err, resp := createUser(requestBody)
	assert.Error(t, err)
	assert.Equal(t, 400, resp.StatusCode(), "Expected status code to be 400")
}

func TestCreateUserWithExistingEmail(t *testing.T) {
	requestBody := `{
		"name": "John Doe",
		"email": "existing.email@example.com",
		"status": "active"
	}`

	_, err, resp := createUser(requestBody)
	assert.Error(t, err)
	assert.Equal(t, 409, resp.StatusCode(), "Expected status code to be 409")
}
