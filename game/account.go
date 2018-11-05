package game

func (g *Game) GameAccountProc() {
	switch g.GameSubStatus {
	case 0:
		if g.MsgCounter.TimeUp() {
			g.GameSubStatus++
			g.Message = "Accounting..."
		}
	case 1:
		g.GameStatus = GameDEMO
		g.GameSubStatus = 0
	}
}
