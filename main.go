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
	allowed := make(map[string]interface{})
	for _, token := range config.Tokens() {
		log.Debug(
			"Adding token",
			"Token", token,
		)
		allowed[token] = struct{}{}
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

	r.POST("/", cowsayHandler(prog, allowed, log))
	r.GET("/healthz", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "not-dead",
		})
	})

	if config.TLS() {
		log.Info(
			"Using auto TLS",
		)
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
