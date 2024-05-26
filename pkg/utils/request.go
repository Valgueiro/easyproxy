package utils

import "net/http"

func MergeHeaders(a http.Header, b http.Header) http.Header {
	o := http.Header{}

	for key, value := range a {
		o[key] = value
	}

	for key, value := range b {
		if aVal, ok := o[key]; ok {
			o[key] = append(aVal, value...)
		} else {
			o[key] = value
		}
	}

	return o
}
