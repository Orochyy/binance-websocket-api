package main

import (
	"github.com/rs/zerolog"
	zlog "github.com/rs/zerolog/log"
)

func main() {

	//log.Print("Hello")
	//zlog.Print("Hello")
	//
	//zlog.Info().
	//	Str("service", "my-service").
	//	Int("Some integer", 10).
	//	Msg("Hello")

	debug()
}

func debug() {
	debug := false
	// Применяем уровень ведения журнала в начале приложения
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}
	// Лог с именем сервиса
	zlog.Info().
		Str("service", "my-service").
		Int("Some integer", 10).
		Msg("Hello")
	// Журнал отладки
	zlog.Debug().Msg("Exiting Program")
}
