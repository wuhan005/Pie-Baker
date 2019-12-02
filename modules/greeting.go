package modules

import "fmt"

type greeting struct {

}

func (m *Module) Greeting(){
	g := new(greeting)
	Baker.CreateModuleFunction("GREET", g.Greet)
}

func (g *greeting) Greet(){
	fmt.Println("Hello! Welcome to Pie Baker!")
}

