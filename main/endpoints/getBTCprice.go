package endpoints

import (
	"encoding/json"
	"net/http"
	"program/functions"
	"program/model"

	"github.com/gin-gonic/gin"
)

func Rate(context *gin.Context) {
	//Getting btc price from third party service 
	responseBody, err := functions.GetResponseBody("https://api.coinbase.com/v2/prices/BTC-UAH/buy")

	//Checking error
	if err != nil {
		//Bad request response if error
		context.IndentedJSON(http.StatusBadRequest, nil)
		return
	}

	//Getting price from json
	var data model.Data
	json.Unmarshal(responseBody, &data)
	
	//Doing response with price
    context.IndentedJSON(http.StatusOK, gin.H{"amount": data.Data.Price})
}