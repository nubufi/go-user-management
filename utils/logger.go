package utils

import (
	"log"
	"sync"
	"time"
)

var once sync.Once

// Initialize a global logger
func InitLogger() {
	once.Do(func() {
		log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	})
}

// LogRequestDuration logs the duration of a request to the console.
// It takes a wait group, a channel to send the duration, and the route name as arguments.
// It logs the route name and the duration of the request.
// It sends the duration to the channel and decrements the wait group when done.
func LogRequestDuration(wg *sync.WaitGroup, ch chan time.Duration, route string) {
	defer wg.Done()
	start := time.Now()

	duration := time.Since(start)
	ch <- duration
	log.Printf("Route: %s | Request duration: %v\n", route, duration)
}
