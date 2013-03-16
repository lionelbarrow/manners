package manners

import (
	"net"
	"net/http"
	"sync"
)

func ListenAndServe(addr string, handler http.Handler) error {
	oldListener, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	tasks := sync.WaitGroup{}
	listener := NewGracefulListener(oldListener, &tasks)
	normalServer := http.Server{Handler: handler}
	gracefulsServer := NewGracefulServer(&listener, &tasks, normalServer)
	gracefulsServer.Serve()
	return err
}
