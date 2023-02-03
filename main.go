package main

import (
	"embed"
	"fmt"
	"io/fs"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"

	"github.com/zserge/lorca"
)

//go:embed frontend/dist/*
var FS embed.FS

func main() {
	// Create instance lorca
	ui, err := lorca.New("", "", 480, 320)
	if err != nil {
		log.Fatal(err)
	}
	defer ui.Close()

	ln, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Fatal(err)
	}
	defer ln.Close()
	
	// 使用
	staticFiles, _ := fs.Sub(FS, "frontend/dist")
	go http.Serve(ln, http.FileServer(http.FS(staticFiles)))
	ui.Load(fmt.Sprintf("http://%s/", ln.Addr()))


	// 监听 ctrl c
	chSignal := make(chan os.Signal, 1)
	signal.Notify(chSignal, os.Interrupt)


	// Wait until UI window is closed or Terminal input Ctrl+C
	select{
	case <-chSignal:
	case  <-ui.Done():
	}
		
	// go func(){ // gin 协程
	// 	gin.SetMode(gin.DebugMode)
	// 	router := gin.Default()
	// 	router.GET("/",func(ctx *gin.Context) {
	// 		ctx.Writer.Write([]byte("hellogin"))
	// 	})
	// 	router.Run(":8080")
	// }()
	// // 监听 ctrl c
	// chSignal := make(chan os.Signal, 1)
	// signal.Notify(chSignal, os.Interrupt)

	// chromePath := "C:\\Program Files (x86)\\Google\\Chrome\\Application\\chrome.exe"
	// cmd := exec.Command(chromePath, "--app=http:127.0.0.1:8080/")
	// cmd.Start()
	// fmt.Println(cmd.Process.Pid)
	// // 等待ctrl c信号
	// <-chSignal // 表示从chSignal里读值，该过程是阻塞的，没有读到值就一直等
	// err := cmd.Process.Kill().Error()
	// fmt.Println("err: ", err)
	// fmt.Println("Process killed with PID: ", cmd.Process.Pid)
}