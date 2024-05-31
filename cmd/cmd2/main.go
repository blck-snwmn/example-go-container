package cmd2

import (
	"fmt"

	greet "github.com/blck-snwmn/example-go-container/geet"
)

func main() {
	fmt.Println(greet.Greeter{}.Greet())
}
