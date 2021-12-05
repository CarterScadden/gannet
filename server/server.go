package server

import (
	"fmt"
	"gannet/server/handlers"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// Run
// used to initialze and run the server, takes a argument of the portnumber to use
func New(port int) *mux.Router {
	// TODO: routers
	r := mux.NewRouter()
	r.Use(loggerMiddleware)

	r.HandleFunc("/ping", handlers.Ping).Methods("GET")
	r.HandleFunc("/api/v1/produce", handlers.GetProduce).Methods("GET")
	r.HandleFunc("/api/v1/produce", handlers.PostProduce).Methods("POST")
	// r.HandleFunc("/api/v1/produce", handlers.PostProduce).Methods("POST")

	http.Handle("/", r)

	return r
}

// logger
// used as middleware to log the `[time] method url` of each request
func loggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("[%s] %s %s\n", time.Now(), r.Method, r.URL)
		next.ServeHTTP(w, r)
	})
}
