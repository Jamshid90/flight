package main

import (
	"log"
	"time"

	"github.com/Jamshid90/flight/api"
	configpkg "github.com/Jamshid90/flight/internal/config"
	"github.com/Jamshid90/flight/internal/database"
	"github.com/Jamshid90/flight/internal/flight"
	"github.com/Jamshid90/flight/internal/http/server"
	"go.uber.org/zap"

	loggerpkg "github.com/Jamshid90/flight/internal/logger"
)

func main() {
	var (
		contextTimeout time.Duration
	)

	// initialization config
	config := configpkg.New()

	// initialization logger
	logger, err := loggerpkg.New(config.LogLevel, config.Environment)
	if err != nil {
		log.Fatal(err)
	}
	defer logger.Sync()

	db, err := database.New(config)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	// context timeout initialization
	contextTimeout, err = time.ParseDuration(config.Context.Timeout)
	if err != nil {
		log.Fatalf("Error during parse duration for context timeout : %v\n", err)
	}

	// repositories initialization
	flightRepo := flight.NewRepository(db)
	flightUsecase := flight.NewUsecase(contextTimeout, flightRepo)

	// api handler initialization
	apiHandler := api.New(api.Option{
		Logger:         logger,
		FlightUsecase:  flightUsecase,
		ContextTimeout: contextTimeout,
	})

	logger.Info("Listen: ", zap.String("address", config.Server.Host+config.Server.Port))
	log.Fatal(server.NewServer(config, apiHandler).Run())
}
