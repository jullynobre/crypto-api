package crypto

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var currencies currencyQuote = getCurrenciesFromJSON()

func getCurrenciesFromJSON() currencyQuote {
	// read local json file
	data, err := ioutil.ReadFile("./crypto/currencies.json")
	if err != nil {
		fmt.Print(err)
	}

	return parseDataToObjct(data)
}

// Converte the received data to currencyQuote object and returns it
func parseDataToObjct(data []byte) currencyQuote {
	var currencies currencyQuote

	type currencyObj struct {
		BRL string `binding:"required"`
		EUR string `binding:"required"`
		CAD string `binding:"required"`
	}

	var obj currencyObj

	var err = json.Unmarshal(data, &obj)
	if err != nil {
		fmt.Println("error:", err)
	}

	// Convert the string values to int
	currencies.BRL, _ = strconv.Atoi(strings.Replace(obj.BRL, ".", "", 1))
	currencies.EUR, _ = strconv.Atoi(strings.Replace(obj.EUR, ".", "", 1))
	currencies.CAD, _ = strconv.Atoi(strings.Replace(obj.CAD, ".", "", 1))

	return currencies
}
