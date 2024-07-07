package main

import (
	s "github.com/frangdelsolar/games/snake"
)

var log *s.Logger

func main() {
	log = s.NewLogger("debug")

	log.Info().Msg("Initializing Snake!")
}
