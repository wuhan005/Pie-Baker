package modules

import (
	"errors"
	"fmt"
	"log"
	"reflect"
)

// 全局的 Baker
var Baker = new(ModuleBaker)

type Module struct{}

// 模块管理器
type ModuleBaker struct {
	functionList map[string]interface{}
}

type ModuleService struct {
	Baker *ModuleBaker // 模块管理器
	Mod   *Module      // 模块
}

// 初始化所有注册的模块
func (m *ModuleService) Init() {
	m.Baker = Baker
	m.Baker.functionList = make(map[string]interface{})

	m.Mod = new(Module)
	mod := reflect.ValueOf(m.Mod)
	for i := 0; i < mod.NumMethod(); i++ {
		mod.Method(i).Call([]reflect.Value{}) // 执行所有 Module 的构造函数
	}
	log.Println("Module Service Init")
}

// 模块内注册函数
func (m *ModuleBaker) CreateModuleFunction(name string, function interface{}) {
	if m.functionList[name] != nil {
		log.Printf("The function name [%s] existed!\n", name)
	}
	m.functionList[name] = function
}

// 执行模块函数
func (m *ModuleBaker) InvokeModuleFunction(funcName string, params []reflect.Value) ([]reflect.Value, error) {
	f, ok := m.functionList[funcName]
	if !ok {
		return nil, errors.New("function not found")
	}
	return m.invokeFunction(f, params)
}

func (m *ModuleBaker) invokeFunction(f interface{}, params []reflect.Value) ([]reflect.Value, error) {
	funcType := reflect.TypeOf(f)
	funcInstance := reflect.ValueOf(f)
	funcParamNum := funcType.NumIn()
	inParamNum := len(params)
	if inParamNum != funcParamNum {
		if funcParamNum == 0 {
			// 若函数无入参，则舍去所有参数继续执行
			params = []reflect.Value{}
		}else{
			return nil, errors.New(fmt.Sprintf("Params number error. Expect %d, got %d.", funcParamNum, inParamNum))
		}
	}

	realParams := make([]reflect.Value, funcParamNum)
	for index, param := range params {
		funcParamType := funcType.In(index)
		inParamType := param.Type()
		if funcParamType != inParamType {
			return nil, errors.New(fmt.Sprintf("Params type error. Expect %s, got %s.", funcParamType, inParamType))
		}
		realParams[index] = param
	}
	result := funcInstance.Call(realParams)
	return result, nil
}
