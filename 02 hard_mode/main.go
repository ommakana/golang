package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	args := os.Args
	fmt.Println(args[1])
	// bs, err := os.ReadFile(args[1])
	// if err != nil {
	// 	fmt.Println("Error reading from file")
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(string(bs))
	// }
	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Error reading file")
		os.Exit(1)
	}
	io.Copy(os.Stdout, f)
}
