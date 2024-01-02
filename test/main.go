package main

import (
    "encoding/json"
    "net/http"

    "github.com/gorilla/mux"
)

type Message struct {
    Text string `json:"text"`
}

func corsMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Access-Control-Allow-Origin", "*")
        w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
        next.ServeHTTP(w, r)
    })
}

func main() {
    router := mux.NewRouter()

    router.Use(corsMiddleware)
    
    router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "client/build/index.html")
	})

    router.HandleFunc("/api/message", func(w http.ResponseWriter, r *http.Request) {
        message := Message{Text: "Hello from Go!"}
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(message)
    })

    router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("client/build/static"))))

    http.Handle("/", router)
    http.ListenAndServe(":8080", nil)
}
