package main

import "fmt"

type contactInfo struct {
  email string
  zipCode int
}

type person struct {
  firstName string
  lastName string
  contactInfo
}



func main(){
    mat := person{
      firstName: "Mat",
      lastName: "Francisco",
      contactInfo: contactInfo{
          email: "fake@gmail.com",
          zipCode: 44444,
        },
    }

    maty := person{
      firstName: "Mat",
      lastName: "Francisco",
      contactInfo: contactInfo{
          email: "fake@gmail.com",
          zipCode: 44444,
        },
    }

    matPointer := &mat
    matPointer.updateName("Matheus")

    mat.print()

    maty.updateName("Xico")
    maty.print()
}

func (pointerToPerson *person) updateName(newFirstName string) {
  (*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
    fmt.Printf("%+v", p)
}
