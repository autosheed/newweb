package main

import (
	"io"
	"log"
	"net/http"
	"os/exec"
)

func reLaunch() {
	cmd := exec.Command("sh","./deploy.sh")
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	err = cmd.Wait()
}

func run() {
	http.HandleFunc("/",blockchainGetHandle)
	http.ListenAndServe(":5000",nil)
}

func blockchainGetHandle(w http.ResponseWriter,r *http.Request) {
	io.WriteString(w,"<h1>hello , i am webserver deploy</h1>")
	reLaunch()
}

func main() {
	run()
}