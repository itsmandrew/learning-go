package main

import (
	"fmt"
	"net/http"
	"time"
)

// func Racer(a, b string) (winner string) {

// 	aTime := measureResponseTime(a)
// 	bTime := measureResponseTime(b)

// 	if aTime < bTime {
// 		return a
// 	}
// 	return b
// }

func Racer(a, b string) (winner string, error error) {
	// Select allows you to wait on multiple channels. The first
	// one to send a value "wins" and the code underneath the case is executed

	// We use ping in our select to set up two channels, one for each of our URLS
	// whichever one writees to its channel first will have its code executed
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(10 * time.Second):
		return "", fmt.Errorf("timed out waiting for %s and %s", a, b)
	}
}

// We don't care what type is sent to the channel, wee just want to signal
// we are done and clsoing the channel

// Inside, we start a goroutine which will send a signal into that channel
// once we have completed http.Get(url)
func ping(url string) chan struct{} {
	ch := make(chan struct{}) // Always use make when creatingg a channel
	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}

// Wee don't care about the exact response times of the request, we just want
// to know which onee comes back first

// func measureResponseTime(url string) time.Duration {
// 	start := time.Now()
// 	http.Get(url)

// 	return time.Since(start)
// }
