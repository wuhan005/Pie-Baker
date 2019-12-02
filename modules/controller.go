package modules

import (
	"errors"
	"reflect"
)

var Baker = new(ModuleBaker)

type ModuleService struct {
	Baker *ModuleBaker
	Mod   *Module
}

func (m *ModuleService) Init() {
	m.Baker = Baker
	m.Baker.functionList = make(map[string]interface{})
	m.Mod = new(Module)

	mod := reflect.ValueOf(m.Mod)
	for i := 0; i < mod.NumMethod(); i++ {
		mod.Method(i).Call([]reflect.Value{})
	}
}

type Module struct{}

type ModuleBaker struct {
	functionList map[string]interface{}
}

func (m *ModuleBaker) CreateModuleFunction(name string, function interface{}) {
	m.functionList[name] = function
}

func (m *ModuleBaker) InvokeModuleFunction(funcName string, params ...interface{}) ([]reflect.Value, error) {
	f, ok := m.functionList[funcName]
	if !ok {
		return nil, errors.New("function not found")
	}
	return m.invokeFunction(f, params)
}

func (m *ModuleBaker) invokeFunction(f interface{}, params []interface{}) ([]reflect.Value, error) {
	funcType := reflect.TypeOf(f)
	funcInstance := reflect.ValueOf(f)
	funcParamNum := funcType.NumIn()
	inParamNum := len(params)
	if inParamNum != funcParamNum {
		return nil, errors.New("params number error")
	}

	realParams := make([]reflect.Value, inParamNum)
	for index, param := range params {
		if funcType.In(index) != reflect.TypeOf(param) {
			return nil, errors.New("params type error")
		}
		realParams[index] = reflect.ValueOf(param)
	}
	result := funcInstance.Call(realParams)
	return result, nil
}
