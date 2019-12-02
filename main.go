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
	pb.moduleSrv = new(modules.ModuleService)
	pb.moduleSrv.Init()



	//moduleSrv.Baker.InvokeModuleFunction("GREET")

	task, err := pb.taskSrv.LoadTaskFile("./tasks/1.json")
	if err != nil {
		panic(err)
	}

	pb.taskSrv.ExecTask(task)

}
