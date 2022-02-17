package main

import (
	"currency/routes"
	"currency/utils"
	"currency/validate"
	"fmt"
	"net/http"
	"time"
)

func main() {
	go validate.ValidateInit()
	go utils.InitCurrencyInfo()
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", 8888),
		Handler:      routes.InitRoute(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	server.ListenAndServe()
}
