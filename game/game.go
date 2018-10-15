package game

import (
	"fmt"
)

type Game struct {
	Player  *Player
	CardSet *Cards
}

func NewGame(credit int) *Game {
	player := NewPlayer(credit)
	cardSet := NewCardSet()
	return &Game{player, cardSet}
}

func (g *Game) AddBet(bet int) {
	g.Player.Bet += bet
}

func (g *Game) Shuffle() {
	g.CardSet.DeckCard()
}

func (g *Game) Deal() {
	for i := 0; i < len(g.Player.Hand); i++ {
		g.Player.Hand[i] = g.CardSet.GetNextCard()
	}
	g.Player.sortHand()
}

func (g *Game) ShowPlayerCard() {
	g.Player.ShowPlayerCard()
	g.Player.ShowPlayerSortCard()
}

func (g *Game) Run() {
	g.Shuffle()
	g.Deal()
	g.ShowPlayerCard()
	fmt.Println(g.Player.CheckWin())
}
