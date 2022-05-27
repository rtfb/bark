package main

import (
	"errors"
	"net/http"
	"time"

	"github.com/rtfb/bark"
)

func main() {
	l := bark.Create()
	e := errors.New("induced")
	l.Log(e)

	l2 := bark.CreateFile("test.log")
	req, _ := http.NewRequest(http.MethodGet, "/?q=foo", nil)
	req.ParseForm()
	l2.LogRq(req, time.Now())
}
