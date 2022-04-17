package controller

import (
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

func TestIndexHndler(t *testing.T) {
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
