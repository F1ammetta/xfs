package main

import (
	"context"
	"dir/back"
	"dir/front"
)

func main() {

	port := 8080

	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	go func() {
		back.Run(false, port)
		cancel()
	}()

	go func() {
		front.Run(port)
		cancel()
	}()

	<-ctx.Done()
}
