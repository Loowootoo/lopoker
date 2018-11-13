package game

func (g *Game) GameWinProc() {
	switch g.GameSubStatus {
	case 0:
		if g.MsgCounter.TimeUp() {
			g.GameSubStatus++
			g.Message = g.WinStr
		}
	case 1:
		if g.WinCounter.TimeUp() {
			g.Player.Credit++
			g.GameWin--
			g.Sound.Coin.Rewind()
			g.Sound.Coin.Play()
			if g.GameWin == 0 {
				g.GameSubStatus++
			}
		}
	case 2:
		g.Player.Bet = 0
		g.GameStatus = GameACCOUNT
		g.GameSubStatus = 0
	}
}
