package main

import (
	"context"
	"dir/back"
	"dir/front"
)

func main() {

	port := 8080
	tls := false

	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	go func() {
		back.Run(tls, port)
		cancel()
	}()

	go func() {
		front.Run(port)
		cancel()
	}()

	<-ctx.Done()
}
