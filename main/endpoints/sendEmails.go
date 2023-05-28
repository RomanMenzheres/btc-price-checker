package endpoints

import (
	"encoding/json"
	"net/http"
	"program/functions"
	"program/model"
	"strings"

	"github.com/gin-gonic/gin"
	gomail "github.com/go-gomail/gomail"
)

func SendEmails(context *gin.Context){
	//Getting reponse body with btc price
	responseBody, err := functions.GetResponseBody("http://localhost:80/api/rate") 
	functions.CheckErr(err)

	//Getting btc price from reponse body
	var response model.Response
	json.Unmarshal(responseBody, &response)

	//Getting emails from db file
    db_emails := strings.Fields(functions.GetDbEmails())
	
	//Sending mail to every email in our db file
	for _, email := range db_emails {
		//Create a message
		m := gomail.NewMessage()
		//Seting sender
 		m.SetHeader("From", "menzheres.romaxaxa2@gmail.com")
		//Seting receiver
		m.SetHeader("To", email)	
		//Setting header of mail
		m.SetHeader("Subject", "Super Duper BTC Price Checker")
		//Setting body with btc price of mail
		m.SetBody("text/plain", "Super Duper BTC Price Checker is telling you that BTC price is only " + response.Price + " UAH!")
   
		//Sending mail
		d := gomail.NewDialer("smtp.gmail.com", 587, "menzheres.romaxaxa2@gmail.com", "wbieqzttljkcugyk")
   
		//Checking error
		if err := d.DialAndSend(m); err != nil {
		   panic(err)
		}
	}

	//Doing response 
	context.IndentedJSON(http.StatusOK, gin.H{"message": "E-mailʼи відправлено"})
}