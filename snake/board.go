package snake

type Board struct {
	Rows int
	Cols int
}

func NewBoard(*AppConfig) *Board {

	return &Board{
		Rows: appConfig.Board.Rows,
		Cols: appConfig.Board.Cols,
	}
}
