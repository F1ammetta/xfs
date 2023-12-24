package main

import (
	// "context"
	"dir/back"
	"dir/front"
)

func main() {

	port := 8080

	go back.Run(false, port)

	front.Run(port)

}
