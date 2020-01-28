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

	// 定时模块初始化
	pb.cronSrv = new(cronService)
	pb.cronSrv.Init()

	// Web 服务初始化
	pb.webSrv = new(webService)
	pb.webSrv.Init()

}
