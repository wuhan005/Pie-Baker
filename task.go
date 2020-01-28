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

// 任务文件数据格式
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
	err := ts.LoadTaskFileList()
	if err != nil{
		panic(err)
	}else{
		log.Println("Task Service Init")
	}
}

// 加载所有任务文件
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

func (ts *taskService) GetTaskList() []*task{
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

	// 并发执行所有 Progress
	for _, v := range progresses {
		if len(errChan) >= 1 {
			break
		}

		go func(name string) {
			err := ts.execProgress(t.Progress[name])
			// 必要模块且出错
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
	var tmpData []reflect.Value		// 模块函数返回值存储

	for _, mod := range modules {
		var backData []reflect.Value
		var err error

		m := mod
		if m.Param == nil{
			// 使用上一模块的返回值
			backData, err = PB.moduleSrv.Baker.InvokeModuleFunction(m.Name, tmpData)
		}else{
			// 将 []interface{} 转换为 []reflect.Value
			var val []reflect.Value
			for _, param := range m.Param{
				val = append(val, reflect.ValueOf(param))
			}
			backData, err = PB.moduleSrv.Baker.InvokeModuleFunction(m.Name, val)
		}
		// 存储返回值
		tmpData = backData

		// 检测模块是否运行成功
		if err != nil{
			log.Printf("Execute function [ %s ] fail: %s\n", m.Name, err)
			if m.Require{
				// 若为必要模块，则返回错误从而不继续往下执行
				return err
			}
		}
	}
	return nil
}
