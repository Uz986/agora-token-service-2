package main

import (
	"github.com/AgoraIO-Community/agora-token-service/service"
	"net/http"
	"github.com/rs/cors"
)

func main() {
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	})
	
	http.HandleFunc("/", c.HandlerFunc(myHandler))
	http.ListenAndServe(":8080", nil)
	
	s := service.NewService()
	// Stop is called on another thread, but waits for an interrupt
	go s.Stop()
	s.Start()
}

func myHandler(w http.ResponseWriter, r *http.Request) {
	// Your handler code here
}
