package router

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	_menteeData "immersiveApp/features/mentees/data"
	_menteeHandler "immersiveApp/features/mentees/delivery"
	_menteeService "immersiveApp/features/mentees/service"
)

func MenteeRouter(db *gorm.DB, e *echo.Echo) {
	data := _menteeData.New(db)
	service := _menteeService.New(data)
	handler := _menteeHandler.New(service)

	g := e.Group("/mentees")
	g.GET("", handler.GetAll)
	g.GET("/:id", handler.GetById)
	g.POST("", handler.Create)
	g.PUT("/:id", handler.Update)
	g.DELETE("/:id", handler.Delete)
}