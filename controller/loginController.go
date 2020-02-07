package controller

import (
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type TokenString struct {
	Token string `json:"token"`
}

type Claims struct {
	TestNo string
	jwt.StandardClaims
}

var expirationTime time.Duration

var JwtKey []byte

func init() {
	expirationTime = 5 * time.Minute
	JwtKey = []byte("5555")
}

// Welcome godoc
// @Description 로그인 response token
// @Router /login [post]
// @Tags auth
// @Accept  json
// @Produce  json
// @Success 200 {object} models.message
// @Failure 500 {object} models.FailedMessage
func Login(c *gin.Context) {
	// Declare the expiration time of the token
	// here, we have kept it as 5 minutes
	time := time.Now().Add(expirationTime)

	// Create the JWT claims, which includes the username and expiry time

	claims := &Claims{
		TestNo: "하이하이",
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: time.Unix(),
		},
	}
	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//Create the JWT string
	tokenString, err := token.SignedString(JwtKey)
	if err != nil {
		// If there is an error in creating the JWT return an internal server error
		// w.WriteHeader(http.StatusInternalServerError)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "failed",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"token": tokenString,
	})
}

func LoginYOU(c *gin.Context) {

	var tokenstring TokenString

	err := c.ShouldBindJSON(&tokenstring)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "failed",
		})
		return
	}
	if tokenstring.Token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "failed",
		})
		return
	}

	claims := &Claims{}

	_, err = jwt.ParseWithClaims(tokenstring.Token, claims, func(token *jwt.Token) (interface{}, error) {
		return JwtKey, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			// statusUnauthorized

			fmt.Println(err)
			fmt.Println("너뭔가이상한데?")
			return
		}
		// satusforbidden
		fmt.Println(err)
		fmt.Println("너뭔가이상한데?")
	}
	fmt.Println("----claims--")
	fmt.Println(claims)
	fmt.Println("----claims TestNO--")
	fmt.Println(claims.TestNo)

}
