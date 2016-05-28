package main

type Config struct {
	Tokens   []string `json:"tokens"`
	CertFile string   `json:"certfile"`
	KeyFile  string   `json:"keyfile"`
	ListenOn string   `json:"listenon"`
}
