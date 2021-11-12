package main

import (
	"github.com/gin-gonic/gin"
	"jwiegmann.de/producer/member"
	"log"
)

const address = "localhost:8080"

func main() {

	// Start REST service to add new members
	router := gin.Default()

	router.POST("/test/:member", member.AddMember)

	err := router.Run(address)
	if err != nil {
		log.Fatalf("Failed to start webserver at %s: %s", address, err)
	}
}
