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
	pb.Init()

}
