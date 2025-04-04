package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

// func TestRacer(t *testing.T) {
// 	slowURL := "http://www.facebook.com"
// 	fastURL := "http://www.quii.dev"

// 	want := fastURL
// 	got := Racer(slowURL, fastURL)

// 	if got != want {
// 		t.Errorf("got %q want %q", got, want)
// 	}
// }

// net/http/httptest allows us to easily create a mock HTTP server
// func TestRacer_v2(t *testing.T) {

// 	// This is how you would write a real HTTP server in Go, however
// 	// we wrap it in httptest which makes it easier to use with testing
// 	slowServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter,
// 		r *http.Request) {
// 		time.Sleep(20 * time.Millisecond)
// 		w.WriteHeader(http.StatusOK)
// 	}))

// 	fastServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter,
// 		r *http.Request) {
// 		w.WriteHeader(http.StatusOK)
// 	}))

// 	slowURL := slowServer.URL
// 	fastURL := fastServer.URL

// 	want := fastURL
// 	got := Racer(slowURL, fastURL)

// 	if got != want {
// 		t.Errorf("got %q, want %q", got, want)
// 	}

// 	slowServer.Close()
// 	fastServer.Close()
// }

func TestRacer(t *testing.T) {

	t.Run("compares speed of servers base", func(t *testing.T) {
		slowServer := makeDelayedServer(50 * time.Millisecond)
		fastServer := makeDelayedServer(0 * time.Millisecond)

		// By prefixing a function call with defer it will now call that function
		// AT THE END OF THE CONTAINING FUNCTION

		// You want this to execute at the end of the function, but keep the instruction
		// near where you created the server for the benefit of future readers
		defer slowServer.Close()
		defer fastServer.Close()

		slowURL := slowServer.URL
		fastURL := fastServer.URL

		got, _ := Racer(slowURL, fastURL)
		want := fastURL

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}

	})

	t.Run("returns an error if a server doesn't respond within 10s",
		func(t *testing.T) {

			serverA := makeDelayedServer(11 * time.Second)
			serverB := makeDelayedServer(12 * time.Second)

			defer serverA.Close()
			defer serverB.Close()

			_, err := Racer(serverA.URL, serverB.URL)

			if err == nil {
				t.Error("expected an error but didn't get one")
			}
		})

}

// Refactor our test case
func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter,
		r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
