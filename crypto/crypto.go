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
	var currencies = getCurrenciesFromJSON()

	var USDRateStr = cryptoResponse.BPI.USD.Rate
	USDRateStr = strings.ReplaceAll(USDRateStr, ".", "")
	USDRateStr = strings.ReplaceAll(USDRateStr, ",", "")
	var USDRate, _ = strconv.Atoi(USDRateStr)

	var BRLRate = strconv.Itoa(USDRate * currencies.BRL)
	var CADRate = strconv.Itoa(USDRate * currencies.CAD)
	var EURRate = strconv.Itoa(USDRate * currencies.EUR)

	BRLRate = BRLRate[0:len(BRLRate)-8] + "." + BRLRate[len(BRLRate)-8:]
	CADRate = CADRate[0:len(CADRate)-8] + "." + CADRate[len(CADRate)-8:]
	EURRate = EURRate[0:len(EURRate)-8] + "." + EURRate[len(EURRate)-8:]

	cryptoResponse.BPI.BRL.Rate = BRLRate
	cryptoResponse.BPI.CAD.Rate = CADRate
	cryptoResponse.BPI.EUR.Rate = EURRate

	cryptoResponse.BPI.BRL.RateFloat, _ = strconv.ParseFloat(BRLRate, 64)
	cryptoResponse.BPI.CAD.RateFloat, _ = strconv.ParseFloat(CADRate, 64)
	cryptoResponse.BPI.EUR.RateFloat, _ = strconv.ParseFloat(EURRate, 64)

	cryptoResponse.BPI.BRL.Code = "BRL"
	cryptoResponse.BPI.CAD.Code = "CAD"
	cryptoResponse.BPI.EUR.Code = "EUR"

	cryptoResponse.BPI.BRL.Description = "Brazilian Real"
	cryptoResponse.BPI.CAD.Description = "Canadian Dollar"
	cryptoResponse.BPI.EUR.Description = "Euro"

	return cryptoResponse
}
