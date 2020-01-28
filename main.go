package main

import (
	"Pie-Baker/modules"
)

type PieBaker struct {
	moduleSrv *modules.ModuleService
	taskSrv   *taskService
	cronSrv   *cronService
	webSrv    *webService
}

var PB *PieBaker

func main() {
	pb := new(PieBaker)
	PB = pb
	pb.Init()
}
