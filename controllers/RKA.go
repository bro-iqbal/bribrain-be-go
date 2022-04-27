package controllers

import (
	"database/sql"
	"net/http"
	"strconv"
	"time"

	"github.com/bro-iqbal/bribrain-be-go/helpers"
	"github.com/bro-iqbal/bribrain-be-go/models"
	"github.com/bro-iqbal/bribrain-be-go/structs"
	echo "github.com/labstack/echo/v4"
)

func GetTargetRKA(c echo.Context) error {
	var rka models.Rka
	response := new(models.Response)

	if c.Param("id") == "0" || c.Param("id")[0:1] == "0" {
		response.Error(http.StatusBadRequest, "UserID cannot be 0", nil)

		return c.JSON(http.StatusBadRequest, response)
	}

	if val := helpers.IsInt(c.Param("id")); val != true {
		response.Error(http.StatusBadRequest, "UserID must be numeric", nil)

		return c.JSON(http.StatusBadRequest, response)
	}

	userId, _ := strconv.Atoi(c.Param("id"))

	t := time.Now()

	yr := t.Year()

	db.Raw("SELECT akuisisi, agen_jawara, agen_juragan,agen_bep FROM rkas WHERE type = 'target' AND user_id = ? AND YEAR(created_at) = ? ", userId, yr).Scan(&rka)

	rka.UserId = userId

	response.Success(http.StatusOK, "Get Target RKA successful", rka)

	return c.JSON(http.StatusOK, response)
}

func GetRKALastYear(c echo.Context) error {
	var rka models.Rka
	response := new(models.Response)

	if c.Param("id") == "0" || c.Param("id")[0:1] == "0" {
		response.Error(http.StatusBadRequest, "UserID cannot be 0", nil)

		return c.JSON(http.StatusBadRequest, response)
	}

	if val := helpers.IsInt(c.Param("id")); val != true {
		response.Error(http.StatusBadRequest, "UserID must be numeric", nil)

		return c.JSON(http.StatusBadRequest, response)
	}

	userId, _ := strconv.Atoi(c.Param("id"))

	t := time.Now()

	yr := t.Year()

	db.Raw("SELECT akuisisi, agen_jawara, agen_juragan,agen_bep FROM rkas WHERE type = 'position' AND user_id = ? AND YEAR(created_at) = ?", userId, (yr - 1)).Scan(&rka)

	rka.UserId = userId

	response.Success(http.StatusOK, "Get RKA Last year successful", rka)

	return c.JSON(http.StatusOK, response)
}

func StoreTargetRKA(c echo.Context) error {
	response := new(models.Response)
	var rka models.Rka
	var user models.User

	if c.Param("id") == "0" || c.Param("id")[0:1] == "0" {
		response.Error(http.StatusBadRequest, "UserID cannot be 0", nil)

		return c.JSON(http.StatusBadRequest, response)
	}

	if val := helpers.IsInt(c.Param("id")); val != true {
		response.Error(http.StatusBadRequest, "UserID must be numeric", nil)

		return c.JSON(http.StatusBadRequest, response)
	}

	Akuisisi, _ := strconv.Atoi(c.FormValue("Akuisisi"))
	AgenJawara, _ := strconv.Atoi(c.FormValue("AgenJawara"))
	AgenJuragan, _ := strconv.Atoi(c.FormValue("AgenJuragan"))
	AgenBEP, _ := strconv.Atoi(c.FormValue("AgenBEP"))
	userId, _ := strconv.Atoi(c.Param("id"))

	user = QueryUser(user, c.Param("id"))

	t := time.Now()

	yr := t.Year()

	date, _ := time.Parse("2006-01-02 15:04:05", time.Now().Local().Format("2006-01-02 15:04:05"))

	data := map[string]interface{}{
		"Akuisisi":    Akuisisi,
		"AgenJawara":  AgenJawara,
		"AgenJuragan": AgenJuragan,
		"AgenBEP":     AgenBEP,
		"Status":      "active",
		"UserId":      userId,
		"Type":        "target",
		"CreatedAt":   structs.NullTime{sql.NullTime{date, true}},
	}

	if user.InputStatusRKA == "Belum Input" {
		go db.Model(&rka).Create(data)
	} else {
		go db.Where("user_id = ? AND YEAR(created_at) = ?", userId, yr).Model(&rka).Updates(data)
	}

	go UpdateStatusUser(user, "rka", c.Param("id"))

	response.Success(http.StatusOK, "Input Target RKA successful", map[string]interface{}{
		"TargetBulanBerjalan": (Akuisisi + AgenJawara + AgenJuragan + AgenBEP) / 12,
	})

	return c.JSON(http.StatusOK, response)
}
