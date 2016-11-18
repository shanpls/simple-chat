package card

import (
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
	played *PlayedPiles 
	sync.Mutex // Protects player's hand and deck
	deck Deck 
	hand Hand 
}

func (d Deck) DrawCard() Card {
	var newCard Card = Card{"", 0}

	if len(d.cards) != 0 {
		newCard = d.cards[0]
		d.cards = d.cards[1: ]
	}

	return newCard;
}

func (d Deck) isEmpty() bool{
	return len(d.cards) == 0
}

func (p PlayedPiles) AddToPile(c Card) bool{
	p.Lock()
	defer p.Lock()

	if (c.rank == p.pileA[len(p.pileA)-1].rank + 1 || c.rank == p.pileA[len(p.pileA)-1].rank - 1){
		p.pileA = append(p.pileA, c)
		return true
	} else if (c.rank == p.pileB[len(p.pileB)-1].rank + 1 || c.rank == p.pileB[len(p.pileB)-1].rank - 1){
		p.pileB = append(p.pileB, c)
		return true
	}

	return false
}

func (p PlayedPiles) CurrentPile() (Card, Card){
	return p.pileA[len(p.pileA)-1], p.pileB[len(p.pileB)-1]
}

func (h Hand) CardAtIndex(index int) Card{
	return h.cards[index]
}

func (h Hand) playCard(index int) {
	if(index < len(h.cards)){
		h.cards = append(h.cards[ :index], h.cards[index+1: ]...)
	}
}

func (h Hand) addToHand(c Card) {
	h.cards = append(h.cards, c)
}

func (c Controller) DrawCard() {
	c.Lock()
	defer c.Unlock()

	if(len(c.hand.cards) < 5 && !c.deck.isEmpty()){
		newCard := c.deck.DrawCard();
		c.hand.addToHand(newCard);
	}
}

func (c Controller) PlayCard(index int) {
	c.Lock()
	defer c.Unlock()

	cardToPlay := c.hand.CardAtIndex(index);
	success := c.played.AddToPile(cardToPlay);
	if(success){
		c.hand.playCard(index)
	}
}

