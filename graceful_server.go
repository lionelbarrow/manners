package manners

import (
	"net/http"
	"sync"
)

func NewGracefulServer(listener *GracefulListener, tasks *sync.WaitGroup, server http.Server) GracefulServer {
	return GracefulServer{
		ShutdownTasks:        func() {},
		server:               server,
		listener:             listener,
		tasks:                tasks,
		notAcceptingRequests: make(chan bool),
		readyToShutdown:      make(chan bool),
	}
}

type GracefulServer struct {
	ShutdownTasks func()

	server               http.Server
	listener             *GracefulListener
	tasks                *sync.WaitGroup
	notAcceptingRequests chan bool
	readyToShutdown      chan bool
}

func (this *GracefulServer) Serve() error {
	go this.waitForShutdown()
	err := this.server.Serve(this.listener)
	if err == nil {
		return nil
	} else if err.Error() == CLOSED_LISTENER_ERROR {
		<-this.readyToShutdown
		return nil
	}
	return err
}

func (this *GracefulServer) waitForShutdown() {
	<-this.notAcceptingRequests
	this.listener.Close()
	this.tasks.Wait()
	this.ShutdownTasks()
	this.readyToShutdown <- true
}

func (this *GracefulServer) Shutdown() {
	this.notAcceptingRequests <- true
}
