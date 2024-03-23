package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{"https://www.google.com", "https://www.stackoverflow.com", "https://ommakana.w3spaces.com/"}
	c := make(chan string)
	for _, link := range links {
		go checkLinkStatus(link, c)
	}

	for l := range c {
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLinkStatus(link, c)
		}(l)
	}
}

func checkLinkStatus(l string, c chan string) {
	_, err := http.Get(l)
	if err != nil {
		fmt.Println(l, "Link not up!")
		c <- l
	}
	fmt.Println(l, "Link is up and running!")
	c <- l
}
