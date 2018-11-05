package game

func (g *Game) GameLoseProc() {
	switch g.GameSubStatus {
	case 0:
		if g.MsgCounter.TimeUp() {
			g.GameSubStatus++
			g.Message = "YOU LOSE"
		}
	case 1:
		g.Player.Bet = 0
		g.GameStatus = GameACCOUNT
		g.GameSubStatus = 0
	}
}
