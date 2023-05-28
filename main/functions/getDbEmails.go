package functions

import (
	"io/ioutil"
)

//Func for getting emails from db as a big string with spaces
func GetDbEmails() string {
	//Reading the whole file and checking error
	b, err := ioutil.ReadFile("db/db.txt")
	CheckErr(err)

	//Return as string
	return string(b)
}