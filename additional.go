package additionalextpoint

import (
	"fmt"

	"github.com/goldeneggg/geggg-go-extpoint-example/extpoints"
)

func init() {
	extpoints.Register(new(additional), "")
}

type additional struct{}

var _ extpoints.Hoge = new(additional)

func (a *additional) Before(args []string) error {
	fmt.Println("Additional Before")
	return nil
}

func (a *additional) Run(args []string) error {
	fmt.Println("Additional Run")
	return nil
}

func (a *additional) After(args []string) error {
	fmt.Println("Additional After")
	return nil
}

func (a *additional) String() string {
	return "I am additional"
}
