package main

import (
	"fmt"
	"github.com/pkg/errors"
	"reflect"
)

type Module struct {

}

type Baker struct {
	FunctionList map[string]interface{}
}

func main(){
	baker := new(Baker)
	baker.FunctionList = make(map[string]interface{})
	baker.FunctionList["TEST_MODULE"] = func(a int) {
		fmt.Println(a)
		fmt.Println("hello!!")
	}

	fmt.Println(invoke(baker.FunctionList["TEST_MODULE"], 1))
}

func invoke(f interface{}, params ...interface{}) ([]reflect.Value, error) {
	funcType := reflect.TypeOf(f)
	funcInstance := reflect.ValueOf(f)

	paramNum := funcType.NumIn()
	if len(params) != paramNum{
		return nil, errors.New("params number error")
	}

	realParams := make([]reflect.Value, len(params))
	for index, item := range params {
		if funcType.In(index) != reflect.TypeOf(item){
			return nil, errors.New("params type error")
		}
		realParams[index] = reflect.ValueOf(item)
	}
	result := funcInstance.Call(realParams)
	return result, nil
}
