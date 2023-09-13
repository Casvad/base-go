package controllers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func TestBaseGoExampleController_GetData(t *testing.T) {

	c := ProvideBaseGoExampleController()
	r := SetUpRouter()
	r.GET("/", c.GetData)
	req, _ := http.NewRequest("GET", "/", strings.NewReader(""))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	var responseMap map[string]string
	err := json.Unmarshal(w.Body.Bytes(), &responseMap)

	assert.Equal(t, nil, err)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "world", responseMap["hello"])
}
