package router

import (
	"immersiveApp/app/middlewares"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitRouter(db *gorm.DB, e *echo.Echo) {
	middlewares.BasicLogger(e)
	AuthRouter(db, e)
	UserRouter(db, e)
	TeamRouter(db, e)
	ClassRouter(db, e)
	MenteeRouter(db, e)
	StatusRouter(db, e)
}
