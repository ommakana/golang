package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{"https://www.google.com", "https://www.stackoverflow.com", "https://ommakana.w3spaces.com/"}
	c := make(chan string)
	for _, link := range links {
		go checkLinkStatus(link, c)
	}
}

func checkLinkStatus(l string, c chan string) {
	_, err := http.Get(l)
	if err != nil {
		fmt.Println(l, "Link not up!")
		c <- "Link is down"
	}
	c <- (l + " Link is up and running")
}
