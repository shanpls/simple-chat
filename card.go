package card

import (
	"fmt"
	"sync"
)

type Card struct{
	suit string
	rank int
}

type Deck struct{
	cards []Card
}

type Hand struct{
	cards []Card
}

type PlayedPiles struct{
	sync.Mutex //Protects the played piles
	pileA []Card
	pileB []Card
}

type Controller struct{
	*PlayedPiles played
	sync.Mutex // Protects player's hand and deck
	Deck deck
	Hand hand
}

func (d Deck) DrawCard() Card {
	var newCard Card = nil

	if len(d.Cards) != 0 {
		newCard = d.cards[0]
		d.cards = d.cards[1: ]
	}

	return newCard;
}

func (d Deck) isEmpty() bool{
	return len(d.cards) == 0
}

func (p Pile) AddToPile(c Card) bool{
	p.Lock()
	defer p.Lock()

	if (c.rank == p.pileA[len(p.pileA)-1].rank + 1 || c.rank == p.pileA[len(p.pileA)-1].rank - 1){
		p.pileA = append(p.pileA, c)
		return true
	}
	else if(c.rank == p.pileB[len(p.pileB)-1].rank + 1 || c.rank == p.pileB[len(p.pileB)-1].rank - 1){
		p.pileB = append(p.pileB, c)
		return true
	}

	return false
}

func (p Pile) CurrentPile() Card, Card{
	return p.pileA[len(p.pileA)-1], p.pileB[len(p.pileB)-1]
}

func (h Hand) CardAtIndex(index int) Card{
	return h.cards[index]
}

func (h Hand) playCard(index int) void{
	if(index < len(h.cards)){
		h.cards = append(h.cards[ :index], h.cards[index+1: ])
	}
}

func (h Hand) addToHand(c Card) void{
	h.cards = append(p.cards, c)
}

func (c Controller) DrawCard() void{
	c.Lock()
	defer c.Unlock()

	if(len(c.hand < 5) && !c.deck.isEmpty()){
		newCard := c.deck.DrawCard();
		c.hand.addToHand(newCard);
	}
}

func (c Controller) PlayCard(index int) void{
	c.Lock()
	defer c.Unlock()

	cardToPlay := c.hand.CardAtIndex(index);
	bool success = c.played.AddToPile(cardToPlay);
	if(success){
		c.hand.playCard(index)
	}
}

