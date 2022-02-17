package controllers

import (
	"currency/models"
	"currency/services"
	"fmt"
	"net/http"

	validateHelper "currency/validate"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type CurrencyController struct {
	serv *services.CurrencyService
}

// NewCurrencyController NewCurrencyController
func NewCurrencyController(service *services.CurrencyService) *CurrencyController {
	return &CurrencyController{
		serv: service,
	}
}

// CalcCurrency CalcCurrency
func (c *CurrencyController) CalcCurrency(ctx *gin.Context) {
	var dataForm models.InputCurrency
	err := ctx.ShouldBind(&dataForm)
	if err != nil {
		fmt.Println(err)
	}
	validate := validateHelper.GetValidate()
	err = validate.Struct(dataForm)
	if errors, ok := err.(validator.ValidationErrors); ok {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errors": validateHelper.Translate(errors),
		})

	}
	price, err := c.serv.GetCurrency(dataForm)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	ctx.JSON(http.StatusOK, gin.H{
		"result": price,
	})
}
