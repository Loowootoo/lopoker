package game

import (
	"fmt"
)

type Player struct {
	Hand     []Card
	HandSort []Card
	Credit   int
	Bet      int
}

func NewPlayer(credit int) *Player {
	hand := make([]Card, 5)
	HandSort := make([]Card, 5)
	return &Player{hand, HandSort, credit, 0}
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
		p.HandSort[i] = p.Hand[i]
		if p.HandSort[i].number == 0 {
			p.HandSort[i].number = 13
		}
	}
	// first sort number
	for i := 0; i < num-1; i++ {
		isChange := false
		for j := 0; j < num-1-i; j++ {
			if p.HandSort[j].number == 0 {
				p.HandSort[j].number = 13
			}
			if p.HandSort[j].number > p.HandSort[j+1].number {
				p.HandSort[j], p.HandSort[j+1] = p.HandSort[j+1], p.HandSort[j]
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
			if p.HandSort[j].number == p.HandSort[j+1].number {
				if p.HandSort[j].pattern < p.HandSort[j+1].pattern {
					p.HandSort[j], p.HandSort[j+1] = p.HandSort[j+1], p.HandSort[j]
					isChange = true
				}
			}
		}
		if !isChange {
			break
		}
	}
	for i := 0; i < num; i++ {
		if p.HandSort[i].number == 13 {
			p.HandSort[i].number = 0
		}
	}

}
func (p *Player) ShowPlayerCard() {
	for i := 0; i < len(p.Hand); i++ {
		fmt.Println(Cardnumber[p.Hand[i].number] + " " + CardPattern[p.Hand[i].pattern])
	}
}
func (p *Player) ShowPlayerSortCard() {
	for i := 0; i < len(p.HandSort); i++ {
		fmt.Println(Cardnumber[p.HandSort[i].number] + " " + CardPattern[p.HandSort[i].pattern])
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
	if p.HandSort[4].number == 0 {
		if p.HandSort[0].number == 9 &&
			p.HandSort[1].number == 10 &&
			p.HandSort[2].number == 11 &&
			p.HandSort[3].number == 12 {
			return true
		}
	}
	return false
}

func (p *Player) chkStraight() bool {
	if p.chkRoyal() {
		return true
	} else if p.HandSort[0].number+1 == p.HandSort[1].number &&
		p.HandSort[1].number+1 == p.HandSort[2].number &&
		p.HandSort[2].number+1 == p.HandSort[3].number &&
		p.HandSort[3].number+1 == p.HandSort[4].number {
		return true
	}
	return false
}

func (p *Player) chkFlush() bool {
	if p.HandSort[0].pattern == p.HandSort[1].pattern &&
		p.HandSort[0].pattern == p.HandSort[2].pattern &&
		p.HandSort[0].pattern == p.HandSort[3].pattern &&
		p.HandSort[0].pattern == p.HandSort[4].pattern {
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
	for i := 0; i < len(p.HandSort); i++ {
		if p.HandSort[i].number == kind {
			count++
		}
	}
	return count
}

func (p *Player) chkFourOfKind() bool {
	max := 0
	for i := 0; i < 5; i++ {
		num := p.getSameKind(p.HandSort[i].number)
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
		num := p.getSameKind(p.HandSort[i].number)
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
	if p.HandSort[0].number == p.HandSort[1].number && p.HandSort[2].number == p.HandSort[3].number {
		return true
	} else if p.HandSort[0].number == p.HandSort[1].number && p.HandSort[3].number == p.HandSort[4].number {
		return true
	} else if p.HandSort[1].number == p.HandSort[2].number && p.HandSort[3].number == p.HandSort[4].number {
		return true
	}
	return false
}

func (p *Player) chkOnePair() bool {
	for i := 0; i < 5; i++ {
		num := p.getSameKind(p.HandSort[i].number)
		if num == 2 {
			return true
		}
	}
	return false
}

func (p *Player) chkHighPair() bool {
	for i := 0; i < 5; i++ {
		if p.HandSort[i].number > 9 || p.HandSort[i].number == 0 {
			num := p.getSameKind(p.HandSort[i].number)
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
