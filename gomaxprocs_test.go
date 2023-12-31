package belajar_go_goroutines

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

func TestGetGomaxprocs(t *testing.T) {
	var group = sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		group.Add(1)
		go func() {
			time.Sleep(3 * time.Second)
			group.Done()
		}()
	}

	totalCPU := runtime.NumCPU()
	fmt.Println("Total CPU:", totalCPU)
	totalThread := runtime.GOMAXPROCS(-1)
	fmt.Println("Total CPU:", totalThread)
	totalGoroutine := runtime.NumGoroutine()
	fmt.Println("Total CPU:", totalGoroutine)

	fmt.Println("selesai")
	group.Wait()
}

func TestChangeThreadNumber(t *testing.T) {
	var group = sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		group.Add(1)
		go func() {
			time.Sleep(3 * time.Second)
			group.Done()
		}()
	}

	totalCPU := runtime.NumCPU()
	fmt.Println("Total CPU:", totalCPU)

	runtime.GOMAXPROCS(20)
	totalThread := runtime.GOMAXPROCS(-1)
	fmt.Println("Total CPU:", totalThread)

	totalGoroutine := runtime.NumGoroutine()
	fmt.Println("Total CPU:", totalGoroutine)

	fmt.Println("selesai")
	group.Wait()
}
