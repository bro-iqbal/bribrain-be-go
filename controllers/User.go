package controllers

import (
	"net/http"

	"github.com/bro-iqbal/bribrain-be-go/helpers"
	"github.com/bro-iqbal/bribrain-be-go/models"
	echo "github.com/labstack/echo/v4"
)

func GetUserByID(c echo.Context) error {
	var user models.User
	response := new(models.Response)

	userId := c.Param("id")

	if userId == "0" || userId[0:1] == "0" {
		response.Error(http.StatusBadRequest, "UserID cannot be 0", nil)

		return c.JSON(http.StatusBadRequest, response)
	}

	if val := helpers.IsInt(userId); val != true {
		response.Error(http.StatusBadRequest, "UserID must be numeric", nil)

		return c.JSON(http.StatusBadRequest, response)
	}

	user = QueryUser(user, userId)

	response.Success(http.StatusOK, "Get User By ID successful", user)

	return c.JSON(http.StatusOK, response)
}

func QueryUser(user models.User, userId string) models.User {

	db.Raw("SELECT id, name, input_status_kunjungan, input_status_rka FROM users WHERE id = ?", userId).Scan(&user)

	return user
}

func UpdateStatusUser(user models.User, source string, userId string) {

	var data map[string]interface{}

	if source == "kunjungan" {
		data = map[string]interface{}{
			"InputStatusKunjungan": "Sudah Input",
		}
	} else {
		data = map[string]interface{}{
			"InputStatusRKA": "Sudah Input",
		}
	}

	db.Where("id = ?", userId).Model(&user).Updates(data)
}
