package services

import (
	"currency/models"
	"currency/validate"
	"testing"
)

func TestGetCurrency(t *testing.T) {
	validate.ValidateInit()
	srv := NewCurrencyService()

	// Case 1
	data := models.InputCurrency{
		InCurrency:  "TWD",
		OutCurrency: "JPY",
		Price:       100,
	}
	n, err := srv.GetCurrency(data)
	if err != nil {
		t.Error(err)
	}
	if n != "366.90" {
		t.Error("wrong in case 1")
	}

	// Case 2
	data = models.InputCurrency{
		InCurrency:  "TWD",
		OutCurrency: "JPY",
		Price:       36.9,
	}
	n, err = srv.GetCurrency(data)
	if err != nil {
		t.Error(err)
	}
	if n != "135.39" {
		t.Error("wrong in case 2")
	}

	// Case 3
	data = models.InputCurrency{
		InCurrency:  "TWD",
		OutCurrency: "JPY",
		Price:       789.77777777,
	}
	n, err = srv.GetCurrency(data)
	if err != nil {
		t.Error(err)
	}
	if n != "2,897.69" {
		t.Error("wrong in case 3")
	}

	// Case 4
	data = models.InputCurrency{
		InCurrency:  "USD",
		OutCurrency: "USD",
		Price:       123456789,
	}
	n, err = srv.GetCurrency(data)
	if err != nil {
		t.Error(err)
	}
	if n != "123,456,789.00" {
		t.Error("wrong in case 4")
	}

	// Case 5
	data = models.InputCurrency{
		InCurrency:  "USD",
		OutCurrency: "AAA",
		Price:       123456789,
	}
	_, err = srv.GetCurrency(data)
	if err == nil {
		t.Error("The out currency is invalid, case 5")
	}

	// Case 6
	data = models.InputCurrency{
		InCurrency:  "USD",
		OutCurrency: "JPY",
		Price:       0,
	}
	n, err = srv.GetCurrency(data)
	if err != nil {
		t.Error(err)
	}
	if n != "0.00" {
		t.Error("wrong in case 6")
	}
}
