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

type mockBeauticianInteractor struct {
	usecase.BeauticianInteractor
}

func (i *mockBeauticianInteractor) GetBeauticians() ([]entity.Beautician, error) {
	result := []entity.Beautician{{Id: "1", Name: "one", Sex: "M", Price: 1000}, {Id: "2", Name: "two", Sex: "F", Price: 2000}}
	return result, nil
}

func (i *mockBeauticianInteractor) GetBeautician(id string) (entity.Beautician, error) {
	result := entity.Beautician{Id: "1", Name: "one", Sex: "M", Price: 1000}
	return result, nil
}

func (i *mockBeauticianInteractor) AddBeautician(name string, sex string, price int) (string, error) {
	return "id", nil
}

func (i *mockBeauticianInteractor) UpdateBeautician(id string, name string, sex string, price int) error {
	return nil
}

func (i *mockBeauticianInteractor) DeleteBeautician(id string) error {
	return nil
}

func TestBeauticianIndexHndler(t *testing.T) {
	expect := "{\"result\":[{\"id\":\"1\",\"name\":\"one\",\"sex\":\"M\",\"price\":1000},{\"id\":\"2\",\"name\":\"two\",\"sex\":\"F\",\"price\":2000}]}"
	controller := NewBeauticianController(&mockBeauticianInteractor{})
	response := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(response)
	c.Request, _ = http.NewRequest(
		http.MethodGet,
		"/beauticians",
		nil,
	)
	controller.IndexHandler(c)
	assert.Equal(t, 200, response.Code)
	assert.Equal(t, expect, response.Body.String())
}

func TestBeauticianShowHandler(t *testing.T){
	expect := "{\"result\":{\"id\":\"1\",\"name\":\"one\",\"sex\":\"M\",\"price\":1000}}"
	controller := NewBeauticianController(&mockBeauticianInteractor{})
	response := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(response)
	c.Request, _ = http.NewRequest(
		http.MethodGet,
		"/beautician/1",
		nil,
	)
	controller.ShowHandler(c)
	assert.Equal(t, 200, response.Code)
	assert.Equal(t, expect, response.Body.String())
}

func TestBeauticianNewHandler(t *testing.T){
	expect := "{\"id\":\"id\"}"
	jsonStr := `{"name":"test","sex":"M","price":1000}`
	controller := NewBeauticianController(&mockBeauticianInteractor{})
	response := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(response)
	c.Request, _ = http.NewRequest(
		http.MethodPost,
		"/beautician",
		bytes.NewBuffer([]byte(jsonStr)),
	)
	c.Header("Content-Type", "application/json")
	controller.NewHandler(c)
	assert.Equal(t, 200, response.Code)
	assert.Equal(t, expect, response.Body.String())
}

func TestBeauticianUpdateHandler(t *testing.T){
	expect := "{\"message\":\"success\"}"
	jsonStr := `{"id":"1","name":"test","sex":"M","price":1000}`
	controller := NewBeauticianController(&mockBeauticianInteractor{})
	response := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(response)
	c.Request, _ = http.NewRequest(
		http.MethodPatch,
		"/beautician",
		bytes.NewBuffer([]byte(jsonStr)),
	)
	c.Header("Content-Type", "application/json")
	controller.UpdateHandler(c)
	assert.Equal(t, 200, response.Code)
	assert.Equal(t, expect, response.Body.String())
}

func TestBeauticianDeleteHandler(t *testing.T){
	expect := "{\"message\":\"success\"}"
	controller := NewBeauticianController(&mockBeauticianInteractor{})
	response := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(response)
	c.Request, _ = http.NewRequest(
		http.MethodGet,
		"/beautician/1",
		nil,
	)
	controller.DeleteHandler(c)
	assert.Equal(t, 200, response.Code)
	assert.Equal(t, expect, response.Body.String())
}
