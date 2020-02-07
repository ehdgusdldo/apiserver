package util

import (
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/joho/godotenv"
)

var Engine *xorm.Engine

func init() {
	// env 라이브러리 호출
	enverr := godotenv.Load()
	if enverr != nil {
		log.Fatal("Error loading .env file")
	}

	var err error
	driverName := os.Getenv("driver-name")
	dbURL := os.Getenv("db-url")
	dbName := os.Getenv("db-name")
	dbUsername := os.Getenv("db-username")
	dbPassword := os.Getenv("db-password")
	datasource := dbUsername + ":" + dbPassword + "@tcp(" + dbURL + ")/" + dbName

	Engine, err = xorm.NewEngine(driverName, datasource)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	// engine.SetMaxOpenConns()
	// engine.SetMaxIdleConns()
	fmt.Println(Engine)

}
