package main

import (
	"log"
	"net/http"
)

// A Client sends http Request and returns response and error
type Client interface {
	Do(*http.Request) (*http.Response, error)
}

// ClientFunc is a custom type for the implementation of the
// Client interface
type ClientFunc func(*http.Request) (*http.Response, error)

// Do implements the Client interface for the ClientFunc
func (f ClientFunc) Do(r *http.Request) (*http.Response, error) {
	return f(r)
}

// Decorator wraps a Client with extra behaviour
type Decorator func(Client) Client

// Logging returns a Decorator that logs a Client's requests.
func Logging(l *log.Logger) Decorator {
	return func(c Client) Client {
		return ClientFunc(func(r *http.Request) (*http.Response, error) {
			l.Printf("%s: %s %s", r.UserAgent(), r.Method, r.URL)
			return c.Do(r)
		})
	}
}

// Decorate decorates a Client c with all the given Decorators, in order.
func Decorate(c Client, ds ...Decorator) Client {
	decorated := c
	for _, decorate := range ds {
		decorated = decorate(decorated)
	}
	return decorated
}
