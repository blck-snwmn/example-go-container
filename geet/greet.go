package greet

type Greeter struct{}

func (g Greeter) Greet() string {
	return "Hello!"
}
