package main

import (
	"fmt"
	"log"
	router "mongoapi/routers"
	"net/http"
)

func main() {
	fmt.Println("this is main function")
	r := router.Router()
	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("server is starting up at port 4000...")
}
