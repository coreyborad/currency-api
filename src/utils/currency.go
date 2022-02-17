package utils

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type CurrenciesInfo struct {
	Currencies map[string]map[string]float64 `json:"currencies"`
}

var currencies CurrenciesInfo

func InitCurrencyInfo() {

	jsonFile, err := os.Open("./utils/currency.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		panic(err)
	}
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal([]byte(byteValue), &currencies)

}

func GetCurrencyInfo() map[string]map[string]float64 {
	return currencies.Currencies
}
