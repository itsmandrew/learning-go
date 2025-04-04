package concurrency

import (
	"testing"
	"time"
)

func slowStubWebsiteCheck(_ string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}
func BenchmarkCheckWebsites(b *testing.B) {

	urls := make([]string, 100)
	for i := 0; i < len(urls); i++ {
		urls[i] = "a url"
	}
	b.ResetTimer()

	// The benchmark tests CheckWebsites using a slice of one hundred urls
	// and uses a new fake implementation of WebsiteChecker.
	// SlowStubWebsiteChecker is deliberately slow. It uses time.Sleep to wait
	// exactly twenty ms and then it returns true.

	// We use b.ResetTimer(0 so that we reset the time of our test before it runs
	for i := 0; i < b.N; i++ {
		CheckWebsites(slowStubWebsiteCheck, urls)
	}
}

func BenchmarkCheckWebsites_v2(b *testing.B) {

	urls := make([]string, 100)
	for i := 0; i < len(urls); i++ {
		urls[i] = "a url"
	}
	b.ResetTimer()

	// The benchmark tests CheckWebsites using a slice of one hundred urls
	// and uses a new fake implementation of WebsiteChecker.
	// SlowStubWebsiteChecker is deliberately slow. It uses time.Sleep to wait
	// exactly twenty ms and then it returns true.

	// We use b.ResetTimer(0 so that we reset the time of our test before it runs
	for i := 0; i < b.N; i++ {
		CheckWebsites_v2(slowStubWebsiteCheck, urls)
	}
}
