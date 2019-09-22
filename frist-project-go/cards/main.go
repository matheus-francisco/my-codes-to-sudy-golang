package main

//import "fmt"

func main() {
  cards := newDeck()
  cards.saveToFile("my_cards")
  my_new_deck := newDeckFromFile("my_cards")
  my_new_deck.shuffle()
  my_new_deck.print()
}


