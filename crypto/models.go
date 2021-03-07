package crypto

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type currencyModel struct {
	BRL int `binding:"required"`
	EUR int `binding:"required"`
	CAD int `binding:"required"`
}

var currencies currencyModel = getCurrenciesFromJSON()

func getCurrenciesFromJSON() currencyModel {
	// read local json file
	data, err := ioutil.ReadFile("./crypto/currencies.json")
	if err != nil {
		fmt.Print(err)
	}

	return parseDataToObjct(data)
}

func parseDataToObjct(data []byte) currencyModel {
	var currencies currencyModel

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

	// Convert the string values to Int
	currencies.BRL, _ = strconv.Atoi(strings.Replace(obj.BRL, ".", "", 1))
	currencies.EUR, _ = strconv.Atoi(strings.Replace(obj.EUR, ".", "", 1))
	currencies.CAD, _ = strconv.Atoi(strings.Replace(obj.CAD, ".", "", 1))

	return currencies
}
