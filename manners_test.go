package manners

import (
	"net"
	"net/http"
	"sync"
	"testing"
	"time"
)

var (
	requestFinished    = make(chan bool)
	requestBeingServed = make(chan bool)
	shutdownSent       = make(chan bool)
)

type testHandler struct{}

func (this testHandler) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	requestBeingServed <- true
	<-shutdownSent
	time.Sleep(1e9)
	requestFinished <- true
}

func TestGracefulness(t *testing.T) {
	taskCounter := sync.WaitGroup{}
	oldListener, err := net.Listen("tcp", ":7000")
	if err != nil {
		t.Error(err.Error())
	}
	gracefulListener := NewGracefulListener(oldListener, &taskCounter)

	handler := testHandler{}
	oldServer := http.Server{Handler: handler}
	server := NewGracefulServer(&gracefulListener, &taskCounter, oldServer)

	go server.Serve()
	go http.Get("http://localhost:7000")
	<-requestBeingServed
	server.Shutdown()
	shutdownSent <- true
	select {
	case <-requestFinished:
	case <-time.After(2e9):
		t.Error("The request did not run completion")
	}
}
