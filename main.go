package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

const (
	goldBlackjack       = 22
	blackjack           = 21
	playerThreshold int = 17
)

type player struct {
	name  string
	cards []card
}

type card struct {
	suit  string
	value string
}

func main() {
	var deck []card
	suits := []string{"D", "C", "H", "S"}
	values := []string{"2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K", "A"}
	p := player{name: "player", cards: []card{}}
	d := player{name: "dealer", cards: []card{}}

	for _, suit := range suits {
		for _, value := range values {
			deck = append(deck, card{suit: suit, value: value})
		}
	}

	deck = shuffleDeck(deck)
	play(deck, &p, &d)
}

func play(deck []card, p *player, d *player) {
	dealCards(deck, p, d)

	if d.getHandValue() == goldBlackjack && p.getHandValue() == goldBlackjack {
		printWinner(d, p)
		return
	}

	if hasBlackjack(d) && hasBlackjack(p) {
		printWinner(d, p)
		return
	}

	if hasBlackjack(d) {
		printWinner(d, p)
	}

	if hasBlackjack(p) {
		printWinner(p, d)
	}

	for p.getHandValue() < playerThreshold {
		drawCard(deck, p)
		if p.getHandValue() > blackjack {
			printWinner(d, p)
			return
		}
	}

	for d.getHandValue() <= p.getHandValue() {
		drawCard(deck, d)
		if d.getHandValue() > blackjack {
			printWinner(p, d)
			return
		}
	}

	if p.getHandValue() > d.getHandValue() {
		printWinner(p, d)
	} else {
		printWinner(d, p)
	}
}

func shuffleDeck(deck []card) []card {
	rand.Seed(time.Now().UnixMilli())
	time.Sleep(time.Millisecond * 50)
	for i := 0; i < len(deck)-1; i++ {
		r := rand.Intn(len(deck))
		deck[i], deck[r] = deck[r], deck[i]
	}
	return deck
}

func dealCards(deck []card, p *player, d *player) {
	deck = drawCard(deck, p)
	deck = drawCard(deck, d)
	deck = drawCard(deck, p)
	deck = drawCard(deck, d)
}

func drawCard(deck []card, p *player) []card {
	rand.Seed(time.Now().UnixMilli())
	time.Sleep(time.Millisecond * 50)
	index := rand.Intn(len(deck))
	p.cards = append(p.cards, deck[index])
	return append(deck[:index], deck[index+1:]...)
}

func hasBlackjack(p *player) bool {
	return p.getHandValue() == 21
}

func printWinner(winner *player, loser *player) {
	fmt.Printf("winner: %v\n", winner.name)
	fmt.Printf("%v: (%v) %v\n", winner.name, winner.getHandValue(), winner.cards)
	fmt.Printf("%v: (%v) %v\n", loser.name, loser.getHandValue(), loser.cards)
}

func (p *player) getHandValue() int {
	result := 0
	for _, card := range p.cards {
		result += card.getCardNominal()
	}
	return result
}

func (c *card) getCardNominal() int {
	result := 0
	switch c.value {
	case "10", "J", "Q", "K", "A":
		result = 10
	default:
		num, _ := strconv.Atoi(c.value)
		result = num
	}
	return result
}
