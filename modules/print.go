package modules

import "log"

type print struct {
}

func (m *Module) Print() {
	p := new(print)
	Baker.CreateModuleFunction("PRINT_LOG", p.Log)
}

func (p *print) Log(content string) {
	log.Println(content)
}
