package main

import (
	"log"
	"net/http"
	"time"

	"github.com/derektamsen/hancock/aws"
	"github.com/derektamsen/hancock/opts"
)

// Config holds the configuration settings for the service
// These are default configs and can be overridden in ./config.yaml or ENV vars prefixed
// with HANCOCK_<var_name>
type Config struct {
	File        string // config file prefix name. Can end in yaml, json, toml
	ListenAddr  string // address to listen for connections on
	ListenPort  string // port the service listens for requests at
	AWSRegion   string // aws region service bucket is located
	S3Bucket    string // bucket assets are located
	PresignTime int    // time in minutes url is valid
	AWSSvc      string // service to generate presigned url for
}

// holds configuration values
var config Config

// httpHandler signs urls using awsurl.S3PreSign then redirects the user to the new location
func httpHandler(w http.ResponseWriter, r *http.Request) {
	var presignedURL string
	switch config.AWSSvc {
	case "s3":
		presignedURL = awsurl.S3PreSign(r.URL.Path[1:], config.AWSRegion, config.S3Bucket, config.PresignTime)
	default:
		log.Fatal("AWS Service not supported: ", config.AWSSvc)
	}

	http.Redirect(w, r, presignedURL, http.StatusFound)
}

// main starts a http server listening for requests on <listen_host>:<listen_port>
func main() {
	config.File = opts.GetConfigString("aws_svc", "config")

	log.Print("Starting Hancock signing server")
	opts.SetConfigLoad("hancock", config.File)

	config.ListenAddr = opts.GetConfigString("listen_address", "0.0.0.0")
	config.ListenPort = opts.GetConfigString("listen_port", "8080")
	config.AWSRegion = opts.GetConfigString("aws_region", "us-east-1")
	config.S3Bucket = opts.GetConfigString("s3_bucket", "testbucket")
	config.PresignTime = opts.GetConfigInt("presign_time", 15)
	config.AWSSvc = opts.GetConfigString("aws_service", "s3")

	httpServer := &http.Server{
		Addr:           config.ListenAddr + ":" + config.ListenPort,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Print("Listening on ", httpServer.Addr)
	http.HandleFunc("/", httpHandler)
	log.Fatal(httpServer.ListenAndServe())
}
