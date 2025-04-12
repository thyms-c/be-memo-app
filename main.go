package main

import (
	"context"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/thyms-c/be-memo-app/internal/configs"
	"github.com/thyms-c/be-memo-app/internal/handlers"
	"github.com/thyms-c/be-memo-app/internal/infrastructure/database"
	"github.com/thyms-c/be-memo-app/internal/repositories"
	"github.com/thyms-c/be-memo-app/internal/services"
)

func main() {
	app := echo.New()

	configs := configs.NewConfig()
	ctx := context.Background()

	app.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{configs.FrontendOrigin},
		AllowHeaders: []string{
			echo.HeaderOrigin,
			echo.HeaderContentType,
			echo.HeaderAccept,
			echo.HeaderAuthorization,
		},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE, echo.OPTIONS},
	}))

	mongoClient := database.NewMongoClient(configs, ctx)

	counterRepository := repositories.NewCounterRepository(configs, mongoClient)
	memoRepository := repositories.NewMemoRepository(configs, mongoClient)

	counterService := services.NewCounterService(counterRepository)
	memoService := services.NewMemoService(counterRepository, memoRepository)

	counterHandler := handlers.NewCounterHandler(counterService)
	memoHandler := handlers.NewMemoHandler(memoService)

	app.GET("/v1/memos", memoHandler.GetAllMemos)
	app.POST("/v1/memos", memoHandler.CreateMemo)
	app.GET("/v1/memos/user-type", memoHandler.GetMemoByUserType)

	app.GET("/v1/counters/user-type", counterHandler.GetCounter)

	app.Logger.Fatal(app.Start(":1323"))

}
