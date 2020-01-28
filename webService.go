package main

import "github.com/gin-gonic/gin"

type webService struct {
	router *gin.Engine
}

func (ws *webService) Init() {
	gin.SetMode(gin.ReleaseMode)
	ws.router = gin.Default()

	ws.router.Run(":9090")
}
