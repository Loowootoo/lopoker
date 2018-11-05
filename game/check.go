package game

var OddsTable = map[string]int{
	"ROYAL FLUSH":     800,
	"STRAIGHT FLUSH":  50,
	"FOUR OF A KIND":  25,
	"FULL HOUSE":      9,
	"FLUSH":           6,
	"STRAIGHT":        4,
	"THREE OF A KIND": 3,
	"TWO PAIR":        2,
	"JACK OR BETTER":  1,
	"NONE":            0,
}

func (g *Game) GameCheckProc() {
	switch g.GameSubStatus {
	case 0:
		g.WinStr = g.Player.CheckWin()
		g.GameWin = OddsTable[g.WinStr] * g.Player.Bet
		g.GameSubStatus++
	case 1:
		if g.GameWin > 0 {
			g.GameStatus = GameWIN
			g.GameSubStatus = 0
		} else {
			g.GameStatus = GameLOSE
			g.GameSubStatus = 0
		}
	}
}
