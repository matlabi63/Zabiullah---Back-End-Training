package main

import (
	"WEEK-6-ORM-Code-Structure/controllers"
	"WEEK-6-ORM-Code-Structure/database"
	"WEEK-6-ORM-Code-Structure/middlewares"

	echojwt "github.com/labstack/echo-jwt/v4"

	"github.com/labstack/echo/v4"
)

func main() {

	database.InitDB()

	database.MigrateDB()
	e := echo.New()
	loggerConfig := middlewares.LoggerConfig{
		Format: "[${time_rfc3339}] ${status} ${method} ${host} ${path} ${latency_human}" + "\n",
	}
	loggerMiddleware := loggerConfig.Init()
	e.Use(loggerMiddleware)
	jwtConfig := middlewares.JWTConfig{
		SecretKey: "secret",
	}
	authMiddlewareConfig := jwtConfig.Init()
	userController := controllers.InitUserController()
	packageController := controllers.InitPackageController()
	e.POST("/api/v1/register", userController.Register)
	e.POST("/api/v1/login", userController.Login)

	packageRoutes := e.Group("/api/v1", echojwt.WithConfig(authMiddlewareConfig), middlewares.VerifyToken)

	packageRoutes.GET("/packages", packageController.GetAll)
	packageRoutes.GET("/packages/:id", packageController.GetByID)
	packageRoutes.POST("/packages", packageController.Create)
	packageRoutes.PUT("/packages/:id", packageController.Update)
	packageRoutes.DELETE("/packages/:id", packageController.Delete)

	e.Logger.Fatal(e.Start(":1323"))

}