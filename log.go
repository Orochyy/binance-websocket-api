package main

import (
	"github.com/rs/zerolog"
	zlog "github.com/rs/zerolog/log"
)

func main() {
	testlog()
}

func testlog() {
	debug := false
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	zlog.Info().
		Str("service", "stream-coin-cap").
		Int("BTC", 0).
		Msg("fine")
}
