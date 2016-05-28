package main

type Config struct {
	Tokens   []string `json:"tokens,omitempty"`
	CertFile string   `json:"certfile"`
	KeyFile  string   `json:"keyfile"`
	ListenOn string   `json:"listenon"`
}
