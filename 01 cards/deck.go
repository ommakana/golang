package main

import (
	"fmt"
	"os"
	"strings"
)

type deck []string

func newDeck() deck {
	cards := deck{}
	suites := []string{"Hearts", "Spades", "Diamonds", "Clubs"}
	suitesNumbers := []string{"Ace", "Two", "Three", "Four"}
	for _, suite := range suites {
		for _, number := range suitesNumbers {
			cards = append(cards, number+" of "+suite)
		}
	}
	return cards
}

func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}

func (d deck) deal(handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) saveToFile(filename string) {
	bs := []byte(strings.Join(d, ","))
	err := os.WriteFile(filename, bs, 0666)
	if err != nil {
		fmt.Println("Error saving file")
		fmt.Println(err)
	}
}

func readFromFile(filename string) deck {
	bs, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading from file")
		fmt.Println(err)
		return newDeck()
	} else {
		cardString := string(bs)
		return deck(strings.Split(cardString, ","))
	}
}
