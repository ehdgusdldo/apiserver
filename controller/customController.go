package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/ehdgusdldo/APIServer/models"
	"github.com/gin-gonic/gin"
)

// Welcome godoc
// @Summary {id}고객아이디
// @Description 고객 id로 고객 단일조회
// @Tags Custom
// @id get-string-by-int
// @Accept  json
// @Produce  json
// @Param id path integer true "고객아이디"
// @Router /customs/{id} [get]
// @Success 200 {object} models.Custom
// @Failure 500 {object} models.FailedMessage
func GetCustom(c *gin.Context) {

	id, _ := strconv.Atoi(c.Param("id"))

	custom := models.Custom{CustID: id}
	has, err := custom.Get()
	// has 와 err의 처리를 따로 해야하는지
	if !has || err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "failed",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"custom":  custom,
	})

}

// Welcome godoc
// @Description 전체고객 목록조회
// @Router /customs [get]
// @Tags Custom
// @Accept  json
// @Produce  json
// @Success 200 {array} models.Custom
// @Failure 500 {object} models.FailedMessage
func GetCustomAll(c *gin.Context) {
	custom := models.Custom{}

	customs, err := custom.GetAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "failed",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"customs": customs,
	})

}

// Welcome godoc
// @Description 고객등록
// @Router /customs [post]
// @Tags Custom
// @Accept  json
// @Produce  json
// @Param id body integer true "기본키 원랜입력안해야함"
// @Param eid body string true "고객아이디 ex SAMSUNG"
// @Param name body string true "고객명 ex 삼성"
// @Param postno body string true "우편번호10자이내"
// @Param basadr body string true "기본주소"
// @Param dtladr body string true "상세주소"
// @Param type body integer true "사업자여부 1 or 0"
// @Param orgid body integer true "그라파나 회사 pk 입력안해야함원랜"
// @Success 200 {object} models.Message
// @Failure 500 {object} models.FailedMessage
func RegisterCustom(c *gin.Context) {
	var custom models.Custom
	if err := c.ShouldBindJSON(&custom); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "failed",
		})
		return

	}

	fmt.Println("-------bindingdata----")
	fmt.Println(custom)
	affected, err := custom.Add()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "failed",
			"error":   err,
		})
		return
	}
	fmt.Println(affected)
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}
