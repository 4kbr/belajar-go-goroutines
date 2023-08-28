package belajar_go_goroutines

import (
	"fmt"
	"sync"
	"testing"
)

///once, hanya sekali dipanggil

var counter int = 0

func OnlyOnce() {
	counter++
}

func TestOnce(t *testing.T) {
	once := sync.Once{}

	group := sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		go func() {
			defer group.Done()
			group.Add(1)
			once.Do(OnlyOnce) //ini cuma kepanggil sekali
			// OnlyOnce() //yan ini berkali-kali, counternya jadi lebih dari 1
		}()
	}

	group.Wait()
	fmt.Println("Counter", counter) //1
}
