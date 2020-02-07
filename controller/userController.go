package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/ehdgusdldo/APIServer/models"
	"github.com/gin-gonic/gin"
)

// Welcome godoc
// @Description 전체사용자 목록조회
// @Router /users [get]
// @Tags User
// @Accept  json
// @Produce  json
// @Success 200 {array} models.Users
// @Failure 500 {object} models.FailedMessage
func GetUserAll(c *gin.Context) {
	user := models.Users{}
	users, err := user.GetAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "failed",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"users":   users,
	})

}

// 단일 고객 조회
// @Summary {id} 사용자아이디
// @Description 사용자 id로 사용자 단일조회
// @Tags User
// @id get-string-by-int
// @Accept  json
// @Produce  json
// @Param id path integer true "사용자아이디"
// @Router /users/{id} [get]
// @Success 200 {object} models.Users
// @Failure 500 {object} models.FailedMessage
func GetUser(c *gin.Context) {

	id, _ := strconv.Atoi(c.Param("id"))

	user := models.Users{UserID: id}
	has, err := user.Get()
	// has 와 err의 처리를 따로 해야하는지
	if !has || err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "failed",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"custom":  user,
	})

}

// 사용자추가
// @Description 사용자추가
// @Router /users [post]
// @Tags User
// @Accept  json
// @Produce  json
// @Param name body string true "사용자명"
// @Param pwd body string true "비밀번호"
// @Param phone body string true "연락처"
// @Param email body string true "email"
// @Param admin body integer true "1"
// @Param custID body integer true "13"
// @Param gfId body integer true "2"
// @Param dept body string false "부서"
// @Param position body string false "직급"
// @Param rec body string false "비고"
// @Success 200 {object} models.Message
// @Failure 500 {object} models.FailedMessage
func RegisterUser(c *gin.Context) {
	var user models.Users
	println("바인딩전")
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "failed",
		})
		return
	}
	println("바인딩완료")

	fmt.Println("-------bindingdata----")
	fmt.Println(user)
	userID, err := user.Add()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "failed",
			"error":   err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"ID":      userID,
	})
}
