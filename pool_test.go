package belajar_go_goroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

///Pool adalah design pattern pool untuk menyimpan data di GO

func TestPool(t *testing.T) {
	pool := sync.Pool{
		New: func() interface{} {
			return "Kosong"
		},
	}

	pool.Put("Aku")
	pool.Put("kamu")
	pool.Put("Dia")

	for i := 0; i < 10; i++ {
		go func() {
			data := pool.Get()
			fmt.Println("data get:", data)
			time.Sleep(1 * time.Second)
			pool.Put(data)
			fmt.Println("data after put:", data)
		}()
	}

	time.Sleep(11 * time.Second)
	fmt.Println("selesai")

}
