package game

func (g *Game) DemoProc() {
	if g.Player.Credit > 0 {
		g.GameStatus = GameSTART
		return
	}
	switch g.GameSubStatus {
	case 0:
		if g.MsgCounter.TimeUp() {
			g.GameSubStatus++
		}
		g.Message = "INSERT COIN"
	case 1:
		if g.MsgCounter.TimeUp() {
			g.GameSubStatus = 0
		}
		g.Message = "           "
	}
	g.SmokeAnim.Update()
}
