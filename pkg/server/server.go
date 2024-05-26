package server

import (
	"easyproxy/pkg/utils"
	"fmt"
	"log"
	"net/http"
)
type Server struct {
	client *http.Client

	port int
}

func New() (Server){
	s := Server{
		client: &http.Client{},
		port: 8080,
	}
	
	return s
}


func (s *Server) proxy(w http.ResponseWriter, r *http.Request) {
	response, err := s.client.Get(r.URL.String())
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	} 

	err = utils.CopyResponseToResponseWriter(response, w)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	} 

}

func (s *Server) Run() {
    // Register the handler function for the root URL path
    http.HandleFunc("/", s.proxy)

    // Specify the address and port to listen on
    addr := fmt.Sprintf(":%d", s.port)
    log.Printf("Starting server on %s\n", addr)

    // Start the HTTP server
    if err := http.ListenAndServe(addr, nil); err != nil {
        log.Fatalf("Could not start server: %s\n", err)
    }
}
