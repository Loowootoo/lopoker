package game

import (
	"fmt"
)

type Player struct {
	Hand     []Card
	handSort []Card
	Credit   int
	Bet      int
}

func NewPlayer(credit int) *Player {
	hand := make([]Card, 5)
	handSort := make([]Card, 5)
	return &Player{hand, handSort, credit, 0}
}

var OddsTable = map[string]int{
	"Royal Flush":     800,
	"Straight Flush":  50,
	"Four of a Kind":  25,
	"Full House":      9,
	"Flush":           6,
	"Straight":        4,
	"Three of a Kind": 3,
	"Two Pair":        2,
	"Jack or Better":  1,
	"None":            0,
}

const (
	Spades   = 0
	Hearts   = 1
	Diamonds = 2
	Clubs    = 3
	Joker    = 4
)

func (p *Player) sortHand() {
	num := len(p.Hand)
	for i := 0; i < num; i++ {
		p.handSort[i] = p.Hand[i]
		if p.handSort[i].number == 0 {
			p.handSort[i].number = 13
		}
	}
	// first sort number
	for i := 0; i < num-1; i++ {
		isChange := false
		for j := 0; j < num-1-i; j++ {
			if p.handSort[j].number == 0 {
				p.handSort[j].number = 13
			}
			if p.handSort[j].number > p.handSort[j+1].number {
				p.handSort[j], p.handSort[j+1] = p.handSort[j+1], p.handSort[j]
				isChange = true
			}
		}
		if !isChange {
			break
		}
	}
	// second sort pattern
	for i := 0; i < num-1; i++ {
		isChange := false
		for j := 0; j < num-1-i; j++ {
			if p.handSort[j].number == p.handSort[j+1].number {
				if p.handSort[j].pattern < p.handSort[j+1].pattern {
					p.handSort[j], p.handSort[j+1] = p.handSort[j+1], p.handSort[j]
					isChange = true
				}
			}
		}
		if !isChange {
			break
		}
	}
	for i := 0; i < num; i++ {
		if p.handSort[i].number == 13 {
			p.handSort[i].number = 0
		}
	}

}
func (p *Player) ShowPlayerCard() {
	for i := 0; i < len(p.Hand); i++ {
		fmt.Println(Cardnumber[p.Hand[i].number] + " " + CardPattern[p.Hand[i].pattern])
	}
}
func (p *Player) ShowPlayerSortCard() {
	for i := 0; i < len(p.handSort); i++ {
		fmt.Println(Cardnumber[p.handSort[i].number] + " " + CardPattern[p.handSort[i].pattern])
	}
}

func (p *Player) getJokerCount() int {
	jokerCount := 0
	for i := 0; i < 5; i++ {
		if p.Hand[i].pattern == Joker {
			jokerCount++
		}
	}
	return jokerCount
}

func (p *Player) getPatternCount(pattern int) int {
	patternCount := 0
	for i := 0; i < 5; i++ {
		if p.Hand[i].pattern == pattern {
			patternCount++
		}
	}
	return patternCount
}

func (p *Player) chkRoyal() bool {
	if p.handSort[4].number == 0 {
		if p.handSort[0].number == 9 &&
			p.handSort[1].number == 10 &&
			p.handSort[2].number == 11 &&
			p.handSort[3].number == 12 {
			return true
		}
	}
	return false
}

func (p *Player) chkStraight() bool {
	if p.chkRoyal() {
		return true
	} else if p.handSort[0].number+1 == p.handSort[1].number &&
		p.handSort[1].number+1 == p.handSort[2].number &&
		p.handSort[2].number+1 == p.handSort[3].number &&
		p.handSort[3].number+1 == p.handSort[4].number {
		return true
	}
	return false
}

func (p *Player) chkFlush() bool {
	if p.handSort[0].pattern == p.handSort[1].pattern &&
		p.handSort[0].pattern == p.handSort[2].pattern &&
		p.handSort[0].pattern == p.handSort[3].pattern &&
		p.handSort[0].pattern == p.handSort[4].pattern {
		return true
	}
	return false
}

func (p *Player) chkRoyalFlush() bool {
	if p.chkRoyal() && p.chkFlush() {
		return true
	}
	return false
}
func (p *Player) chkStraightFlush() bool {
	if p.chkStraight() && p.chkFlush() {
		return true
	}
	return false
}

func (p *Player) getSameKind(kind int) int {
	count := 0
	for i := 0; i < len(p.handSort); i++ {
		if p.handSort[i].number == kind {
			count++
		}
	}
	return count
}

func (p *Player) chkFourOfKind() bool {
	max := 0
	for i := 0; i < 5; i++ {
		num := p.getSameKind(p.handSort[i].number)
		if num > max {
			max = num
		}
	}
	if max == 4 {
		return true
	}
	return false
}

func (p *Player) chkThreeOfKind() bool {
	max := 0
	for i := 0; i < 5; i++ {
		num := p.getSameKind(p.handSort[i].number)
		if num > max {
			max = num
		}
	}
	if max == 3 {
		return true
	}
	return false
}

func (p *Player) chkTwoPair() bool {
	if p.handSort[0].number == p.handSort[1].number && p.handSort[2].number == p.handSort[3].number {
		return true
	} else if p.handSort[0].number == p.handSort[1].number && p.handSort[3].number == p.handSort[4].number {
		return true
	} else if p.handSort[1].number == p.handSort[2].number && p.handSort[3].number == p.handSort[4].number {
		return true
	}
	return false
}

func (p *Player) chkOnePair() bool {
	for i := 0; i < 5; i++ {
		num := p.getSameKind(p.handSort[i].number)
		if num == 2 {
			return true
		}
	}
	return false
}

func (p *Player) chkHighPair() bool {
	for i := 0; i < 5; i++ {
		if p.handSort[i].number > 9 || p.handSort[i].number == 0 {
			num := p.getSameKind(p.handSort[i].number)
			if num == 2 {
				return true
			}
		}
	}
	return false
}

func (p *Player) chkFullHouse() bool {
	if p.chkThreeOfKind() && p.chkOnePair() {
		return true
	}
	return false
}

func (p *Player) CheckWin() string {
	if p.chkRoyalFlush() {
		return "Royal Flush"
	} else if p.chkStraightFlush() {
		return "Straight Flush"
	} else if p.chkFourOfKind() {
		return "Four of a Kind"
	} else if p.chkFullHouse() {
		return "Full House"
	} else if p.chkFlush() {
		return "Flush"
	} else if p.chkStraight() {
		return "Straight"
	} else if p.chkThreeOfKind() {
		return "Three of a Kind"
	} else if p.chkTwoPair() {
		return "Two Pair"
	} else if p.chkHighPair() {
		return "Jack or Better"
	}
	return "None"
}
