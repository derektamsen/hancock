package main

import (
	"crypto/sha512"
	"github.com/derektamsen/awss3urlsigner/aws"
	"log"
	"net/http"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	presigned_url := awsurl.PreSign(r.URL.Path[1:])
	http.Redirect(w, r, presigned_url, http.StatusFound)
}

func shasum(s string) []byte {
	sha := sha512.New()
	sha.Write([]byte(s))
	return sha.Sum(nil)
}

func main() {
	httpserver := &http.Server{
		Addr:           ":8080",
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	http.HandleFunc("/", handler)
	log.Fatal(httpserver.ListenAndServe())
}
