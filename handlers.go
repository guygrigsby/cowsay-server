package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	cowsay "github.com/Code-Hex/Neo-cowsay"
	"github.com/inconshreveable/log15"
)

func cowsayHandler() http.HandlerFunc {
	return http.HandlerFunc(

		func(w http.ResponseWriter, r *http.Request) {
			log := log15.New()
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

			text := r.PostFormValue("text")

			say, err := cowsay.Say(
				cowsay.Phrase(text),
				cowsay.Type("default"),
			)

			if err != nil {
				log.Error(
					"Cowsay error",
					"Error", err,
				)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			codeMark := []byte("```")
			out := append(codeMark, say...)
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
