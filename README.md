# Crypto API
### Requirements
Go v1.16 ([Download](https://golang.org/dl/))
### Running
Execute this command on project root folder:
```Bash
go run main.go
```
### Endpoints
This API is setted to run on port 8000 in order to avoid conflicts with the frontend server.
- [POST] http://localhost:8000/api/login  
- [GET]  http://localhost:8000/api/crypto/btc 
- [POST] http://localhost:8000/api/crypto/btc
### The currencies.json file
A "0" was added to the last decimal place of the currencies values to simplify calculations and conversions.
```JSON
{
  "BRL": "5.4000",
  "EUR": "0.9200",
  "CAD": "1.4400"
}
```
### The approach to deal with float points
The float values are converted to int before calcs do avoid rounding and inconsistencies of floats.
When converting the value to integer the unit of measurement used is 1/1000 of the currency.
Ex.: 4.8 = 4.8000 = 48000

### The CORS problem
This API was implemented in Go using the [gin](https://github.com/gin-gonic/gin) frameword and is was tested during the development using the [Insomnia](https://insomnia.rest) API Client, consequently there was no CORS requests being created.
But in the moment of requesting from API on frontend appeared an error invalidating requests to be made from another source. <br>
After research it was discovered that there is currently an issue in the framework used that is causing this error: [CORS Issue with Basic Auth #1799](https://github.com/gin-gonic/gin/issues/1799)
