package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	// listen for system shutdown
	stop := make(chan os.Signal)
	signal.Notify(stop, os.Interrupt)

	r := handleRoutes()

	srv := &http.Server{Handler: r, Addr: ":8080"}

	go func() {
		log.Printf("Listening on port 8080")
		if err := srv.ListenAndServe(); err != nil {
			log.Printf("listen: %s\n", err)
		}
	}()

	<-stop

	log.Printf("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	srv.Shutdown(ctx)

	log.Printf("Shutting gracefully shutdown")
}

func handleRoutes() *mux.Router {
	r := mux.NewRouter()

	s := r.PathPrefix("/api/").Subrouter()
	s.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusForbidden)
	})

	r.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("./views"))))

	return r
}
