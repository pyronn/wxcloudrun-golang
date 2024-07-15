package main

import (
	"fmt"
	"log"
	"net/http"
	"wxcloud-test/db"
	"wxcloud-test/service"
)

func main() {
	if err := db.Init(); err != nil {
		panic(fmt.Sprintf("mysql init failed with %+v", err))
	}

	http.HandleFunc("/", service.IndexHandler)
	http.HandleFunc("/api/count", service.CounterHandler)
	http.HandleFunc("/api/recMsg", service.MsgHandler)

	log.Fatal(http.ListenAndServe(":80", nil))
}
