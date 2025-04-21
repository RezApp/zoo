package animal

type Animal struct {
	Name string
}

type AnimalBehavior interface {
	Speak() string
}
