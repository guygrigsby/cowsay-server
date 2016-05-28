package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os/exec"

	"github.com/inconshreveable/log15"
)

func cowsayHandler(tokens map[string]bool, log log15.Logger) http.HandlerFunc {
	return http.HandlerFunc(

		func(w http.ResponseWriter, r *http.Request) {
			defer r.Body.Close()
			err := r.ParseForm()

			if err != nil {
				log.Error(
					"Cannot parse form",
					"Error", err,
					"Form", fmt.Sprintf("%+v", r.Form),
				)
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			if len(tokens) > 0 {
				token := r.PostFormValue("token")
				if tokens[token] == false {
					log.Error(
						"Token Rejected",
						"Token", token,
					)
					w.WriteHeader(http.StatusForbidden)
					return
				}

			}

			text := r.PostFormValue("text")
			prog := "/usr/games/cowsay"
			out, err := exec.Command(prog, string(text)).Output()
			if err != nil {
				log.Error(
					"Cowsay not found",
					"Error", err,
					"Location", prog,
				)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			codeMark := []byte("```")
			out = append(out, codeMark...)
			out = append(codeMark, out...)
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
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			w.Header().Set("content-type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write(resp)
		},
	)
}

type CowsayResponse struct {
	Response_type string   `json:"response_type"`
	Text          string   `json:"text"`
	Attachments   []string `json:"attachments"`
}
