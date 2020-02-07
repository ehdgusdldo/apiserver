package routers

import (
	"database/sql"

	"github.com/gin-gonic/gin"

	ginSwagger "github.com/swaggo/gin-swagger"
	swaggerFiles "github.com/swaggo/gin-swagger/swaggerFiles"

	"github.com/ehdgusdldo/APIServer/controller"
)

var db *sql.DB

// InitRouter
func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(gin.Logger())

	//연결테스트
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	r.GET("/login", controller.Login)
	r.POST("/loginyou", controller.LoginYOU)
	//swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// custom

	c := r.Group("/customs")
	{
		c.GET("", controller.GetCustomAll)
		c.GET("/:id", controller.GetCustom)
		c.POST("", controller.RegisterCustom)
	}

	u := r.Group("/users")
	{
		u.GET("", controller.GetUserAll)
		u.POST("", controller.RegisterUser)
		u.GET("/:id", controller.GetUser)

	}

	return r
}
