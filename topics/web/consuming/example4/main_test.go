// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Tests to show how to consume a web API using the default http
// support in the standard library. This shows a Client timeout.
package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

// App returns a handler takes a while to respond.
func App() http.Handler {

	// Handler function will be used for mocking. It waits
	// 500 milliseconds before responding.
	h := func(res http.ResponseWriter, req *http.Request) {
		time.Sleep(500 * time.Millisecond)
	}

	// Return the handler function.
	return http.HandlerFunc(h)
}

func TestApp(t *testing.T) {

	// Start a server to handle these requests.
	ts := httptest.NewServer(App())
	defer ts.Close()

	// Create a new request for the GET call.
	req, err := http.NewRequest("GET", ts.URL, nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a Client value with a timeout.
	client := http.Client{
		Timeout: 50 * time.Millisecond,
	}

	// Perform the GET call with the excepted timeout. We are expecting an
	// error here so we fail if err == nil.
	if _, err := client.Do(req); err == nil {
		t.Fatal("request was supposed to timeout")
	}
}
