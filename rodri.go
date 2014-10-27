package main

import (
	"log"
	"net/http"
	"os"
)

func main(){
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)
	log.Println("Listening...")
	err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if err != nil {
		panic(err)
	}
}
