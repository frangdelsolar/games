package snake

import (
	"os"
	"strconv"
)

type Board struct {
	Rows int
	Cols int
}

func NewBoard() (*Board, error) {
	var err error

	rows, err := strconv.Atoi(os.Getenv("ROWS"))
	if err != nil {
		rows = 20
	}

	cols, err := strconv.Atoi(os.Getenv("COLS"))
	if err != nil {
		cols = 10
	}

	return &Board{rows, cols}, err
}
