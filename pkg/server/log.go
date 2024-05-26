package server

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func generateAccessLog(r *http.Request, startTime time.Time) {
	logStr := fmt.Sprintf("[%s] %s %s", startTime.Format(time.RFC3339), r.Method, r.URL.String())

	log.Println(logStr)
}
