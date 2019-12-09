package main

import (
	"Pie-Baker/modules"
)

func (pb *PieBaker) Init(){
	// 模块服务初始化
	pb.moduleSrv = new(modules.ModuleService)
	pb.moduleSrv.Init()

	// 任务初始化
	pb.taskSrv = new(taskService)
	pb.taskSrv.Init()
}
