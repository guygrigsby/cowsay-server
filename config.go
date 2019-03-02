package main

import (
	"bytes"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

type cfg struct {
	tokens []string
	domain string
	tls    bool
}

func FromEnv() Config {
	conf := &cfg{}
	conf.tokens = getEnvValCSV("COWSAY_TOKENS")
	conf.domain = getEnvVal("COWSAY_TLS_DOMAIN")
	conf.tls = getEnvVal("COWSAY_AUTO_TLS") == "TRUE"
	return conf
}

func (c *cfg) Tokens() []string {
	return c.tokens
}
func (c *cfg) CowsayExec() string {
	uri, err := whichCowsay()
	if err != nil {
		return DefaultCowsay
	}
	return uri
}
func (c *cfg) Domain() string {
	return c.domain
}
func (c *cfg) TLS() bool {
	return c.tls
}

func whichCowsay() (string, error) {
	cmd := exec.Command("which", "cowsay")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return "", err
	}
	return strings.TrimRight(out.String(), "\n"), nil
}

type Config interface {
	Tokens() []string
	CowsayExec() string
	Domain() string
	TLS() bool
}

func getEnvValCSV(envVars ...string) []string {
	str := getEnvVal(envVars...)
	if str == "" {
		return nil
	}
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
