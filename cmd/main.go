package main

import (
	"errors"

	"github.com/rtfb/bark"
)

func main() {
	l := bark.Create()
	e := errors.New("induced")
	l.Log(e)
}
