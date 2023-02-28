package main

import (
	"flag"
	"os"

	"github.com/pauljamescleary/gomin/pkg/common/handler"
	router "github.com/pauljamescleary/gomin/pkg/common/router"
	"github.com/rs/zerolog/log"
)

func main() {
	configPath := flag.String("configpath", "", "Config Path")
	flag.Parse()
	if configPath == nil || len(*configPath) == 0 {
		log.Fatal().Msgf("Unable to load config path. Empty Path specified. %s", *configPath)
	}
	if _, err := os.Stat(*configPath); os.IsNotExist(err) {
		// path/to/whatever does not exist
		log.Fatal().Msgf("Unable to load config path. Path not found. %s", *configPath)
	}
	h := handler.LoadHandler(configPath)
	e := router.SetupRouter(h)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
