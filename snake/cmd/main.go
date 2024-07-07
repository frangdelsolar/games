package main

import (
	s "github.com/frangdelsolar/games/snake"
)

var log *s.Logger

func main() {

	config, err := s.LoadConfig("config.yaml")
	if err != nil {
		panic(err)
	}

	log = s.NewLogger(config.LogLevel)

	log.Info().Interface("config", config).Msg("Initializing Snake with config")

	board := s.NewBoard(config)
	log.Debug().Interface("board", board).Msg("Created new board")
}
