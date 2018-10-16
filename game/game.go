package game

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

var testCount float64

func (g *Game) Run() {
	testCount++
	if testCount > 60 {
		g.Shuffle()
		g.Deal()
		testCount = 0
	}
}

type MainGameState int

const (
	GameDEMO MainGameState = iota
	GameSTART
	GamePLAY
	GameCHECK
	GameWIN
	GameLOSE
	GameACCOUNT
)

var GameMessage = [...]string{
	"GAME DEMO",
	"GAME START",
	"GAME PLAY",
	"GAME CHECK",
	"GAME WIN",
	"GAME LOSE",
	"GAME ACCOUNT",
}
