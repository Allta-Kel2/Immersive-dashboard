package main

import (
	"immersiveApp/app/configs"
	"immersiveApp/app/database"
	"immersiveApp/app/router"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	cfg := configs.InitConfig()
	db := database.InitDBMysql(*cfg)
	database.InitMigration(db)

	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	router.InitRouter(db, e)
	e.Logger.Fatal(e.Start(":8080"))
}
