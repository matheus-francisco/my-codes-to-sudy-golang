# Frist Project in Golang... (Cards)

### Cards

* newDeck: Create a list of playing cards. Essentially an array of strings
* print: Log out the contents of a deck
* shuffle: Shuffles all the cards in a deck
* deal: create a 'hand' of cards.
* saveToFile: Save a list of cards to a file on the local machine
* newDeckFromFile: Load a list of cards from the local machine


## Go basic types
* bool
* string
* int
* flaot

## Go Data Sturct
* Array: Fixed length list of things
* Slice: An array that can grow or shrink ( every element in a slice must be same type )
* type newtype  []string : new type slice strings

* this will be print (byte ascii to "hello")
```
package main

import "fmt"

func main() {
  fmt.Println([]byte("Hello"))
}

```


## Packages standard library
https://golang.org/pkg/

