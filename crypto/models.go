package crypto

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
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

// Converte the received data to currencyModel object and returns it
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

	// Convert the string values to int
	currencies.BRL, _ = strconv.Atoi(strings.Replace(obj.BRL, ".", "", 1))
	currencies.EUR, _ = strconv.Atoi(strings.Replace(obj.EUR, ".", "", 1))
	currencies.CAD, _ = strconv.Atoi(strings.Replace(obj.CAD, ".", "", 1))

	return currencies
}

func getCryptoBtc() gin.H {
	var response gin.H

	return response
}
