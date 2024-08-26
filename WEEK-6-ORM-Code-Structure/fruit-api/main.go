package main

import (
	"fruit-api/controllers"
	"fruit-api/database"

	"github.com/labstack/echo/v4"
)

func main() {
    database.InitDB()
    database.MigrateDB()

    fruitController := controllers.InitFruitController()
    e := echo.New()

    e.GET("/api/v1/fruits", fruitController.GetAll)
    e.GET("/api/v1/fruits/:id", fruitController.GetByID)
    e.POST("/api/v1/fruits", fruitController.Create)
    e.PUT("/api/v1/fruits/:id", fruitController.Update)
    e.DELETE("/api/v1/fruits/:id", fruitController.Delete)

    e.Logger.Fatal(e.Start(":1323"))
}
