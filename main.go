package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/inconshreveable/log15"
)

const (
	DefaultCowsay = "/usr/games/cowsay"
)

func main() {
	configFile := flag.String("c", "./config.json", "Path to configuration file")
	flag.Parse()
	log := log15.New()
	b, err := ioutil.ReadFile(*configFile)
	if err != nil {
		log.Error(
			"Cannot read config",
			"Error", err,
		)
	}
	var config Config
	err = json.Unmarshal(b, &config)
	if err != nil {
		log.Error(
			"Invalid JSON in config",
			"Error", err,
		)
	}
	log.Debug(
		"Config",
		"Tokens", fmt.Sprintf("%+v", config.Tokens),
		"Cert Path", config.CertFile,
		"Key Path", config.KeyFile,
		"ListenOn", config.ListenOn,
		"CowsayExec", config.CowsayExec,
	)
	tokens := make(map[string]bool)
	for _, token := range config.Tokens {
		log.Debug(
			"Adding token",
			"Token", token,
		)
		tokens[token] = true
	}

	prog := config.CowsayExec
	if prog == "" {
		prog = DefaultCowsay
		log.Info(
			"Using default cowsay location",
			"Location", prog,
		)
	}
	mux := http.NewServeMux()
	mux.Handle(
		"/cowsay",
		cowsayHandler(prog, tokens, log),
	)
	log.Info(
		"Starting server...",
	)
	for err = http.ListenAndServeTLS(config.ListenOn, config.CertFile, config.KeyFile, mux); err != nil; {
		time.Sleep(time.Duration(2) * time.Second)
		log.Crit("Restarting", "Error", err)
	}

}
