package belajar_go_goroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func RunAsychronous(group *sync.WaitGroup) {
	defer group.Done()

	group.Add(1)

	fmt.Println("Hello")
	time.Sleep(1 * time.Second) 
}

func TestWaitGroup(t *testing.T) {
	group := &sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		go RunAsychronous(group)
	}

	group.Wait()
	fmt.Println("Selesai")
}
 