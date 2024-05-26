package server

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"easyproxy/pkg/utils"
)

type Server struct {
	client *http.Client

	headers http.Header

	port int
}

func New() Server {
	headers := &http.Header{}
	headers.Add("Proxied-From", "easyproxy")

	s := Server{
		client:  &http.Client{},
		port:    8080,
		headers: *headers,
	}

	return s
}

func (s *Server) checkResponse(r *http.Response) error {
	// Check if the response is nil
	if r == nil {
		log.Println("r nil")
		return fmt.Errorf("response is nil")
	}

	// Check if the response body is nil
	if r.Body == nil {
		log.Println("r body")
		return fmt.Errorf("response body is nil")
	}

	return nil
}

func (s *Server) buildProxiedResponse(r *http.Response, rw http.ResponseWriter) error {
	// Need to copy the request data from http.Response to http.ResponseWrite
	rwHeaders := rw.Header()

	headersToAdd := utils.MergeHeaders(r.Header, s.headers)
	// copy headers
	for key, values := range headersToAdd {
		for _, v := range values {
			rwHeaders.Add(key, v)
		}
	}

	// copy status code
	rw.WriteHeader(r.StatusCode)

	buf := make([]byte, 8)
	if _, err := io.CopyBuffer(rw, r.Body, buf); err != nil {
		// handle the error
		return err
	}

	return r.Body.Close()
}

func (s *Server) proxy(w http.ResponseWriter, r *http.Request) {
	response, err := s.client.Get(r.URL.String())
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = s.checkResponse(response)
	if err != nil {
		log.Println("Check response failed: %w", err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	err = s.buildProxiedResponse(response, w)
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
