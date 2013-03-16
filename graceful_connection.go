package manners

import (
	"net"
	"sync"
)

type GracefulConnection struct {
	net.Conn
	taskCounter *sync.WaitGroup
}

func (this GracefulConnection) Close() error {
	err := this.Conn.Close()
	this.taskCounter.Done()
	return err
}
