package game

func (g *Game) getHeld() {
	if g.IsHeldKey1() {
		g.Player.Held[0] = !g.Player.Held[0]
	}
	if g.IsHeldKey2() {
		g.Player.Held[1] = !g.Player.Held[1]
	}
	if g.IsHeldKey3() {
		g.Player.Held[2] = !g.Player.Held[2]
	}
	if g.IsHeldKey4() {
		g.Player.Held[3] = !g.Player.Held[3]
	}
	if g.IsHeldKey5() {
		g.Player.Held[4] = !g.Player.Held[4]
	}
}
func (g *Game) GamePlayProc() {
	if g.Player.Credit <= 0 && g.Player.Bet <= 0 {
		g.GameStatus = GameDEMO
		g.GameSubStatus = 0
		return
	}
	switch g.GameSubStatus {
	case 0:
		if g.MsgCounter.TimeUp() {
			g.GameSubStatus++
			g.Player.SetBackCard(0, false)
		}
	case 1:
		if g.MsgCounter.TimeUp() {
			g.GameSubStatus++
			g.Player.SetBackCard(1, false)
		}
	case 2:
		if g.MsgCounter.TimeUp() {
			g.GameSubStatus++
			g.Player.SetBackCard(2, false)
		}
	case 3:
		if g.MsgCounter.TimeUp() {
			g.GameSubStatus++
			g.Player.SetBackCard(3, false)
		}
	case 4:
		if g.MsgCounter.TimeUp() {
			g.GameSubStatus++
			g.Player.SetBackCard(4, false)
			g.Player.CheckWin()
		}
	case 5:
		if g.MsgCounter.TimeUp() {
			g.GameSubStatus++
			g.Message = "SELECT HELD"
		}
	case 6:
		g.getHeld()
		if g.IsStartKey() {
			g.GameSubStatus++
		}
	case 7:
		g.DealWithHeld()
		g.GameSubStatus++
	case 8:
		if g.MsgCounter.TimeUp() {
			g.GameSubStatus++
		}
	case 9:
		if g.Player.BackCard[0] == false {
			g.GameSubStatus++
		}
		if g.MsgCounter.TimeUp() {
			g.GameSubStatus++
			g.Player.SetBackCard(0, false)
		}
	case 10:
		if g.Player.BackCard[1] == false {
			g.GameSubStatus++
		}
		if g.MsgCounter.TimeUp() {
			g.GameSubStatus++
			g.Player.SetBackCard(1, false)
		}
	case 11:
		if g.Player.BackCard[2] == false {
			g.GameSubStatus++
		}
		if g.MsgCounter.TimeUp() {
			g.GameSubStatus++
			g.Player.SetBackCard(2, false)
		}
	case 12:
		if g.Player.BackCard[3] == false {
			g.GameSubStatus++
		}
		if g.MsgCounter.TimeUp() {
			g.GameSubStatus++
			g.Player.SetBackCard(3, false)
		}
	case 13:
		if g.Player.BackCard[4] == false {
			g.GameSubStatus++
		}
		if g.MsgCounter.TimeUp() {
			g.GameSubStatus++
			g.Player.SetBackCard(4, false)
		}
	case 14:
		g.GameStatus = GameCHECK
		g.GameSubStatus = 0
	}
}
