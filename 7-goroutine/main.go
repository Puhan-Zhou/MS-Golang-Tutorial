package main

import (
	"fmt"
	"net/http"
	"time"
)

func process(ch chan string) {
	time.Sleep(3 * time.Second)
	ch <- "Done processing!"
}

func replicate(ch chan string) {
	time.Sleep(1 * time.Second)
	ch <- "Done replicating!"
}

func main() {
	// show the usage of select
	ch1 := make(chan string)
	ch2 := make(chan string)
	go process(ch1)
	go replicate(ch2)

	for i := 0; i < 2; i++ {
		// `select` is a switch-case like keyword only for channel.
		// you can run different branch in select when you receive something
		// from a specific channel.
		// if select receive data from a channel in any branch,
		// the select block will be quit.
		// so if you delete the for-loop here, this select will only print the data
		// that receive firstly, in this case is `Done replicating!`
		select {
		case process := <-ch1:
			fmt.Println(process)
		case replicate := <-ch2:
			fmt.Println(replicate)
		}
	}

	// Non-buffer channel
	// ch := make(chan string)
	// Buffered channel
	ch := make(chan string, 10)
	start := time.Now()

	apis := []string{
		"https://management.azure.com",
		"https://dev.azure.com",
		"https://api.github.com",
		"https://outlook.office.com/",
		"https://api.somewhereintheinternet.com/",
		"https://graph.microsoft.com",
	}

	for _, api := range apis {
		go CheckApi(api, ch)
	}

	for i := 0; i < len(apis); i++ {
		// same as fmt.Println(<-ch)
		// write like this is easy to read
		// and can limit the action of ch
		fmt.Println(receive(ch))
	}

	elapsed := time.Since(start)
	fmt.Printf("Done! It took %v seconds!\n", elapsed.Seconds())

}

// specify the chan is send-only
func send(ch chan<- string, data string) {
	ch <- data
}

// specify the chan is receive-only
func receive(ch <-chan string) string {
	return <-ch
}

func CheckApi(api string, ch chan string) {
	_, err := http.Get(api)
	if err != nil {
		// same as `ch <- fmt.Sprintf("ERROR: %s is down", api)`
		// write like this is easy to read
		// and can limit the action of ch
		send(ch, fmt.Sprintf("ERROR: %s is down", api))
		return
	}

	send(ch, fmt.Sprintf("SUCCEED: %s is up and running", api))
}
