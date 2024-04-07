package main

import (
	"context"
	"dir/back"
	"dir/front"
	"os"
)

func main() {

	var port = 8080

	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	go func() {
		back.Run(port)
		cancel()
	}()

	if os.Getenv("FRONT") == "true" {
		go func() {
			front.Run(port)
			cancel()
		}()
	}

	<-ctx.Done()
}
