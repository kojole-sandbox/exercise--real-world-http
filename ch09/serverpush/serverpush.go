package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

var image []byte

func init() {
	var err error
	image, err = ioutil.ReadFile("./image.png")
	if err != nil {
		panic(err)
	}
}

func handlerHTML(w http.ResponseWriter, r *http.Request) {
	pusher, ok := w.(http.Pusher)
	if ok {
		pusher.Push("/image", nil)
	}
	w.Header().Add("Content-Type", "text/html")
	fmt.Fprintf(w, `<html><body><img src="/image"></body></html>`)
}

func handlerImage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/png")
	w.Write(image)
}

func main() {
	http.HandleFunc("/", handlerHTML)
	http.HandleFunc("/image", handlerImage)
	fmt.Println("start http listening :18433")
	err := http.ListenAndServeTLS(":18443", "../../ch06/tlsserver/server.crt", "../../ch06/tlsserver/server.key", nil)
	fmt.Println(err)
}
