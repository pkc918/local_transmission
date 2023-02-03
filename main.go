package main

import (
	"os/exec"

	"github.com/gin-gonic/gin"
)

func main() {
	go func(){ // gin 协程
		gin.SetMode(gin.DebugMode)
		router := gin.Default()
		router.GET("/",func(ctx *gin.Context) {
			ctx.Writer.Write([]byte("hello gin"))
		})
		router.Run(":8080")
	}()

	chromePath := "C:\\Program Files (x86)\\Google\\Chrome\\Application\\chrome.exe"
	cmd := exec.Command(chromePath, "--app=http:127.0.0.1:8080/")
	cmd.Start()
	select{}
}