package main

import (
	"gsfilm/go-proxy-newbing/api"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	http.HandleFunc("/sydney/ChatHub", api.ChatHub)

	http.HandleFunc("/web/", api.WebStatic)

	http.HandleFunc("/", api.Index)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	addr := ":" + port

	log.Println("Starting newbing Proxy At " + addr)

	srv := &http.Server{
		Addr:         addr,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
