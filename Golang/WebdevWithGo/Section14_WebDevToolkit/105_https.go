package main

import (
	"net/http"
)

//go run /usr/local/go/src/crypto/tls/generate_cert.go --host=localhost
// access with https://localhost:10443/

func main() {
	http.HandleFunc("/", a)
	http.ListenAndServeTLS(":10443", "cert.pem", "key.pem", nil)
}

func a(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("This is an example server.\n"))
}
