package services

import (
	"currency/models"
	"currency/utils"
	"currency/validate"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

// CurrencyService CurrencyService
type CurrencyService struct {
}

// NewCurrencyService New CurrencyService
func NewCurrencyService() *CurrencyService {
	return &CurrencyService{}
}

func (s *CurrencyService) GetCurrency(data models.InputCurrency) (string, error) {
	if data.Price == 0 {
		return "0.00", nil
	}
	validate := validate.GetValidate()
	err := validate.Struct(data)
	if err != nil {
		return "", err
	}
	currencys := utils.GetCurrencyInfo()
	price := currencys[data.InCurrency][data.OutCurrency] * data.Price

	p := message.NewPrinter(language.English)
	return p.Sprintf("%.2f", price), nil
}
