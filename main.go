package main

import (
	"log"
	"net/http"
	"time"

	"github.com/derektamsen/hancock/aws"
	"github.com/stevenroose/gonfig"
)

// config holds configuration values for app
var config struct {
	ConfigFile  string `short:"f" desc:"Path to config file"`
	ListenAddr  string `short:"a" default:"0.0.0.0" desc:"Address to listen for connections"`
	ListenPort  string `short:"p" default:"8080" desc:"Port the service listens for requests"`
	S3Bucket    string `short:"b" default:"testbucket" desc:"Bucket assets are located"`
	PresignTime int    `short:"t" default:"15" desc:"time in minutes url is valid"`
	AWSSvc      string `short:"s" default:"s3" desc:"service to generate presigned url for"`
}

// httpHandler signs urls using awsurl.S3PreSign then redirects the user to the new location
func httpHandler(w http.ResponseWriter, r *http.Request) {
	var presignedURL string
	switch config.AWSSvc {
	case "s3":
		presignedURL = awsurl.S3PreSign(r.URL.Path[1:], config.S3Bucket, config.PresignTime)
	default:
		log.Fatal("AWS Service not supported: ", config.AWSSvc)
	}

	http.Redirect(w, r, presignedURL, http.StatusFound)
}

// main starts a http server listening for requests on <ListenAddr>:<ListenPort>
func main() {
	err := gonfig.Load(&config, gonfig.Conf{
		ConfigFileVariable:  "configfile", // enables passing --configfile myfile.conf
		FileDefaultFilename: "hancock.conf",
		FileDecoder:         gonfig.DecoderTOML,
		EnvPrefix:           "HANCOCK_",
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Print("Starting Hancock signing server")

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
