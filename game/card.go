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
	Pattern int
	Number  int
	Val     int
}

type Cards struct {
	Sets         []Card
	CurrentIndex int
}

func NewCardSet() *Cards {
	Sets := make([]Card, 52)
	for i := 0; i < len(Sets); i++ {
		Sets[i].Pattern = i / 13
		Sets[i].Number = i % 13
		Sets[i].Val = i
	}
	return &Cards{Sets, 0}
}

func (card *Cards) cardSwap(i, j int) {
	card.Sets[i].Val, card.Sets[j].Val = card.Sets[j].Val, card.Sets[i].Val
	card.Sets[i].Pattern, card.Sets[j].Pattern = card.Sets[j].Pattern, card.Sets[i].Pattern
	card.Sets[i].Number, card.Sets[j].Number = card.Sets[j].Number, card.Sets[i].Number
}

func (card *Cards) DeckCard() {
	rand.Shuffle(len(card.Sets), card.cardSwap)
}

func (card *Cards) ShowCard(num int) {
	println(Cardnumber[card.Sets[num].Number] + " of " + CardPattern[card.Sets[num].Pattern])
}

func (card *Cards) GetNextCard() Card {
	c := card.Sets[card.CurrentIndex]
	card.CurrentIndex++
	if card.CurrentIndex >= 51 {
		card.DeckCard()
		card.CurrentIndex = 0
	}
	return c
}
