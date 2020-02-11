package mid

import (
	"log"
	"os"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/ehdgusdldo/APIServer/models"
	"github.com/ehdgusdldo/APIServer/util"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// User demo
type User struct {
	UserName string
	CID      int
	UID      int
	IsAdmin  bool
}
type login struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

var identityKey = "id"

// jwt secret key
var secretKey = os.Getenv("SECRETKEY")

// AuthMiddleware jwtmiddleware
var AuthMiddleware *jwt.GinJWTMiddleware

func init() {
	var err error
	log.Println("init 통과")
	AuthMiddleware, err = jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "test zone",
		Key:         []byte(secretKey),
		Timeout:     time.Hour * 24 * 7,
		MaxRefresh:  time.Hour * 24 * 7,
		IdentityKey: identityKey,
		//PayloadFunc 셋팅
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*User); ok {
				return jwt.MapClaims{
					"UserName": v.UserName,
					"CID":      v.CID,
					"UID":      v.UID,
					"IsAdmin":  v.IsAdmin,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			return &User{
				UserName: claims["UserName"].(string),
				CID:      int(claims["CID"].(float64)),
				UID:      int(claims["UID"].(float64)),
				IsAdmin:  claims["IsAdmin"].(bool),
			}
		},
		// 로그인로직
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var loginVals login
			// ID와 password 바인딩
			if err := c.ShouldBind(&loginVals); err != nil {
				return "", jwt.ErrMissingLoginValues
			}
			userID := loginVals.Username
			password := loginVals.Password
			// user객체에 email 값셋팅후 DBselect
			user := models.Users{Email: userID}
			util.Engine.Get(&user)
			log.Println(user)
			// select한 user객체의 admin값에따라 isadmin true false 설정
			var isAdmin bool
			if user.Admin == 0 {
				isAdmin = false
			} else {
				isAdmin = true
			}

			//  bcrypt로 암호화된 패스워드와 사용자가입력한 패스워드가 일치하는지 확인후 일치하면 user객체 리턴
			if pwdErr := bcrypt.CompareHashAndPassword([]byte(user.Pwd), []byte(password)); pwdErr == nil {
				log.Println("로그인성공")
				return &User{
					UserName: user.Name,
					CID:      user.CustID,
					UID:      user.UserID,
					IsAdmin:  isAdmin,
				}, nil
			}
			log.Println("아이디 혹은 비밀번호오류")
			return nil, jwt.ErrFailedAuthentication
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			// if v, ok := data.(*User); ok && v.IsAdmin == true {
			// 	return true
			// }

			// return false
			return true
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		// TokenLookup is a string in the form of "<source>:<name>" that is used
		// to extract token from the request.
		// Optional. Default value "header:Authorization".
		// Possible values:
		// - "header:<name>"
		// - "query:<name>"
		// - "cookie:<name>"
		// - "param:<name>"
		TokenLookup: "header: Authorization, query: token, cookie: jwt",
		// TokenLookup: "query:token",
		// TokenLookup: "cookie:token",

		// TokenHeadName is a string in the header. Default value is "Bearer"
		TokenHeadName: "Bearer",

		// TimeFunc provides the current time. You can override it to use another time value. This is useful for testing or if your server uses a different time zone than your tokens.
		TimeFunc: time.Now,
	})

	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}
	log.Println("init 끝")
	log.Println(AuthMiddleware)
}