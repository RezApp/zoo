package animal

type Cow struct {
	Animal
}

func (c Cow) Speak() string {
	return "Moo! My name is " + c.Name
}
