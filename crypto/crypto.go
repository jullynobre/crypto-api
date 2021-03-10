package crypto

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func getCryptoBtc() (response cryptoResponse) {
	response = requestDataFromCoindeskAPI()
	response = calcCurrenciesRate(response)
	return
}

func requestDataFromCoindeskAPI() cryptoResponse {
	var cryptoRespose cryptoResponse

	// Request from API
	resp, err := http.Get("https://api.coindesk.com/v1/bpi/currentprice/BTC.json")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// Read the response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Convert the []btye object to cryptoResponse object
	err = json.Unmarshal(body, &cryptoRespose)
	if err != nil {
		fmt.Println("error:", err)
	}

	return cryptoRespose
}

func calcCurrenciesRate(cryptoResponse cryptoResponse) cryptoResponse {
	var currencies = getCurrencies()

	fmt.Println(currencies)

	var USDRateStr = cryptoResponse.BPI.USD.Rate
	USDRateStr = strings.ReplaceAll(USDRateStr, ".", "")
	USDRateStr = strings.ReplaceAll(USDRateStr, ",", "")
	var USDRate, _ = strconv.Atoi(USDRateStr)

	var BRLRateInt = (USDRate * currencies.BRL) / 10000
	var CADRateInt = (USDRate * currencies.CAD) / 10000
	var EURRateInt = (USDRate * currencies.EUR) / 10000

	cryptoResponse.BPI.BRL.Rate = convertCurrencyIntToStr(BRLRateInt)
	cryptoResponse.BPI.CAD.Rate = convertCurrencyIntToStr(CADRateInt)
	cryptoResponse.BPI.EUR.Rate = convertCurrencyIntToStr(EURRateInt)

	cryptoResponse.BPI.BRL.RateFloat, _ = strconv.ParseFloat(cryptoResponse.BPI.BRL.Rate, 64)
	cryptoResponse.BPI.CAD.RateFloat, _ = strconv.ParseFloat(cryptoResponse.BPI.CAD.Rate, 64)
	cryptoResponse.BPI.EUR.RateFloat, _ = strconv.ParseFloat(cryptoResponse.BPI.EUR.Rate, 64)

	cryptoResponse.BPI.BRL.Code = "BRL"
	cryptoResponse.BPI.CAD.Code = "CAD"
	cryptoResponse.BPI.EUR.Code = "EUR"

	cryptoResponse.BPI.BRL.Description = "Brazilian Real"
	cryptoResponse.BPI.CAD.Description = "Canadian Dollar"
	cryptoResponse.BPI.EUR.Description = "Euro"

	return cryptoResponse
}
