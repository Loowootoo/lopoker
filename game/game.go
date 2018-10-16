package game

type Game struct {
	Player        *Player
	CardSet       *Cards
	Message       string
	GameStatus    MainGameState
	GameSubStatus int
}

func NewGame(credit int) *Game {
	player := NewPlayer(credit)
	cardSet := NewCardSet()
	return &Game{player, cardSet, "", GameDEMO, 0}
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
		g.GameStatus = GameSTART
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
