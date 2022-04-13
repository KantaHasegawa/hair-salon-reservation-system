package controller

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestGreetingHandler(t *testing.T) {
	rootController := NewRootController()
	expect := map[string]interface{}{"message": "helloWorld"}
	response := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(response)
	c.Request, _ = http.NewRequest(
		http.MethodGet,
		"/",
		strings.NewReader(""),
	)
	rootController.GreetingHandler(c)
	if response.Code != 200 {
		t.Errorf("invalid status")
	}
	var responseBody map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &responseBody)
	for key, value := range responseBody {
		if value != expect[key] {
			t.Errorf("invalid body")
		}
	}
}
