package manners

import (
	"errors"
	"net"
	"sync"
)

var CLOSED_LISTENER_ERROR = "This listener is closed."

func NewGracefulListener(oldListener net.Listener, taskCounter *sync.WaitGroup) GracefulListener {
	listener := GracefulListener{oldListener, true, taskCounter}
	return listener
}

type GracefulListener struct {
	net.Listener
	open        bool
	taskCounter *sync.WaitGroup
}

func (this *GracefulListener) Accept() (net.Conn, error) {
	conn, err := this.Listener.Accept()
	if err != nil {
		if !this.open {
			err = errors.New(CLOSED_LISTENER_ERROR)
		}
		return nil, err
	}
	this.taskCounter.Add(1)
	return GracefulConnection{conn, this.taskCounter}, nil
}

func (this *GracefulListener) Close() error {
	if !this.open {
		return nil
	}
	this.open = false
	err := this.Listener.Close()
	return err
}
