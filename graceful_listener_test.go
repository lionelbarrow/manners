package manners

import (
	"sync"
	"testing"
)

func TestAcceptIncrementsCounter(t *testing.T) {
	taskCounter := sync.WaitGroup{}
	fakeListener := FakeListener{open: true}

	listener := NewGracefulListener(fakeListener, &taskCounter)
	listener.Accept()

	// WaitGroups panic when their internal counter goes negative. If accept did
	// not increment the counter, this will panic.
	taskCounter.Done()
}

func TestAcceptReturnsAMannersErrorWhenCalledAfterClose(t *testing.T) {
	taskCounter := sync.WaitGroup{}
	fakeListener := FakeListener{}

	listener := NewGracefulListener(fakeListener, &taskCounter)
	err := listener.Close()
	if err != nil {
		t.Error("Unexpected error when closing listener: " + err.Error())
	}
	_, err = listener.Accept()
	if err == nil {
		t.Error("Did not receive an error after calling Accept on a closed Listener.")
	}
	if err.Error() != "This listener is closed." {
		t.Error("Did not receive a manners error after calling Accept() on a closed listener.")
	}
}

func TestAcceptAllowsOtherErrorsThroughBeforeClosed(t *testing.T) {
	taskCounter := sync.WaitGroup{}
	fakeListener := FakeListener{open: false}

	listener := NewGracefulListener(fakeListener, &taskCounter)
	_, err := listener.Accept()
	if err == nil {
		t.Error("Did not receive an error when calling Accept() on an invalid listener.")
	}
}
