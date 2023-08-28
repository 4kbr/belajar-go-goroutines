package belajar_go_goroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestAtomic(t *testing.T) {
	var x int64 = 0
	group := sync.WaitGroup{}

	for i := 1; i <= 1000; i++ {
		go func() {
			defer group.Done()
			group.Add(1)
			for j := 1; j <= 100; j++ { 0 
				x = x + 1
			}
		}()

		time.Sleep(5 * time.Second)
		fmt.Println("Counter =", x)
	}

}
