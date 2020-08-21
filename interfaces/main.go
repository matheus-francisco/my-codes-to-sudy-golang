package main

import "fmt"

type bot interface {
  getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
  eb := englishBot{}
  sp := spanishBot{}

  printGreeting(eb)
  printGreeting(sp)
}

func printGreeting(b bot) {
  fmt.Println(b.getGreeting())
}

//func printGreeting(eb englishBot) {
//  fmt.Println(eb.getGreeting())
//}
//
//func printGreeting(sp spanishBot) {
//  fmt.Println(eb.getGreeting())
//}

func (englishBot) getGreeting() string {
  // Very custom logic for genereting an english
  return "Hi There!"
}

func (spanishBot) getGreeting() string {
  // Very custom logic for genereting hola
  return "Hola"
}
