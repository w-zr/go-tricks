package stopper

import (
	"context"
	"testing"
	"time"
)

func TestStopper(t *testing.T) {
	s := NewStopper()
	i := 0
	s.Run(func(ctx context.Context) {
		defer t.Log("stopped")
		for {
			select {
			case <-ctx.Done():
				return
			default:
			}
			t.Log("counting...", i)
			i++
			time.Sleep(time.Second)
		}
	})
	time.Sleep(5 * time.Second)
	s.Stop()
}
