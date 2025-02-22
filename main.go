package main

import (
	"flag"
	"log"
	"net/http"
	"strconv"
)

func main() {
	var Port int

	flag.IntVar(&Port, "Port", 1008, "Port")

	flag.Parse()
	log.Println("main", "StartParam", Port)
	http.HandleFunc("/", func(Writer http.ResponseWriter, Request *http.Request) {
		Writer.Write([]byte("Hello!World!"))
	})
	log.Println(http.ListenAndServe(("0.0.0.0:" + strconv.Itoa(Port)), nil))
}
