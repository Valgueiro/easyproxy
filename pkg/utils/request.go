package utils

import (
	"fmt"
	"io"
	"log"
	"net/http"
)


func CopyResponseToResponseWriter(r *http.Response, rw http.ResponseWriter) error {
	// Need to copy the request data from http.Response to http.ResponseWriter
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
	
	// copy status code
	rw.WriteHeader(r.StatusCode)

	// copy headers
	for key, values := range r.Header {
		for _, v := range values {
			rw.Header().Add(key, v)
		}
	}
	buf := make([]byte, 8)
    if _, err := io.CopyBuffer(rw, r.Body, buf); err != nil {
        // handle the error
        return err
    }

	return r.Body.Close()
}
