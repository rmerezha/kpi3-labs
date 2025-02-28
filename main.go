package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/time", handleFunc)
	server := http.Server{
		Addr:    ":8795",
		Handler: mux,
	}
	log.Println("Server is starting, port: 8795")
	err := server.ListenAndServe()
	if err != nil {
		log.Println(err.Error())
	}
}

func handleFunc(w http.ResponseWriter, r *http.Request) {
	log.Println("New requerst from " + r.RemoteAddr)
	t := Resp{
		Time: time.Now().Format(time.RFC3339),
	}
	log.Println(t)
	json, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		log.Println(err.Error())
	}
	log.Println("Send response")
	w.Write(json)
}

type Resp struct {
	Time string `json:"time"`
}
