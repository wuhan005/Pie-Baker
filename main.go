package main

import (
	"Pie-Baker/modules"
)

func main(){
	moduleSrv := new(modules.ModuleService)
	moduleSrv.Init()

	moduleSrv.Baker.InvokeModuleFunction("GREET")
}
