package main

import (
	"fmt"
	"os"

	"github.com/bro-iqbal/bribrain-be-go/controllers"
	"github.com/bro-iqbal/bribrain-be-go/routes"
	"github.com/joho/godotenv"
	echo "github.com/labstack/echo/v4"
)

func main() {
	err := godotenv.Load()

	fmt.Println(err)

	e := echo.New()
	controllers.Init()

	routes.Route(e)

	e.Logger.Fatal(e.Start(":" + os.Getenv("APP_PORT")))
}
