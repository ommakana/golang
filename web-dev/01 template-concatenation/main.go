package main

import "fmt"

func main() {
	name := "omi"
	html := `<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<title>Go lang</title>
	</head>
	<body>
		<h1> Hi ` + name + `</h1>
		</body>
		</html>`
	fmt.Println(html)
}
