package modules

import "log"

type variable struct {
	data map[string]interface{}
}

func (m *Module) Variable() {
	v := new(variable)
	v.data = make(map[string]interface{})
	Baker.CreateModuleFunction("SET_VAR", v.SetVar)
	Baker.CreateModuleFunction("GET_VAR", v.GetVar)
}

func (v *variable) SetVar(name string, value interface{}) interface{} {
	v.data[name] = value
	return value
}

func (v *variable) GetVar(name string) interface{} {
	value, ok := v.data[name]
	if ok{
		return value
	}else{
		log.Printf("Try to get undefined variable: %s\n", name)
		return nil
	}
}
