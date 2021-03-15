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
