# speed-golang
An application to play the card game speed, based on indy-golang/simple-chat (description below).
This application is meant as a learning excercise for concurrency in golang

## Rules
At the beginning of the game, the deck is split into 6 piles:
  - 5 card hands for each player (2 players total)
  - Two played piles, 1 card each
  - The rest is split into two decks, one for each player
At any time, a player can play a card as long as the card's value is one above or below something in the played pile.
At any time, a player can draw cards, but can only have a maximum of 5 cards at a time in their hand.
The player who runs out of cards in hand and in deck wins
If there is a point where both players don't have legal plays, the top card of both decks is placed on top of the played piles.

# Original README ----------

# simple-chat
A simple IRC-like chat created for the Apr 7, 2015 Indy Golang Meeting

## General

This is a very simple chat server written in Go, for demonstrating concurrency and the use of channels.

This package is self contained and doesn't require any dependencies. 

The server is basically a webserver which serves the files necessary to run the chat clients in a browser, plus it handles the chat traffic via websockets.

## Usage

- Download / clone the project
- run the server: `go run server.go`
- connect to the server from any computer in your local network: `http://_server_ip_or_name_:8000`

That's it! 
