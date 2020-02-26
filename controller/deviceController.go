package controller

import (
	"log"
	"net/http"

	"github.com/ehdgusdldo/APIServer/mid"
	"github.com/ehdgusdldo/APIServer/models"
	"github.com/ehdgusdldo/APIServer/util"
	"github.com/gin-gonic/gin"
)

// GetDeviceAll 로그인한 유저 고객사의 모델리스트
// @Description 로그인한 유저 고객사의 모델리스트  Authorization Bearer token 을 넣으세요
// @Router /auth/device [get]
// @Tags device
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Success 200 {array} models.Device
func GetDeviceAll(c *gin.Context) {
	user, _ := c.Get("id")
	custID := user.(*mid.User).CID
	var device []models.Device
	err := util.Engine.Where("cust_id = ?", custID).Find(&device)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "faile check parameter",
			"error":   err,
		})
	}

	c.JSON(http.StatusOK, device)
}

// GetDevice ..
// @Description 모델 단일조회  Authorization Bearer token 을 넣으세요
// @Router /auth/device/{id} [get]
// @Tags device
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param id path string true "장치아이디"
// @Success 200 {object} models.Device
func GetDevice(c *gin.Context) {
	// jwt 토큰 claims
	user, _ := c.Get("id")
	// 토큰안에 custID
	custID := user.(*mid.User).CID
	var device models.Device
	device.DevID = c.Param("id")
	has, err := util.Engine.Where("dev_id = ?", device.DevID).Get(&device)
	// 에러처리
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "faile check parameter",
			"error":   err,
		})
		return
	}
	// select 결과가없을경우
	if !has {
		log.Println("존재하지않는 사이트")
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "no data",
		})
		return
	}
	// 조회된 사이트의 custID와 인증된사용자의 custID가 다를경우
	if device.CustID != custID {
		log.Println("다른고객사임")
		c.JSON(http.StatusForbidden, gin.H{
			"message": "permission denied",
		})
		return
	}

	c.JSON(http.StatusOK, device)
}
