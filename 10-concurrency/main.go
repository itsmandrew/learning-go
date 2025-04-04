package concurrency

type WebsiteChecker func(string) bool

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := map[string]bool{}

	for _, url := range urls {
		results[url] = wc(url)
	}
	return results
}

// To start a new goroutine we turn a function call into a go statement
// by putting the keyword go in front of it: go doSomething()

// func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {

// 	results := map[string]bool{}

// 	// Each iteration of the loop will start a new goroutine, concurrent
// 	// with the current process. Each goroutine will add its result to the map
// 	for _, url := range urls {
// 		go func() {
// 			results[url] = wc(url)
// 		}()
// 	}

// 	time.Sleep(2 * time.Second)
// 	// -> sometimes, when we run our tests, two of the goroutines write to
// 	// the results map at exactly the same time. maps in Go don't like it when more
// 	// than one thing tries to write to them at once, thus fatal error

// 	// THIS IS CALLED A RACE CONDITION, bug that occurs when output of software
// 	// is dependent on the timing and sequence of events that we have no control
// 	//over, We cannot control exactly when each goroutine writes to the results
// 	// map, we are vulnerablee to two goroutines writing to it at the same time
// 	return results

// }

type result struct {
	string
	bool
}

// We can solve this race condition by coordinating goroutines using channels.
// Channels are a Go data structure that can both receive and send values.
// These operations, along w/ their details, allow communication between
// different processes.
func CheckWebsites_v2(wc WebsiteChecker, urls []string) map[string]bool {
	results := map[string]bool{}
	resultChannel := make(chan result) // a channel of result

	// Iterate over the urls, instead of writing to the map directly we're
	// sending a result struct for each call to wc to the resultChannel w/
	// a send statement. This uses the <- operator, taking a channel on the left
	// and a value on the right
	for _, url := range urls {
		go func() {
			resultChannel <- result{url, wc(url)}
		}()
	}
	// this for loop interates once for each of the urls.
	// we use a receive expression, which assigns a value received from a channel
	// to a variable. also uses the <- operator but reversed
	for i := 0; i < len(urls); i++ {
		r := <-resultChannel
		results[r.string] = r.bool
	}
	return results
}

// By sending the results into a channel, we can control the timing of each
// write into the results map, ensuring that it happens one at a time.
// Although each of the calls of wc, and each send to the result channel, is
// happening concurrently inside its own process, each of the results is being
// dealt w/ onee at a time as we take values out of the result channel.

// We have used concurrency for the part of the code that we wanted to make faster, while
// making sure that the part that cannot happen simultaneuously still happens linearly.
// We have communicated across the multiple processes involved by using channels
