package game

import (
	"math/rand"
)

var CardPattern = map[int]string{
	0: "Spades",
	1: "Hearts",
	2: "Diamonds",
	3: "Clubs",
	4: "Joker",
}

var Cardnumber = map[int]string{
	0:  "Ace",
	1:  "2",
	2:  "3",
	3:  "4",
	4:  "5",
	5:  "6",
	6:  "7",
	7:  "8",
	8:  "9",
	9:  "10",
	10: "Jack",
	11: "Queen",
	12: "King",
}

type Card struct {
	pattern int
	number  int
	val     int
}

type Cards struct {
	Sets         []Card
	CurrentIndex int
}

func NewCardSet() *Cards {
	Sets := make([]Card, 52)
	for i := 0; i < len(Sets); i++ {
		Sets[i].pattern = i / 13
		Sets[i].number = i % 13
		Sets[i].val = i
	}
	return &Cards{Sets, 0}
}

func (card *Cards) cardSwap(i, j int) {
	card.Sets[i].val, card.Sets[j].val = card.Sets[j].val, card.Sets[i].val
	card.Sets[i].pattern, card.Sets[j].pattern = card.Sets[j].pattern, card.Sets[i].pattern
	card.Sets[i].number, card.Sets[j].number = card.Sets[j].number, card.Sets[i].number
}

func (card *Cards) DeckCard() {
	rand.Shuffle(len(card.Sets), card.cardSwap)
}

func (card *Cards) ShowCard(num int) {
	println(Cardnumber[card.Sets[num].number] + " of " + CardPattern[card.Sets[num].pattern])
}

func (card *Cards) GetNextCard() Card {
	c := card.Sets[card.CurrentIndex]
	card.CurrentIndex++
	card.CurrentIndex %= 52
	return c
}
