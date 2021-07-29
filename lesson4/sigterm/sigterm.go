package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var done = make(chan bool, 1)

func main() {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancelFunc()

	sigs := make(chan os.Signal, 1)

	go func(ctx context.Context) {
		time.Sleep(10 * time.Second)
		done <- true
	}(ctx)
	signal.Notify(sigs, syscall.SIGINT)
	go func(ctx context.Context) {
		time.Sleep(1 * time.Second)
		//signal.Notify(sigs, syscall.SIGINT)

	}(ctx)

	select {
	case <-ctx.Done():
		err := ctx.Err()
		fmt.Println(err)
	case sig := <-sigs:
		fmt.Println(sig)
	case <-done:
		fmt.Println("finish")
	}
}
