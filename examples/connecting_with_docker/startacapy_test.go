package main

import (
	"flag"
	"fmt"
	"github.com/ldej/go-acapy-client"
	"math/rand"
	"strconv"
	"testing"
	"time"
)

func TestApp_StartACApyWithDocker(t *testing.T) {
	var port = 4455
	var ledgerURL = "http://localhost:9000"
	var name = ""

	flag.IntVar(&port, "port", 4455, "port")
	flag.StringVar(&name, "name", "Alice", "alice")
	flag.Parse()

	acapyURL := fmt.Sprintf("http://localhost:%d", port+2)

	app := App{
		client:    acapy.NewClient(acapyURL),
		ledgerURL: ledgerURL,
		port:      port,
		label:     name,
		rand:      strconv.Itoa(rand.New(rand.NewSource(time.Now().UnixNano())).Intn(100000)),
	}
	app.StartACApyWithDocker()
}
