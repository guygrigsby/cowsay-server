package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

type cfg struct {
	tokens []string
	cert   string
	key    string
	port   int
}

func FromEnv() Config {
	conf := &cfg{}
	conf.tokens = getEnvValCSV("COWSAY_TOKENS")
	conf.cert = getEnvVal("COWSAY_TLS_CERT")
	conf.key = getEnvVal("COWSAY_TLS_KEY")
	port, err := getEnvValInt("COWSAY_PORT")
	if err != nil {
		conf.port = 8080
	}
	conf.port = port
	return conf
}

func (c *cfg) Tokens() []string {
	return c.tokens
}

func (c *cfg) CertFile() string {
	return c.cert
}

func (c *cfg) KeyFile() string {
	return c.key
}

func (c *cfg) ListenOn() string {
	return fmt.Sprintf("%d", c.port)
}

func (c *cfg) CowsayExec() string {
	uri, err := whichCowsay()
	if err != nil {
		return DefaultCowsay
	}
	return uri
}

func whichCowsay() (string, error) {
	cmd := exec.Command("which", "cowsay")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return "", err
	}
	return out.String(), nil
}

type Config interface {
	Tokens() []string
	CertFile() string
	KeyFile() string
	ListenOn() string
	CowsayExec() string
}

func getEnvValCSV(envVars ...string) []string {
	str := getEnvVal(envVars...)
	return strings.Split(str, ",")

}

func getEnvValInt(envVars ...string) (int, error) {

	str := getEnvVal(envVars...)

	val, err := strconv.Atoi(str)

	return val, err
}

func getEnvVal(envVars ...string) string {

	for _, envVar := range envVars {

		if v := os.Getenv(envVar); v != "" {
			return v
		}
	}

	return ""
}
