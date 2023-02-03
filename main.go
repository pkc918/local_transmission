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

	// listen tcp of net/http
	ln, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Fatal(err)
	}
	defer ln.Close()
	
	// 使用
	staticFiles, _ := fs.Sub(FS, "frontend/dist")
	go http.Serve(ln, http.FileServer(http.FS(staticFiles)))
	fmt.Println(fmt.Sprintf("http://%s/", ln.Addr()))
	ui.Load(fmt.Sprintf("http://%s/", ln.Addr()))

	// 监听 ctrl c
	chSignal := make(chan os.Signal, 1)
	signal.Notify(chSignal, os.Interrupt)

	// Wait until UI window is closed or Terminal input Ctrl+C
	select{
	case <-chSignal:
	case  <-ui.Done():
	}
}