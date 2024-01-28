package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/naolson2025/go-rest-api/models"
	"github.com/stretchr/testify/assert"
)

func clearDB() {
	os.Remove("test.db")
}

func TestGetEvents(t *testing.T) {
	os.Setenv("DB_NAME", "test.db")

	server := setupServer()

	response := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/events", nil)

	server.ServeHTTP(response, request)

	assert.Equal(t, 200, response.Code)

	clearDB()
}

func TestGetEvent(t *testing.T) {
	os.Setenv("DB_NAME", "test.db")

	server := setupServer()

	newEvent := []byte(`{
		"name": "Test Event",
		"description": "Test Description",
		"location": "Test Location",
		"dateTime": "2020-01-01T00:00:00Z",
		"userID": 1
	}`)

	// create event
	createResponse := httptest.NewRecorder()
	createRequest, _ := http.NewRequest("POST", "/events", bytes.NewBuffer(newEvent))

	server.ServeHTTP(createResponse, createRequest)
	assert.Equal(t, 201, createResponse.Code)

	response := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/events/1", nil)

	server.ServeHTTP(response, request)

	var event models.Event
	err := json.Unmarshal([]byte(response.Body.Bytes()), &event)
	if err != nil {
		fmt.Println("Error parsing response body:", err)
		return
	}
	fmt.Println(event)

	assert.Equal(t, 200, response.Code)
	assert.Equal(t, "Test Event", event.Name)

	clearDB()
}
