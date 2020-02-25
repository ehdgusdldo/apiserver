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

// User payload 셋팅용 struct
type User struct {
	UserName string
	CID      int
	UID      int
	IsAdmin  bool
}

// 아이디 패스워드 바인딩용 struct
type login struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

// 스웨거 페이지보여주기위한 struct
type loginSuccess struct {
	code   string `example:"200"`
	expire string `example:"2020-02-28T09:15:29+09:00"`
	token  string `example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJDSUQiOjIsIklzQWRtaW4iOnRydWUsIlVJRCI6NDgsIlVzZXJOYW1lIjoi6rmA7ZiV6re8IiwiZXhwIjoxNTgyODQ4OTI5LCJvcmlnX2lhdCI6MTU4MjI0NDEyOX0.ErN3ajVe7Sj5s6jaJQFGfMMaeP8jietm9uP0feacfxA"`
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
		Realm: "test zone",
		//SigningAlgorithm : HS256, 해싱알고리즘 default값 hs256
		//서명부 시크릿키
		Key: []byte(secretKey),
		//토큰 유효기간 설정
		Timeout:    time.Hour * 24 * 7,
		MaxRefresh: time.Hour * 24 * 7,
		// identityKey
		IdentityKey: identityKey,
		//PayloadFunc 셋팅 암호화되지않음
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
		//
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
		Authenticator: loginfunc,
		//true false로 admin권한여부정할수있는듯? 일단 true리턴
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
		TokenLookup: "header: Authorization", //, query: token, cookie: jwt
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

// 로그인핸들러 ..
// @Description 로그인후 jwt토큰발행
// @Router /login [post]
// @Tags login
// @Accept  json
// @Produce  json
// @Param user body login true "아이디 및 패스워드"
// @Success 200 {object} loginSuccess
func loginfunc(c *gin.Context) (interface{}, error) {
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
}
