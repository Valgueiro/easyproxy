package server

import (
	"easyproxy/pkg/utils"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

var supportedSchemes = []string{"http", "https"}

type Proxy struct {
	headers http.Header
}

// var defaultHeaders = &http.Header{}
// defaultHeaders.Add("Proxied-From", "easyproxy")

func (p *Proxy) getClient() *http.Client {
	return &http.Client{}
}

func (p *Proxy) prepareRequest(req *http.Request) error {
	//http: Request.RequestURI can't be set in client requests.
	//http://golang.org/src/pkg/net/http/client.go
	req.RequestURI = ""

	// TODO: during CONNECT scheme is empty which is leading to
	// "Connect "//httpbin.org:443": unsupported protocol scheme """
	// We need to find a way to fix this here
	// Set url manually? Add a scheme?

	// if req.URL.Scheme == "" {
	// 	req.Ur
	// }

	return nil
}

func (p *Proxy) checkRequest(req *http.Request) error {
	// Check if the response is nil
	if req == nil {
		return fmt.Errorf("req is nil")
	}

	// Check if the response body is nil
	// if !slices.Contains(supportedSchemes, req.URL.Scheme) {
	// 	return fmt.Errorf("unsupported scheme %s", req.URL.Scheme)
	// }

	return nil
}

func (p *Proxy) prepareResponse(r *http.Response, rw http.ResponseWriter) error {
	// Need to copy the request data from http.Response to http.ResponseWrite
	rwHeaders := rw.Header()

	headersToAdd := utils.MergeHeaders(r.Header, p.headers)
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

func (p *Proxy) checkResponse(r *http.Response) error {
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

func (p *Proxy) ServeHTTP(wr http.ResponseWriter, req *http.Request) {
	startTime := time.Now()
	err := p.checkRequest(req)

	if err != nil {
		log.Println(fmt.Errorf("check request failed: %w", err))
		return
	}

	err = p.prepareRequest(req)
	if err != nil {
		log.Println(fmt.Errorf("could not prepare the request: %w", err))
		return
	}

	c := p.getClient()
	log.Println(req)
	response, err := c.Do(req)
	if err != nil {
		log.Println(fmt.Errorf("could not do the request: %w", err))
		return
	}

	err = p.checkResponse(response)
	if err != nil {
		log.Println(fmt.Errorf("check response failed: %w", err))
		return
	}

	// print status
	// log.Println("Status code: %d", response.StatusCode)

	// copy header from response to wr
	err = p.prepareResponse(response, wr)
	if err != nil {
		log.Println(fmt.Errorf("prepare response failed: %w", err))
		wr.WriteHeader(http.StatusInternalServerError)
		return
	}

	generateAccessLog(req, startTime)
}
