package game

func (g *Game) GameStartProc() {
	g.Deal()
	g.GameStatus = GameBET
	g.GameSubStatus = 0
}
