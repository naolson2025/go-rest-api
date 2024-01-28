package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetEvents(t *testing.T) {
	server := setupServer()

	response := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/events", nil)

	server.ServeHTTP(response, request)

	assert.Equal(t, 200, response.Code)
}