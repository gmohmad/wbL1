package main

import (
	"context"
	"fmt"
	"time"
)

func main() {

	// done channel
	// Launch a goroutine that print "waiting on DONE CHANNEL" while waiting one done channel
	done := make(chan struct{})

	go func() {
		for {
			select {
			case <-done:
				return
			default:
				fmt.Println("waiting on DONE CHANNEL")
			}
			time.Sleep(time.Millisecond * 100)
		}
	}()

	// Timer channel
	// Launch a goroutine that print "waiting on TIMER" while waiting one time.Tick channel
	go func() {
		for {
			select {
			case <-time.Tick(time.Second * 5):
				return
			default:
				fmt.Println("waiting on TIMER")
			}
			time.Sleep(time.Millisecond * 100)
		}
	}()

	// context
	ctx, cancel := context.WithCancel(context.Background()) // We could use WithDeadline or WithTimeout too

	// Launch a goroutine that print "waiting on CONTEXT" while waiting one ctx.Done() channel
	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				fmt.Println("waiting on CONTEXT")
			}
			time.Sleep(time.Millisecond * 100)
		}
	}(ctx)

	time.Sleep(time.Second)
	cancel()
	close(done)

	// etc
}
