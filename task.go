package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"reflect"
	"sync"
)

type taskService struct {
}

type task struct {
	Name     string              `json:"name"`
	Progress map[string][]module `json:"progress"`
}

type module struct {
	Name  string        `json:"module"`
	Param []interface{} `json:"param"`
}

func (ts *taskService) LoadTaskFile(path string) (*task, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, errors.New("file not found")
	}
	t := new(task)
	err = json.Unmarshal(data, t)
	if err != nil {
		return nil, err
	}
	return t, nil
}

func (ts *taskService) ExecTask(t *task) error {
	fmt.Printf("Execute [ %s ]\n", t.Name)
	progresses := reflect.ValueOf(t.Progress).MapKeys()
	wg := &sync.WaitGroup{}
	for _, v := range progresses {
		wg.Add(1)
		go func(value string) {
			ts.execProgress(t.Progress[value])
			wg.Done()
		}(v.String())
	}
	wg.Wait()
	return nil
}

func (ts *taskService)execProgress(modules []module)  {
	for _, mod := range modules {
		_, _ = PB.moduleSrv.Baker.InvokeModuleFunction(mod.Name, mod.Param...)
	}
}
