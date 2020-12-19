package data

type Greeter struct {
}

func (g Greeter) GetGreeter(name string) (age int, size string) {

	age = 18
	size = "B"

	return
}

func NewGreeter() Greeter {
	return Greeter{}
}
