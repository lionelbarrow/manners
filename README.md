# Manners

A *polite* webserver for Go.

Manners allows you to shut your Go webserver down gracefully, without dropping any requests. It can act as a drop-in replacement for the standard library's http.ListenAndServe function:

```go
func main() {
  handler := MyHTTPHandler()
  server := manners.NewServer()
  server.ListenAndServe(":7000", handler)
}
```

Then, when you want to shut the server down:

```go
server.Shutdown <- true
```

(Note that this does not block until all the requests are finished. Rather, the call to server.ListenAndServe will stop blocking when all the requests are finished.)

Manners ensures that all requests are served by incrementing a WaitGroup when a request comes in and decrementing it when the request finishes.

If your request handler spawns Goroutines that are not guaranteed to finish with the request, you can ensure they are also completed with the `StartRoutine` and `FinishRoutine` functions on the server.

### Compatability

Manners 0.3.0 and above use standard library functionality introduced in Go 1.3.

### Installation

`go get github.com/braintree/manners`
