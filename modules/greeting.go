package modules

import "fmt"

type greeting struct {

}

func (m *Module) Greeting(){
	g := new(greeting)
	Baker.CreateModuleFunction("GREET", g.Greet)
	Baker.CreateModuleFunction("HELLO_WORLD", g.HelloWorld)
	Baker.CreateModuleFunction("GET_NAME", g.GetName)
}

func (g *greeting) Greet(name string){
	fmt.Printf("Hello %s! Welcome to Pie Baker!\n", name)
}

func (g *greeting) HelloWorld(){
	fmt.Println("Hello world!")
}

func (g *greeting) GetName() string{
	return "E99E99E99!"
}