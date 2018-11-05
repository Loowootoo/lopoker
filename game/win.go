package game

import (
	"strconv"
)

func (g *Game) GameWinProc() {
	switch g.GameSubStatus {
	case 0:
		if g.MsgCounter.TimeUp() {
			g.GameSubStatus++
			g.Message = g.WinStr + strconv.Itoa(g.GameWin)
		}
	case 1:
		g.Player.Credit += g.GameWin
		g.Player.Bet = 0
		g.GameStatus = GameACCOUNT
		g.GameSubStatus = 0
	}
}
