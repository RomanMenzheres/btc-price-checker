package main

import (
	endpoint "program/endpoints"

	"github.com/gin-gonic/gin"
)

//Dunc for initializing
func initialization(){
	router := gin.Default()

	router.GET("/api/rate", endpoint.Rate)
	router.GET("/api/sendEmails", endpoint.SendEmails)
	router.POST("/api/subscribe", endpoint.Subscribe)

	router.Run(":80")
}