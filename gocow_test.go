package main

import (
	"fmt"
	"testing"

	cowsay "github.com/Code-Hex/Neo-cowsay"
)

func TestGoCow(t *testing.T) {
	say, err := cowsay.Say(
		cowsay.Phrase("Hello"),
		cowsay.Type("default"),
		cowsay.BallonWidth(40),
	)
	if err != nil {
		panic(err)
	}
	fmt.Println(say)
}
