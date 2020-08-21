
package main

import "fmt"

func main(){
  //maps in go
  // all value the same type!
  // key -> value
  // key -> value

   colors := map[string]string {
     "red": "#ff0000",
     "green": "#4bf745",
   }
  //colors := make(map[string]string)
  //colors_int := make(map[int]string)

  //colors["white"] ="#ffffff"
  //colors_int[10] ="#ffffff"

  //fmt.Println(colors)
  //fmt.Println(colors_int)

  //delete(colors_int, 10)
  //delete(colors, "white")

  //fmt.Println(colors)
  //fmt.Println(colors_int)
  printMap(colors)
}

func printMap(c map[string]string) {
  for color, hex := range c {
      fmt.Println("Hex code for", color, "is", hex)
  }
}
