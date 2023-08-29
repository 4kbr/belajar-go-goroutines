package belajar_go_goroutines

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

func TestAtomic(t *testing.T) {
	var x int64 = 0
	group := sync.WaitGroup{}

	for i := 1; i <= 1000; i++ {
		group.Add(1)
		go func() {
			for j := 1; j <= 10; j++ {
				x = x + 1
				atomic.AddInt64(&x, 1)
			}
			group.Done()
		}()
		// time.Sleep(5 * time.Second)
		group.Wait()
		fmt.Println("Counter =", x)
	}
}


