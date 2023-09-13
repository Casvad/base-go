package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type BaseGoExampleController interface {
	GetData(context *gin.Context)
}

type baseGoExampleController struct {
}

func (b *baseGoExampleController) GetData(context *gin.Context) {

	context.JSON(http.StatusOK, map[string]string{"hello": "world"})
}

func ProvideBaseGoExampleController() BaseGoExampleController {

	return &baseGoExampleController{}
}
