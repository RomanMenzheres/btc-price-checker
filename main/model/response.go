package model

// Struct type for response from api which gives me a btc ptice
type Response struct {
	BaseCurrency string `json:"base"`
	Currency     string `json:"currency"`
	Price        string `json:"amount"`
}

type Data struct {
	Data Response `json:"data"`
}