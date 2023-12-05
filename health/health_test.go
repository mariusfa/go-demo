package health

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"todo/middleware"

	"github.com/gin-gonic/gin"
)

func TestGetHealth(t *testing.T) {
	// Setup
	r := gin.New()
	r.Use(middleware.Marius)
	HealthSetup(r)

	// Perform request
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/health", nil)
	r.ServeHTTP(w, req)

	// Then
	if w.Code != 200 {
		t.Errorf("Expected status code 200, got %d", w.Code)
	}
	if w.Body.String() != "ok Marius" {
		t.Errorf("Expected body to be 'ok', got %s", w.Body.String())
	}
}