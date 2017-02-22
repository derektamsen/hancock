package main

import (
	"log"
	"net/http"
	"time"

	"github.com/derektamsen/hancock/aws"
	"github.com/derektamsen/hancock/opts"
)

// These are default configs and can be overridden in ./config.yaml or ENV vars prefixed
// with HANCOCK_<var_name>
var defConfigFile = "config"   // config file prefix name. Can end in yaml, json, toml
var defListenPort = "8080"     // port the service listens for requests at
var defAWSRegion = "us-east-1" // aws region service bucket is located
var defS3Bucket = "testbucket" // bucket assets are located
var defPresignTime = 15        // time in minutes url is valid
var defAWSSvc = "s3"           // service to generate presigned url for

// httpHandler signs urls using awsurl.S3PreSign then redirects the user to the new location
func httpHandler(w http.ResponseWriter, r *http.Request) {
	awsSvc := opts.GetConfigString("aws_svc", defAWSSvc)
	var presignedURL string
	switch awsSvc {
	case "s3":
		presignedURL = awsurl.S3PreSign(r.URL.Path[1:],
			opts.GetConfigString("aws_region", defAWSRegion),
			opts.GetConfigString("s3_bucket", defS3Bucket),
			opts.GetConfigInt("presign_time", defPresignTime),
		)
	default:
		log.Fatal("AWS Service not supported: ", awsSvc)
	}

	http.Redirect(w, r, presignedURL, http.StatusFound)
}

// main starts a http server listening for requests on <listen_host>:<listen_port>
func main() {
	log.Print("Starting Hancock signing server")
	opts.SetConfigLoad("hancock", defConfigFile)

	listenPort := opts.GetConfigString("listen_port", defListenPort)

	httpServer := &http.Server{
		Addr:           ":" + listenPort,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Print("Listening on port: ", listenPort)
	http.HandleFunc("/", httpHandler)
	log.Fatal(httpServer.ListenAndServe())
}
