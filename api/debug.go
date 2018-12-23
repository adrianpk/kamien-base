package api

import (
	"bytes"
	"fmt"
	"net/http"
	"strings"
)

func debugRequestHeader(r *http.Request) {
	log.Debugf("Request header:\n%v", r.Header)
}

func debugRequestBody(r *http.Request) {
	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)
	str := buf.String()
	log.Debugf("Request body\n%s", str)
}

func debugRequest(r *http.Request) {
	log.Debug(formatRequest(r))
}

// formatRequest generates ascii representation of a request
func formatRequest(r *http.Request) string {
	// Create return string
	var request []string
	// Add the request string
	url := fmt.Sprintf("%v %v %v", r.Method, r.URL, r.Proto)
	request = append(request, url)
	// Add the host
	request = append(request, fmt.Sprintf("Host: %v", r.Host))
	// Loop through headers
	for name, headers := range r.Header {
		name = strings.ToLower(name)
		for _, h := range headers {
			request = append(request, fmt.Sprintf("%v: %v", name, h))
		}
	}
	// If this is a POST, add post data
	if r.Method == "POST" || r.Method == "PUT" || r.Method == "PATCH" {
		r.ParseForm()
		request = append(request, "\n")
		request = append(request, r.Form.Encode())
	}
	// Return the request as a string
	return strings.Join(request, "\n")
}
