package controller

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/kantahasegawa/hair-salon-reservation-system/src/entity"
	"github.com/kantahasegawa/hair-salon-reservation-system/src/usecase"
	"github.com/stretchr/testify/assert"
)

type mockMenuInteractor struct {
	usecase.MenuInteractor
}

func (i *mockMenuInteractor) GetMenus() ([]entity.Menu, error) {
	result := []entity.Menu{{Id: "1", Name: "one", Sex: "M", Price: 1000, Time: 60}, {Id: "2", Name: "two", Sex: "F", Price: 2000, Time: 60}}
	return result, nil
}

func (i *mockMenuInteractor) GetMenu(id string) (entity.Menu, error) {
	result := entity.Menu{Id: "1", Name: "one", Sex: "M", Price: 1000, Time: 60}
	return result, nil
}

func (i *mockMenuInteractor) AddMenu(name string, sex string, price int, time int) (string, error) {
	return "id", nil
}

func (i *mockMenuInteractor) UpdateMenu(id string, name string, sex string, price int, time int) error {
	return nil
}

func (i *mockMenuInteractor) DeleteMenu(id string) error {
	return nil
}

func TestMenuIndexHndler(t *testing.T) {
	expect := "{\"result\":[{\"id\":\"1\",\"name\":\"one\",\"sex\":\"M\",\"price\":1000,\"time\":60},{\"id\":\"2\",\"name\":\"two\",\"sex\":\"F\",\"price\":2000,\"time\":60}]}"
	controller := NewMenuController(&mockMenuInteractor{})
	response := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(response)
	c.Request, _ = http.NewRequest(
		http.MethodGet,
		"/Menus",
		nil,
	)
	controller.IndexHandler(c)
	assert.Equal(t, 200, response.Code)
	assert.Equal(t, expect, response.Body.String())
}

func TestMenuShowHandler(t *testing.T) {
	expect := "{\"result\":{\"id\":\"1\",\"name\":\"one\",\"sex\":\"M\",\"price\":1000,\"time\":60}}"
	controller := NewMenuController(&mockMenuInteractor{})
	response := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(response)
	c.Request, _ = http.NewRequest(
		http.MethodGet,
		"/Menu/1",
		nil,
	)
	controller.ShowHandler(c)
	assert.Equal(t, 200, response.Code)
	assert.Equal(t, expect, response.Body.String())
}

func TestMenuNewHandler(t *testing.T) {
	expect := "{\"id\":\"id\"}"
	jsonStr := `{"name":"test","sex":"M","price":1000,"time":60}`
	controller := NewMenuController(&mockMenuInteractor{})
	response := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(response)
	c.Request, _ = http.NewRequest(
		http.MethodPost,
		"/Menu",
		bytes.NewBuffer([]byte(jsonStr)),
	)
	c.Header("Content-Type", "application/json")
	controller.NewHandler(c)
	assert.Equal(t, 200, response.Code)
	assert.Equal(t, expect, response.Body.String())
}

func TestMenuUpdateHandler(t *testing.T) {
	expect := "{\"message\":\"success\"}"
	jsonStr := `{"id":"1","name":"test","sex":"M","price":1000,"time":60}`
	controller := NewMenuController(&mockMenuInteractor{})
	response := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(response)
	c.Request, _ = http.NewRequest(
		http.MethodPatch,
		"/Menu",
		bytes.NewBuffer([]byte(jsonStr)),
	)
	c.Header("Content-Type", "application/json")
	controller.UpdateHandler(c)
	assert.Equal(t, 200, response.Code)
	assert.Equal(t, expect, response.Body.String())
}

func TestMenuDeleteHandler(t *testing.T) {
	expect := "{\"message\":\"success\"}"
	controller := NewMenuController(&mockMenuInteractor{})
	response := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(response)
	c.Request, _ = http.NewRequest(
		http.MethodGet,
		"/Menu/1",
		nil,
	)
	controller.DeleteHandler(c)
	assert.Equal(t, 200, response.Code)
	assert.Equal(t, expect, response.Body.String())
}
