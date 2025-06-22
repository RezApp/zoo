package main

import (
	"fmt"

	"github.com/RezApp/zoo/animal"
)

func MakeItSpeak(a animal.AnimalBehavior) {
	fmt.Println(a.Speak())
}

func main() {
	dog := animal.Dog{Animal: animal.Animal{Name: "Buddy"}}
	cat := animal.Cat{Animal: animal.Animal{Name: "Tom"}}
	cow := animal.Cow{Animal: animal.Animal{Name: "John"}}
	MakeItSpeak(dog)
	MakeItSpeak(cat)
	MakeItSpeak(cow)
}
