package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"reflect"
)

type taskService struct {
	Tasks []*task
}

type task struct {
	Name     string              `json:"name"`
	Progress map[string][]module `json:"progress"`
}

type module struct {
	Name    string        `json:"module"`
	Param   []interface{} `json:"param"`
	Require bool          `json:"require"`
}

func (ts *taskService) Init(){
	_ = ts.LoadTaskFileList()
	log.Println("Task Service Init")
}

func (ts *taskService) LoadTaskFileList() error{
	ts.Tasks = make([]*task, 0)

	files, err := ioutil.ReadDir("./tasks")
	if err != nil{
		return err
	}
	for _, item := range files{
		fileName := item.Name()
		if !item.IsDir() && fileName[len(fileName) - 5:] == ".json"{
			taskItem, _ := ts.LoadTaskFile("./tasks/" + fileName)
			ts.Tasks = append(ts.Tasks, taskItem)
		}
	}
	return nil
}

func (ts *taskService) GetTaskFileList() []*task{
	return ts.Tasks
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
	errChan := make(chan error)

	fmt.Printf("Execute [ %s ]\n", t.Name)
	progresses := reflect.ValueOf(t.Progress).MapKeys()
	//wg := &sync.WaitGroup{}
	// task 任务并发执行
	for _, v := range progresses {
		if len(errChan) >= 1 {
			break
		}

		go func(value string) {
			err := ts.execProgress(t.Progress[value])
			// 判断当前 Module 是否出错
			if err != nil {
				errChan <- err
			}
		}(v.String())
	}

	if len(errChan) >= 1 {
		return <-errChan
	}
	return nil
}

// 执行单个模块
func (ts *taskService) execProgress(modules []module) error {
	for _, mod := range modules {
		m := mod
		_, err := PB.moduleSrv.Baker.InvokeModuleFunction(m.Name, m.Param...)

		// 检测必要模块是否运行成功
		if m.Require && err != nil {
			return err
		}
	}
	return nil
}
