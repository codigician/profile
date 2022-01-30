package main

import (
	"log"

	"github.com/codigician/profile/internal"
	"github.com/codigician/profile/internal/about"
	aboutmongo "github.com/codigician/profile/internal/about/mongo"
	"github.com/codigician/profile/internal/analytics"
	"github.com/codigician/profile/internal/submission"
	"github.com/labstack/echo/v4"
)

const _address = ":8081"

func main() {
	echoServer := echo.New()

	aboutMongoRepository := aboutmongo.New("mongodb://localhost:27017")
	aboutService := about.NewService(aboutMongoRepository)
	submissionService := submission.NewService()
	analyticsService := analytics.NewService()

	profileHandler := internal.NewProfileHandler(aboutService, submissionService, analyticsService)
	profileHandler.RegisterRoutes(echoServer)

	log.Fatal(echoServer.Start(_address))
}
