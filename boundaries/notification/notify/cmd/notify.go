/*
Bot application
*/
package main

import (
	"os"

	"github.com/spf13/viper"

	notify_di "github.com/shortlink-org/shortlink/boundaries/notification/notify/di"
	"github.com/shortlink-org/shortlink/pkg/graceful_shutdown"
)

func main() {
	viper.SetDefault("SERVICE_NAME", "shortlink-notify")

	// Init a new service
	s, cleanup, err := notify_di.InitializeFullBotService()
	if err != nil { // TODO: use as helpers
		panic(err)
	}

	defer func() {
		if r := recover(); r != nil {
			s.Log.Error(r.(string)) //nolint:forcetypeassert // simple type assertion
		}
	}()

	// Handle SIGINT, SIGQUIT and SIGTERM.
	graceful_shutdown.GracefulShutdown()

	// Stop the service gracefully.
	cleanup()

	// Exit Code 143: Graceful Termination (SIGTERM)
	os.Exit(143) //nolint:gocritic // TODO: research
}
