package main

import (
	"fmt"

	"github.com/gin-gonic/autotls"
	"github.com/gin-gonic/gin"
	"github.com/inconshreveable/log15"
)

const (
	DefaultCowsay = "/usr/games/cowsay"
	DefaultPort   = 80
)

func gmain() {
}

func main() {
	log := log15.New()
	config := FromEnv()
	log.Debug(
		"Config",
		"Tokens", fmt.Sprintf("%+v", config.Tokens()),
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
	r := gin.Default()

	r.POST("/", cowsayHandler(prog, tokens, log))
	log.Info(
		"Using auto TLS",
	)

	if config.TLS() {
		err := autotls.Run(r, config.Domain())
		if err != nil {
			log.Error(
				"TLS failure.",
				"Error", err,
			)
		}
	}

	log.Error(r.Run(":80").Error())

}
