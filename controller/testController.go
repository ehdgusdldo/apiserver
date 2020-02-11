package controller

import (
	"log"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/ehdgusdldo/APIServer/mid"
	"github.com/gin-gonic/gin"
)

//HelloHandler 미들웨어테스트
func HelloHandler(c *gin.Context) {
	// payload, _ := c.Get("JWT_PAYLOAD")
	// log.Println(payload)
	// log.Println("====================")
	log.Println(c)
	claims := jwt.ExtractClaims(c)
	log.Println("====================")
	log.Println(claims)
	log.Println("================parseToken===============")
	log.Println(mid.AuthMiddleware.ParseToken(c))

	user, _ := c.Get("id")
	log.Println("===user====")
	log.Println(user)
	c.JSON(200, gin.H{
		"userName":  user.(*mid.User).UserName,
		"cid":       user.(*mid.User).CID,
		"userId":    user.(*mid.User).UID,
		"isadmin :": user.(*mid.User).IsAdmin,
		"text":      "Hello World.",
	})
}
