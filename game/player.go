package game

import (
	"fmt"
)

type Player struct {
	Hand     [5]Card
	HandSort [5]Card
	Held     [5]bool
	BackCard [5]bool
	Credit   int
	Bet      int
}

func NewPlayer(credit int) *Player {
	player := new(Player)
	player.Credit = credit
	player.Bet = 0
	player.ResetBackCard()
	player.ResetHeld()
	return player
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
		if p.HandSort[i].Number == 0 {
			p.HandSort[i].Number = 13
		}
	}
	// first sort number
	for i := 0; i < num-1; i++ {
		isChange := false
		for j := 0; j < num-1-i; j++ {
			if p.HandSort[j].Number == 0 {
				p.HandSort[j].Number = 13
			}
			if p.HandSort[j].Number > p.HandSort[j+1].Number {
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
			if p.HandSort[j].Number == p.HandSort[j+1].Number {
				if p.HandSort[j].Pattern < p.HandSort[j+1].Pattern {
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
		if p.HandSort[i].Number == 13 {
			p.HandSort[i].Number = 0
		}
	}

}
func (p *Player) ShowPlayerCard() {
	for i := 0; i < len(p.Hand); i++ {
		fmt.Println(Cardnumber[p.Hand[i].Number] + " " + CardPattern[p.Hand[i].Pattern])
	}
}
func (p *Player) ShowPlayerSortCard() {
	for i := 0; i < len(p.HandSort); i++ {
		fmt.Println(Cardnumber[p.HandSort[i].Number] + " " + CardPattern[p.HandSort[i].Pattern])
	}
}

func (p *Player) getJokerCount() int {
	jokerCount := 0
	for i := 0; i < 5; i++ {
		if p.Hand[i].Pattern == Joker {
			jokerCount++
		}
	}
	return jokerCount
}

func (p *Player) getPatternCount(pattern int) int {
	patternCount := 0
	for i := 0; i < 5; i++ {
		if p.Hand[i].Pattern == pattern {
			patternCount++
		}
	}
	return patternCount
}

func (p *Player) ResetHeld() {
	for i := 0; i < 5; i++ {
		p.Held[i] = false
	}
}
func (p *Player) ResetBackCard() {
	for i := 0; i < 5; i++ {
		p.BackCard[i] = true
	}
}
func (p *Player) setHeld(sortIndex int) {
	for i := 0; i < 5; i++ {
		if p.HandSort[sortIndex].Number == p.Hand[i].Number && p.HandSort[sortIndex].Pattern == p.Hand[i].Pattern {
			p.Held[i] = true
		}
	}
}

func (p *Player) SetBackCard(i int, status bool) {
	p.BackCard[i] = status
}

func (p *Player) chkRoyal() bool {
	if p.HandSort[4].Number == 0 {
		if p.HandSort[0].Number == 9 &&
			p.HandSort[1].Number == 10 &&
			p.HandSort[2].Number == 11 &&
			p.HandSort[3].Number == 12 {
			p.Held[0] = true
			p.Held[1] = true
			p.Held[2] = true
			p.Held[3] = true
			p.Held[4] = true
			return true
		}
	}
	return false
}

func (p *Player) chkStraight() bool {
	if p.chkRoyal() {
		return true
	} else if p.HandSort[0].Number+1 == p.HandSort[1].Number &&
		p.HandSort[1].Number+1 == p.HandSort[2].Number &&
		p.HandSort[2].Number+1 == p.HandSort[3].Number &&
		p.HandSort[3].Number+1 == p.HandSort[4].Number {
		p.Held[0] = true
		p.Held[1] = true
		p.Held[2] = true
		p.Held[3] = true
		p.Held[4] = true
		return true
	}
	return false
}

func (p *Player) chkFlush() bool {
	if p.HandSort[0].Pattern == p.HandSort[1].Pattern &&
		p.HandSort[0].Pattern == p.HandSort[2].Pattern &&
		p.HandSort[0].Pattern == p.HandSort[3].Pattern &&
		p.HandSort[0].Pattern == p.HandSort[4].Pattern {
		p.Held[0] = true
		p.Held[1] = true
		p.Held[2] = true
		p.Held[3] = true
		p.Held[4] = true
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
		if p.HandSort[i].Number == kind {
			count++
		}
	}
	return count
}

func (p *Player) chkFourOfKind() bool {
	max := 0
	card := 0
	for i := 0; i < 5; i++ {
		num := p.getSameKind(p.HandSort[i].Number)
		if num > max {
			max = num
			card = p.HandSort[i].Number
		}
	}
	if max == 4 {
		for i := 0; i < 5; i++ {
			if p.Hand[i].Number == card {
				p.Held[i] = true
			}
		}
		return true
	}
	return false
}

func (p *Player) chkThreeOfKind() bool {
	max := 0
	card := 0
	for i := 0; i < 5; i++ {
		num := p.getSameKind(p.HandSort[i].Number)
		if num > max {
			max = num
			card = p.HandSort[i].Number
		}
	}
	if max == 3 {
		for i := 0; i < 5; i++ {
			if p.Hand[i].Number == card {
				p.Held[i] = true
			}
		}
		return true
	}
	return false
}
func (p *Player) chkTwoPair() bool {
	if p.HandSort[0].Number == p.HandSort[1].Number && p.HandSort[2].Number == p.HandSort[3].Number {
		p.setHeld(0)
		p.setHeld(1)
		p.setHeld(2)
		p.setHeld(3)
		return true
	} else if p.HandSort[0].Number == p.HandSort[1].Number && p.HandSort[3].Number == p.HandSort[4].Number {
		p.setHeld(0)
		p.setHeld(1)
		p.setHeld(3)
		p.setHeld(4)
		return true
	} else if p.HandSort[1].Number == p.HandSort[2].Number && p.HandSort[3].Number == p.HandSort[4].Number {
		p.setHeld(1)
		p.setHeld(2)
		p.setHeld(3)
		p.setHeld(4)
		return true
	}
	return false
}

func (p *Player) chkOnePair() bool {
	card := 0
	for i := 0; i < 5; i++ {
		num := p.getSameKind(p.HandSort[i].Number)
		if num == 2 {
			card = p.HandSort[i].Number
			for j := 0; j < 5; j++ {
				if p.Hand[j].Number == card {
					p.Held[j] = true
				}
			}
			return true
		}
	}
	return false
}

func (p *Player) chkHighPair() bool {
	for i := 0; i < 5; i++ {
		if p.HandSort[i].Number > 9 || p.HandSort[i].Number == 0 {
			num := p.getSameKind(p.HandSort[i].Number)
			if num == 2 {
				card := p.HandSort[i].Number
				for j := 0; j < 5; j++ {
					if p.Hand[j].Number == card {
						p.Held[j] = true
					}
				}
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

var WinKind = [10]string{
	"ROYAL FLUSH",
	"STRAIGHT FLUSH",
	"FOUR OF A KIND",
	"FULL HOUSE",
	"FLUSH",
	"STRAIGHT",
	"THREE OF A KIND",
	"TWO PAIR",
	"JACK OR BETTER",
	"NONE",
}

func (p *Player) CheckWin() string {
	if p.chkRoyalFlush() {
		return WinKind[0]
	} else if p.chkStraightFlush() {
		return WinKind[1]
	} else if p.chkFourOfKind() {
		return WinKind[2]
	} else if p.chkFullHouse() {
		return WinKind[3]
	} else if p.chkFlush() {
		return WinKind[4]
	} else if p.chkStraight() {
		return WinKind[5]
	} else if p.chkThreeOfKind() {
		return WinKind[6]
	} else if p.chkTwoPair() {
		return WinKind[7]
	} else if p.chkHighPair() {
		return WinKind[8]
	}
	return WinKind[9]
}
