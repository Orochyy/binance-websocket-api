package main

import (
	"github.com/rs/zerolog"
	zlog "github.com/rs/zerolog/log"
)

func main() {

}

func testlog() {
	debug := false
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}
}

func debug() {
	debug := false
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	zlog.Info().
		Str("service", "my-service").
		Int("Some integer", 10).
		Msg("Hello")
	// Журнал отладки
	zlog.Debug().Msg("Exiting Program")
}
