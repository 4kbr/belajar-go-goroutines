package belajar_go_goroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// /praktek di lapangannya pakai waitgroup bukan time.sleep
func RunAsychronous(group *sync.WaitGroup) {
	defer group.Done()

	group.Add(1)

	fmt.Println("Hello")
	time.Sleep(1 * time.Second)
}

func TestWaitGroup(t *testing.T) {
	group := &sync.WaitGroup{}

	group.Wait()
	fmt.Println("Selesai")
}
