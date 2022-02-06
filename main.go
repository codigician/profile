package main

import (
	"context"
	"log"
	"os"

	"github.com/codigician/profile/internal"
	"github.com/codigician/profile/internal/about"
	aboutmongo "github.com/codigician/profile/internal/about/mongo"
	"github.com/codigician/profile/internal/analytics"
	"github.com/codigician/profile/internal/submission"
	"github.com/labstack/echo/v4"
)

const _address = ":8082"

func main() {
	echoServer := echo.New()

	mongoURL := os.Getenv("MONGO_DB_URL")
	aboutMongoRepository := aboutmongo.New(mongoURL)
	if err := aboutMongoRepository.Connect(context.Background()); err != nil {
		log.Fatalf("about mongo repository: %v\n", err)
	}

	aboutService := about.NewService(aboutMongoRepository)
	submissionService := submission.NewService()
	analyticsService := analytics.NewService()

	profileHandler := internal.NewProfileHandler(aboutService, submissionService, analyticsService)
	profileHandler.RegisterRoutes(echoServer)

	log.Fatal(echoServer.Start(_address))
}
