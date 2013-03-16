package manners

import (
	"net"
	"net/http"
	"sync"
	"testing"
)

func TestServerCallsShutdownTasksWhenShutdown(t *testing.T) {
	fakeHandler := FakeHandler{}
	taskCounter := sync.WaitGroup{}
	oldListener, err := net.Listen("tcp", ":7000")
	if err != nil {
		t.Error(err.Error())
	}
	oldServer := http.Server{Handler: fakeHandler}
	gracefulListener := NewGracefulListener(oldListener, &taskCounter)

	testChannel := make(chan bool)

	server := NewGracefulServer(&gracefulListener, &taskCounter, oldServer)
	server.ShutdownTasks = func() { testChannel <- true }

	go server.Serve()
	server.Shutdown()

	// This will cause a deadlock if the server does not call the ShutdownTasks
	<-testChannel
}
