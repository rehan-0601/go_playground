package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	// channel for communicating 'string' data type
	// will be used for the checkLink subroutine to communicate with main go routine
	c := make(chan string)

	// concurrency is not the same as parallelism
	// for a single core, concurrency: ability to context switch between sub routines when
	// one is blocked.
	// parallelism is possible on a multi core cpu
	for _, link := range links {
		go checkLink(link, c)
	}

	// print is receiving the message in the channel. once it receives the first mesasge
	// it will quit the main go routine
	// receiving from channel is a BLOCKING call
	// fmt.Println(<-c)

	// for i := 0; i < len(links); i++ {
	// 	// each time we get smhng in the channel. we want to send out another request
	// 	go checkLink(<-c, c)
	// }

	// infinite for. NOTE: <- is blocking call
	// for {
	// 	// note first argument is interpreted as string, voila!
	// 	go checkLink(<-c, c)
	// }

	// following code is refactor of above loop
	for l := range c {
		// haivng a sleep here is not a good idea. main routine will sleep and might lose
		// messages that came earlier than 5 secs
		// time.Sleep(5 * time.Second)

		// function literal below
		go func(link string) {
			time.Sleep(5 * time.Second)

			// never share variable between 2 subroutines
			// l in outerscope might change by the time checkLink starts execution
			// instead of reference to outer scope, pass it as an argument. since its pass by value,
			// it wont be affecting by outer l, and u will be good in checkLink execution
			checkLink(link, c)
		}(l)
	}

}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down")
		c <- link
		return
	}
	fmt.Println(link, "is up")
	c <- link
}
