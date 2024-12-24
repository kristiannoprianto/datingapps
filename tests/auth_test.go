package tests

import (
	"apps/datingapps/routes/auth"
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestSignUp(t *testing.T) {
	r := gin.Default()
	r.POST("/signup", auth.SignUp)

	// Test data for signing up
	payload := `{"email": "test@example.com", "password": "securepassword", "name": "Test User"}`

	// Make request to /signup
	w := performRequest(r, "POST", "/signup", payload)

	assert.Equal(t, 200, w.Code) // Expecting HTTP status 200
}

func performRequest(r *gin.Engine, method, path string, payload string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, bytes.NewBufferString(payload))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}
