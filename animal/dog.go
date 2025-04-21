package animal

type Dog struct {
	Animal
}

func (d Dog) Speak() string {
	return "Woof! My name is " + d.Name
}
