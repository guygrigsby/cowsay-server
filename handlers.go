package main

import (
	"encoding/json"
	"net/http"
	"os/exec"

	"github.com/gin-gonic/gin"
	"github.com/inconshreveable/log15"
)

func cowsayHandler(cowsayExec string, tokens map[string]bool, log log15.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		if len(tokens) > 0 {
			token, _ := c.GetPostForm("token")

			if tokens[token] == false {
				log.Error(
					"Token Rejected",
					"Token", token,
				)
				c.AbortWithStatus(http.StatusForbidden)
				return
			}

		}

		text, _ := c.GetPostForm("text")
		out, err := exec.Command(cowsayExec, string(text)).Output()
		if err != nil {
			log.Error(
				"Cowsay not found",
				"Error", err,
				"Location", cowsayExec,
			)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		codeMark := []byte("```")
		out = append(codeMark, out...)
		out = append(out, codeMark...)
		back := CowsayResponse{
			Response_type: "in_channel",
			Text:          string(out),
		}

		resp, err := json.Marshal(back)
		if err != nil {
			log.Error(
				"Cannot marshal response",
				"Error", err,
				"Response", string(resp),
			)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		c.Header("content-type", "application/json")
		c.JSON(http.StatusOK, resp)

	}
}

type CowsayResponse struct {
	Response_type string   `json:"response_type"`
	Text          string   `json:"text"`
	Attachments   []string `json:"attachments"`
}
