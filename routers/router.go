package routers

import (
	"database/sql"
	"log"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"

	ginSwagger "github.com/swaggo/gin-swagger"
	swaggerFiles "github.com/swaggo/gin-swagger/swaggerFiles"

	"github.com/ehdgusdldo/APIServer/controller"
	"github.com/ehdgusdldo/APIServer/mid"
)

var db *sql.DB

// InitRouter router Init
func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(gin.Logger())

	//연결테스트
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	//swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// 404
	r.NoRoute(mid.AuthMiddleware.MiddlewareFunc(), func(c *gin.Context) {
		claims := jwt.ExtractClaims(c)
		log.Printf("NoRoute claims: %#v\n", claims)
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	// custom
	c := r.Group("/customs")
	{
		c.GET("", controller.GetCustomAll)
		c.GET("/:id", controller.GetCustom)
		// c.POST("", controller.RegisterCustom)
	}
	// users
	u := r.Group("/users")
	{
		u.GET("", controller.GetUserAll)
		// u.POST("", controller.RegisterUser)
		u.GET("/:id", controller.GetUser)

	}

	auth := r.Group("/auth")
	// jwt 미들웨어 적용 전체라우터에할수도있고 그룹별로 따로도 가능
	// 토큰발행 로그인
	r.POST("/login", mid.AuthMiddleware.LoginHandler)
	// refresh token 발행
	r.GET("/refresh_token", mid.AuthMiddleware.RefreshHandler)

	auth.Use(mid.AuthMiddleware.MiddlewareFunc())
	{
		//test
		auth.GET("/hello", controller.HelloHandler)
	}

	// influx select
	r.GET("/equip", controller.Equip)
	r.GET("/site", controller.Site)
	r.GET("/event", controller.Event)

	return r
}
