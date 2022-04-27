package routes

import (
	"github.com/bro-iqbal/bribrain-be-go/controllers"
	echo "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Route(e *echo.Echo) {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// config := middleware.JWTConfig{
	// 	SigningKey: []byte(os.Getenv("JWT_SECRET")),
	// }

	e.GET("/", controllers.Documentation)

	e.GET("/user/:id", controllers.GetUserByID)

	e.GET("/target-kunjungan/:id", controllers.GetTargetKunjungan)

	e.GET("/rka-last-year/:id", controllers.GetRKALastYear)
	e.GET("/target-rka/:id", controllers.GetTargetRKA)

	e.POST("/target-kunjungan/:id", controllers.StoreTargetKunjungan)
	e.POST("/target-rka/:id", controllers.StoreTargetRKA)

}
