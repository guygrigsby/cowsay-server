package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/inconshreveable/log15"
)

const (
	DefaultCowsay = "/usr/games/cowsay"
	DefaultPort   = 8080
)

func gmain() {
}

func main() {
	log := log15.New()
	config := FromEnv()
	log.Debug(
		"Config",
		"Tokens", fmt.Sprintf("%+v", config.Tokens()),
		"Cert Path", config.CertFile(),
		"Key Path", config.KeyFile(),
		"Port", config.Port(),
		"CowsayExec", config.CowsayExec(),
	)
	tokens := make(map[string]bool)
	for _, token := range config.Tokens() {
		log.Debug(
			"Adding token",
			"Token", token,
		)
		tokens[token] = true
	}

	prog := config.CowsayExec()
	if prog == "" {
		prog = DefaultCowsay
		log.Info(
			"Using default cowsay location",
			"Location", prog,
		)
	}
	addr := fmt.Sprintf(":%d", config.Port())
	r := gin.Default()

	// Ping handler
	r.POST("/", cowsayHandler(prog, tokens, log))

	//log.Error(autotls.Run(r, "cowsay.guygrigsby.com").Error())
	log.Error(r.Run(addr).Error())

}
