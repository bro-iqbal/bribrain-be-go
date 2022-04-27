package controllers

import (
	"database/sql"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/bro-iqbal/bribrain-be-go/helpers"
	"github.com/bro-iqbal/bribrain-be-go/models"
	"github.com/bro-iqbal/bribrain-be-go/structs"
	echo "github.com/labstack/echo/v4"
)

func GetTargetKunjungan(c echo.Context) error {
	var kunjungan models.Kunjungan
	response := new(models.Response)

	userId, _ := strconv.Atoi(c.Param("id"))

	if c.Param("id") == "0" || c.Param("id")[0:1] == "0" {
		response.Error(http.StatusBadRequest, "UserID cannot be 0", nil)

		return c.JSON(http.StatusBadRequest, response)
	}

	if val := helpers.IsInt(c.Param("id")); val != true {
		response.Error(http.StatusBadRequest, "UserID must be numeric", nil)

		return c.JSON(http.StatusBadRequest, response)
	}

	db.Raw("SELECT akuisisi, agen_jawara, agen_juragan, agen_bep FROM kunjungans WHERE user_id = ?", userId).Scan(&kunjungan)

	kunjungan.UserId = userId

	response.Success(http.StatusOK, "Get Target Kunjungan by UserId successful", kunjungan)

	return c.JSON(http.StatusOK, response)
}

func StoreTargetKunjungan(c echo.Context) error {
	response := new(models.Response)
	var kunjungan models.Kunjungan
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
		"CreatedAt":   structs.NullTime{sql.NullTime{date, true}},
	}

	log.Println(user.InputStatusKunjungan)

	if user.InputStatusKunjungan == "Belum Input" {
		go CreateData(kunjungan, data)
	} else {
		go UpdateData(kunjungan, data, userId, yr)
	}

	go UpdateStatusUser(user, "kunjungan", c.Param("id"))

	response.Success(http.StatusOK, "Input Target Kunjungan successful", data)

	return c.JSON(http.StatusOK, response)
}

func CreateData(kunjungan models.Kunjungan, data interface{}) {
	db.Model(&kunjungan).Create(data)
}

func UpdateData(kunjungan models.Kunjungan, data interface{}, userId int, yr int) {
	db.Where("user_id = ? AND YEAR(created_at) = ?", userId, yr).Model(&kunjungan).Updates(data)
}
