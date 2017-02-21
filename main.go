package main

import (
	"crypto/sha512"
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Path
	hashed_url := shasum(url)
	fmt.Fprintf(w, "sha: %x\nrequest: %s\n", hashed_url, url)
}

func shasum(s string) []byte {
	sha := sha512.New()
	sha.Write([]byte(s))
	return sha.Sum(nil)
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
