package functions

import (
	"io/ioutil"
	"net/http"
)

//func for gettig body from response
func GetResponseBody(url string) ([]byte, error) {
	//Getting response from url and checking error
	response, err := http.Get(url)
	if err != nil {
		//Return if error
		return nil, err
	}

	//Getting body from response and checking error
	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		//Return if error
		return nil, err
	}

	//Return body
	return responseBody, nil
}