package manners

import (
	"bytes"
	"errors"
	"net"
	"net/http"
	"time"
)

type FakeHandler struct{}

func (this FakeHandler) ServeHTTP(response http.ResponseWriter, request *http.Request) {}

func NewResponseWriter() FakeResponseWriter {
	return FakeResponseWriter{FakeHeader: make(map[string][]string), Content: bytes.NewBuffer([]byte("content"))}
}

type FakeResponseWriter struct {
	FakeHeader http.Header
	Content    *bytes.Buffer
}

func (this *FakeResponseWriter) Header() http.Header {
	return this.FakeHeader
}

func (this *FakeResponseWriter) Write(content []byte) (int, error) { return 0, nil }
func (this *FakeResponseWriter) WriteHeader(status int)            {}

type FakeAddr struct{}

func (this FakeAddr) Network() string { return "" }
func (this FakeAddr) String() string  { return "" }

type FakeConnection struct{}

func (this FakeConnection) Read(b []byte) (n int, err error)   { return 0, nil }
func (this FakeConnection) Write(b []byte) (n int, err error)  { return 0, nil }
func (this FakeConnection) Close() error                       { return nil }
func (this FakeConnection) LocalAddr() net.Addr                { return FakeAddr{} }
func (this FakeConnection) RemoteAddr() net.Addr               { return FakeAddr{} }
func (this FakeConnection) SetDeadline(t time.Time) error      { return nil }
func (this FakeConnection) SetReadDeadline(t time.Time) error  { return nil }
func (this FakeConnection) SetWriteDeadline(t time.Time) error { return nil }

type FakeListener struct {
	open bool
}

func (this FakeListener) Accept() (c net.Conn, err error) {
	if this.open {
		return FakeConnection{}, nil
	}
	return nil, errors.New("Accept called on closed FakeListener")
}

func (this FakeListener) Close() error {
	this.open = false
	return nil
}

func (this FakeListener) Addr() net.Addr { return FakeAddr{} }
