package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/ehdgusdldo/APIServer/docs"
	"github.com/ehdgusdldo/APIServer/routers"
	"github.com/joho/godotenv"
)

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	//swagger
	// programatically set swagger info
	docs.SwaggerInfo.Title = "LPWA API"
	docs.SwaggerInfo.Description = "LPWA API SAMPLE SWAGGER"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8082"
	// docs.SwaggerInfo.BasePath = "/v2"
	// env 라이브러리 호출
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	port := os.Getenv("PORT")

	fmt.Println("start API Server LPWA with port :: " + port)

	router := routers.InitRouter()

	server := &http.Server{
		Addr:           port,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	server.ListenAndServe()

}
