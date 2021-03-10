package crypto

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"io/ioutil"
	"strconv"
	"strings"
)

func getCurrencies() currencyQuote {
	// read local json file
	data, err := ioutil.ReadFile("./crypto/currencies.json")
	if err != nil {
		fmt.Print(err)
	}

	var obj = parseDataToObjct(data)
	var currencies currencyQuote

	currencies.BRL = convertCurrencyStrToInt(obj.BRL)
	currencies.EUR = convertCurrencyStrToInt(obj.EUR)
	currencies.CAD = convertCurrencyStrToInt(obj.CAD)

	return currencies
}

func updateCurrency(currency string, value int) {
	data, err := ioutil.ReadFile("./crypto/currencies.json")
	if err != nil {
		fmt.Print(err)
	}
	var currencies = parseDataToObjct(data)

	switch currency {
	case "BRL":
		currencies.BRL = convertCurrencyIntToStr(value)
	case "CAD":
		currencies.CAD = convertCurrencyIntToStr(value)
	case "EUR":
		currencies.EUR = convertCurrencyIntToStr(value)
	default:
		fmt.Println("Invalid currency")
	}

	// Converting currenciesObj to data
	data, err = json.Marshal(currencies)
	if err != nil {
		fmt.Println("error:", err)
	}
	// Saving data into file
	err = ioutil.WriteFile("./crypto/currencies.json", data, fs.ModeType)
	if err != nil {
		fmt.Println("error:", err)
	}
}

// Converte the received data to currencyQuote object and returns it
func parseDataToObjct(data []byte) currencyQuoteStr {
	var obj currencyQuoteStr

	var err = json.Unmarshal(data, &obj)
	if err != nil {
		fmt.Println("error:", err)
	}

	return obj
}

func convertCurrencyStrToInt(currencyStr string) int {
	currencyInt, _ := strconv.Atoi(strings.Replace(currencyStr, ".", "", 1))
	return currencyInt
}

func convertCurrencyIntToStr(currencyInt int) string {
	currencyStr := strconv.Itoa(currencyInt)
	currencyStr = currencyStr[0:len(currencyStr)-4] + "." + currencyStr[len(currencyStr)-4:]
	return currencyStr
}
