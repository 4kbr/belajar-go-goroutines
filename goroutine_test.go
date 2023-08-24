package belajar_go_goroutines

import (
	"fmt"
	"testing"
	"time"
)

func RunHelloWorld() {
	fmt.Println("Hello word")
}

func TestCreateGoRoutine(t *testing.T) {
	//tambahkan syntax go untuk menjalankan function dengan goroutine
	go RunHelloWorld()
	fmt.Println("Ini printl")
	time.Sleep(1 * time.Second)
}

func DisplayNumber(number int) {
	fmt.Println("Display", number)
}

func TestManyGoroutine(t *testing.T) {
	for i := 0; i < 100000; i++ {
		go DisplayNumber(i)
	}

	time.Sleep(10 * time.Second)
}

//tanpa go routine
/*
--- PASS: TestManyGoroutine (14.89s)
PASS
ok      belajar-go-goroutines   15.647s
*/

