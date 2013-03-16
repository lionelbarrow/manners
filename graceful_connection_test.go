package manners

import (
	"sync"
	"testing"
)

func TestCloseDecrementsCounter(t *testing.T) {
	taskCounter := sync.WaitGroup{}
	taskCounter.Add(1)
	testConn := GracefulConnection{FakeConnection{}, &taskCounter}

	testConn.Close()
	// Will deadlock if the connection does not work correctly
	taskCounter.Wait()
}
