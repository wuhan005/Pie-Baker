package modules

import "fmt"

type greeting struct {

}

func (m *Module) Greeting(){
	g := new(greeting)
	Baker.CreateModuleFunction("GREET", g.Greet)
}

func (g *greeting) Greet(name string){
	fmt.Printf("Hello %s! Welcome to Pie Baker!\n", name)
}

