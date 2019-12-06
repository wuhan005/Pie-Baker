package main

import "Pie-Baker/modules"

func (pb *PieBaker) init(){
	// 模块服务初始化
	pb.moduleSrv = new(modules.ModuleService)
	pb.moduleSrv.Init()
}
