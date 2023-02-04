package main

import (
	"embed"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/zserge/lorca"
)

//go:embed frontend/dist/*
var FS embed.FS
const serverPort int = 8080

func main() {
	// Create instance lorca
	ui, err := lorca.New("", "", 480, 320)
	if err != nil {
		log.Fatal(err)
	}
	defer ui.Close()

	// go 协程
	go func ()  {
		mux := http.NewServeMux()
		staticFiles, _ := fs.Sub(FS, "frontend/dist")
		mux.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.FS(staticFiles))))
		
		mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			fmt.Printf("server: %s /\n", r.Method)
			fmt.Fprint(w, "Hello Go")
		})
		server := &http.Server{
			Addr:    fmt.Sprintf(":%d", serverPort),
			Handler: mux,
		}
		go server.ListenAndServe()
		ui.Load(fmt.Sprintf("http://127.0.0.1:%d/static/", serverPort))
	}()

	// 监听 ctrl c
	chSignal := make(chan os.Signal, 1)
	signal.Notify(chSignal, os.Interrupt)

	// Wait until UI window is closed or Terminal input Ctrl+C
	select{
	case <-chSignal:
	case  <-ui.Done():
	}
}