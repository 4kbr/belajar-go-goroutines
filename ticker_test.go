package belajar_go_goroutines

import (
	"fmt"
	"testing"
	"time"
)

func TestTicker(t *testing.T) {
	ticker := time.NewTicker(1 * time.Second)
	counter := 0
	for time := range ticker.C {
		if counter == 5 {
			ticker.Stop()
		}
		fmt.Println(time, counter)
		counter++
	}
	fmt.Println("selesai")
}

func TestTick(t *testing.T) {
	channel := time.Tick(1 * time.Second)

	for time := range channel {
		fmt.Println(time)
	}
	fmt.Println("selesai")
}
  