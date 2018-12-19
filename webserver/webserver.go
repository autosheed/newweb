package main

import (
	"io"
	"net/http"
)

func run() {
	http.HandleFunc("/",blockchainGetHandle)
	http.ListenAndServe(":8000",nil)
}

func blockchainGetHandle(w http.ResponseWriter,r *http.Request) {
	io.WriteString(w,"<h1>this is a auto deploy webserver</h1>")
}

func main() {
	run()
}