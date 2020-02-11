package controller

import (
	"log"

	"github.com/ehdgusdldo/APIServer/mid"
	"github.com/gin-gonic/gin"
)

//HelloHandler 미들웨어테스트
func HelloHandler(c *gin.Context) {
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
