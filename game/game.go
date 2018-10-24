package game

import (
	"Loowootoo/lopoker/sprlib"
	"Loowootoo/lopoker/ui2d/assets/pcard"

	"github.com/hajimehoshi/ebiten"
)

type Game struct {
	Player        *Player
	CardSet       *Cards
	SmokeAnim     *sprlib.Sprite
	Message       string
	GameStatus    MainGameState
	GameSubStatus int
}

func NewGame(credit int) *Game {
	player := NewPlayer(credit)
	cardSet := NewCardSet()
	smokeAnim := sprlib.NewSprite()
	smokeAnim.AddAnimFrameFromBytes("default", pcard.SmokePNG, 2000, 15, ebiten.FilterDefault)
	smokeAnim.CenterCoordonnates = true
	smokeAnim.Pos = sprlib.Vector{576, 153, 100}
	smokeAnim.Speed = 1
	smokeAnim.Direction = sprlib.Vector{0, 0, -1}
	smokeAnim.Start()

	return &Game{player, cardSet, smokeAnim, "", GameDEMO, 0}
}

func (g *Game) AddBet(bet int) {
	g.Player.Bet += bet
}

func (g *Game) Shuffle() {
	g.CardSet.DeckCard()
}

func (g *Game) Deal() {
	g.Player.ResetHeld()
	for i := 0; i < len(g.Player.Hand); i++ {
		g.Player.Hand[i] = g.CardSet.GetNextCard()
	}
	g.Player.sortHand()
}

func (g *Game) ShowPlayerCard() {
	g.Player.ShowPlayerCard()
	g.Player.ShowPlayerSortCard()
}

var testCount float64

func (g *Game) Run() {
	testCount++
	if testCount > 60 {
		g.Deal()
		g.Message = g.Player.CheckWin()
		testCount = 0
	}
	g.SmokeAnim.Update()
}

type MainGameState int

const (
	GameDEMO MainGameState = iota
	GameSTART
	GameBET
	GamePLAY
	GameCHECK
	GameWIN
	GameLOSE
	GameACCOUNT
)

var GameMessage = [...]string{
	"GAME DEMO",
	"GAME START",
	"GAME BET",
	"GAME PLAY",
	"GAME CHECK",
	"GAME WIN",
	"GAME LOSE",
	"GAME ACCOUNT",
}

func (g *Game) GameLoop() {
	switch g.GameStatus {
	case GameDEMO:
		game.DemoProc()
	case GameSTART:
		g.GameStatus = GameBET
	case GameBET:
		g.GameStatus = GamePLAY
	case GamePLAY:
	case GameCHECK:
	case GameWIN:
	case GameLOSE:
	case GameACCOUNT:
	}
}
