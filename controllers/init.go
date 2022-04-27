package controllers

import (
	"net/http"

	"github.com/bro-iqbal/bribrain-be-go/database"
	echo "github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init() {
	db = database.Connection()
}

func Documentation(c echo.Context) error {
	return c.Redirect(http.StatusMovedPermanently, "https://documenter.getpostman.com/view/1585979/UyrDEbaB")
}
