package logger

import (
	"io"
	"log"
	"os"
)

func Log2file(msg interface{}) {
	f, err := os.OpenFile("log.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	ws := io.MultiWriter(f, os.Stdout)
	ls := log.New(ws, "[WSChat]", log.Ldate | log.Ltime)
	ls.Println(msg)
}
