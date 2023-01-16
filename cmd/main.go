package main

import (
    "net/http"
	"github.com/AgoraIO-Community/agora-token-service/service"
)

func CORS(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        w.Header().Add("Access-Control-Allow-Origin", "*")
        w.Header().Add("Access-Control-Allow-Credentials", "true")
        w.Header().Add("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
        w.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")

        if r.Method == "OPTIONS" {
            http.Error(w, "No Content", http.StatusNoContent)
            return
        }

        next(w, r)
    }
}

func startWithCORS(w http.ResponseWriter, r *http.Request) {
    s := service.NewService()
    // Stop is called on another thread, but waits for an interrupt
    go s.Stop()
    s.Start()
}

func main() {
    http.HandleFunc("/", CORS(startWithCORS))
    http.ListenAndServe(":8080", nil)
}
