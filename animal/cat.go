package animal

type Cat struct {
	Animal
}

func (c Cat) Speak() string {
	return "Meow! My name is " + c.Name
}
