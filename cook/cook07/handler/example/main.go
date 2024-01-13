package main

import (
	"cook/cook07/handler"
	"fmt"
	"net/http"
)
///home/start/go/src/cook/cook07/handler
func main(){
	http.HandleFunc("/name", handler.HelloHandler)
	http.HandleFunc("/greeting", handler.GreetingHandler)
	fmt.Println("Listening on port :3333")
	err := http.ListenAndServe(":3333", nil)
	panic(err)
}

//curl "http://localhost:3333/name?name=Reader" -X GET
//curl "http://localhost:3333/greeting" -X POST -d 'greeting=Goodbye&name=Reader'