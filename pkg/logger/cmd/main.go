package main

import (
	"fmt"

	"github.com/frangdelsolar/go-games/pkg/logger"
)

func main(){

	log := logger.NewLogger("debug", "Logger PKG Main", "1.0.0")

	log.Info().Msg("This is a test of Info log")

	log.Warn().Msg("This is a test of Warn log")


	err:= fmt.Errorf("This is a fake error for testing purposes")
	
	log.Err(err).Msg("This is a test of Error log")

}