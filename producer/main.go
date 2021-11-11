package main

import (
	"github.com/gin-gonic/gin"
	"jwiegmann.de/group-member-service1/member"
)

func main() {

	// Start REST service to add new members
	router := gin.Default()

	router.POST("/test/:member", member.AddMember)

	err := router.Run("localhost:8080")
	if err != nil {
		return
	}
}
