package crypto

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func getCryptoBtc() gin.H {
	var response gin.H

	resp, err := http.Get("https://api.coindesk.com/v1/bpi/currentprice/BTC.json")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var cryptoRespose cryptoResponse
	err = json.Unmarshal(body, &cryptoRespose)
	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Println(cryptoRespose)

	return response
}
