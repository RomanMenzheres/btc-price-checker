package endpoints

import (
	"fmt"
	"net/http"
	"os"
	"program/functions"
	model "program/model"
	"strings"

	"github.com/gin-gonic/gin"
)

func Subscribe(context *gin.Context) {
	var newEmail model.Email

	//Getting email from context and checking error
	if err := context.BindJSON(&newEmail); err != nil {
		return
	}

	//Opening file or creating if not found and checking error
	file, err := os.OpenFile("db/db.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0660);
    functions.CheckErr(err)

	//Getting email addresses from our file as big string with spaces
    db_emails := functions.GetDbEmails()

	//if our big string does not have new email we write it in file, do response and return
    if !strings.Contains(db_emails, newEmail.Email){
		//Writing in file
		fmt.Fprintln(file, newEmail.Email)

		//Doing appropriate response
		context.IndentedJSON(http.StatusOK, gin.H{"message": "E-mail додано!"})
		return
	}

	//if our big string already have new email we do appropriate response 
	context.IndentedJSON(http.StatusConflict, gin.H{"conflict": "E-mail вже є в базі даних!"})
}