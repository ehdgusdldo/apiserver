package controller

import (
	"log"
	"net/http"
	"strconv"

	"github.com/ehdgusdldo/APIServer/mid"
	"github.com/ehdgusdldo/APIServer/models"
	"github.com/ehdgusdldo/APIServer/util"
	"github.com/gin-gonic/gin"
)

// GetSiteAll 로그인한 유저 고객사의 사이트리스트
// @Description 로그인한 유저 고객사의 사이트리스트  Authorization Bearer token 을 넣으세요
// @Router /auth/site [get]
// @Tags site
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Success 200 {array} models.Site
func GetSiteAll(c *gin.Context) {
	user, _ := c.Get("id")
	custID := user.(*mid.User).CID
	var site []models.Site
	err := util.Engine.Where("cust_id = ?", custID).Find(&site)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "faile check parameter",
			"error":   err,
		})
	}

	c.JSON(http.StatusOK, site)
}

// GetSite ..
// @Descriptiond 사이트단일조회  Authorization Bearer token 을 넣으세요
// @Router /auth/site/{id} [get]
// @Tags site
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @id get-string-by-int
// @Param id path integer true "사이트아이디"
// @Success 200 {object} models.Site
func GetSite(c *gin.Context) {
	// jwt 토큰 claims
	user, _ := c.Get("id")
	// 토큰안에 custID
	custID := user.(*mid.User).CID
	var site models.Site
	site.SiteID, _ = strconv.Atoi(c.Param("id"))
	has, err := util.Engine.Where("site_id = ?", site.SiteID).Get(&site)
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
	if site.CustID != custID {
		log.Println("다른고객사임")
		c.JSON(http.StatusForbidden, gin.H{
			"message": "permission denied",
		})
		return
	}

	c.JSON(http.StatusOK, site)
}
