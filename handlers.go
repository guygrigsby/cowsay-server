package main

import (
	"net/http"
	"os/exec"

	"github.com/gin-gonic/gin"
	"github.com/inconshreveable/log15"
)

func cowsayHandler(cowsayExec string, allowed map[string]interface{}, log log15.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		if len(allowed) > 0 {
			log.Info(
				"Processing tokens",
				"Count", len(allowed),
			)
			token, _ := c.GetPostForm("token")
			_, ok := allowed[token]

			if !ok {
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
		c.Header("content-type", "application/json")
		c.JSON(http.StatusOK, back)

	}
}

type CowsayResponse struct {
	Response_type string   `json:"response_type"`
	Text          string   `json:"text"`
	Attachments   []string `json:"attachments"`
}
