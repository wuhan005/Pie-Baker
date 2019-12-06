package main

import (
	"Pie-Baker/modules"
)

type PieBaker struct {
	moduleSrv *modules.ModuleService
	taskSrv   *taskService
}

var PB *PieBaker
func main() {
	pb := new(PieBaker)
	PB = pb
	pb.init()

	task, err := pb.taskSrv.LoadTaskFile("./tasks/1.json")
	if err != nil {
		panic(err)
	}

	pb.taskSrv.ExecTask(task)
}
