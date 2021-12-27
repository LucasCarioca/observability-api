package auth

import (
	"github.com/gin-gonic/gin"
	"log"
)

//AuthCheck middleware function that validates if the request is authenticated
func AuthCheck(ctx *gin.Context) {
	log.Println("checking apikey")
	log.Println("not yet implemented")
}
