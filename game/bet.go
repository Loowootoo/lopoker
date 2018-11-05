package game

func (g *Game) GameBetProc() {
	if g.Player.Credit <= 0 && g.Player.Bet <= 0 {
		g.GameStatus = GameDEMO
		g.GameSubStatus = 0
		return
	}
	if g.IsBetKey() {
		g.Player.Credit--
		g.Player.Bet++
	}
	if g.Player.Bet > 0 && g.IsStartKey() {
		g.GameStatus = GamePLAY
		g.GameSubStatus = 0
	}
	switch g.GameSubStatus {
	case 0:
		if g.MsgCounter.TimeUp() {
			g.GameSubStatus++
		}
		g.Message = "PUSH    BET"
	case 1:
		if g.MsgCounter.TimeUp() {
			g.GameSubStatus = 0
		}
		if g.Player.Bet > 0 {
			g.Message = "PUSH  START"
		} else {
			g.Message = "           "
		}
	}
}
