package main

import (
    "net/http"
    "github.com/AgoraIO-Community/agora-token-service/service"
)

func main() {
    s := service.NewService()
    http.HandleFunc("/", CORS(s.Start))
    http.ListenAndServe(":8080", nil)
}

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
