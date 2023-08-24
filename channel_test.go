package belajar_go_goroutines

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

//Channel
/*
channel di Go bisa sebagai pengganti async dan synchronus
biasanya ada penerima dan pengirim
penerima dan pengirimnya biasanya goroutine yang berbeda (jadi ada 2 goroutine)
channel hanya bisa mengirim 1 data
channel di Go direpresentasikan dengan tipe data chan
untuk membuat channel bisa menggunakan make(), mirip saat membuat map
namun perlu menentukan tipe data apa yang bisa dimasukan kedalam channel tersebut

channel := make(chan string)

*/

func TestCreateChannel(t *testing.T) {
	channel := make(chan string)
	//jangan lupa close channel
	defer close(channel)

	go func() {
		time.Sleep(2 * time.Second)
		//mengirim data ke channel
		channel <- "Data"
		fmt.Println("Selesai mengirim data ke channel")
	}()

	//menerima data
	data := <-channel
	fmt.Println(data)

	time.Sleep(5 * time.Second)

}

/*
=== RUN   TestCreateChannel
Selesai mengirim data ke channel
Data
--- PASS: TestCreateChannel (7.00s)
PASS
ok      belajar-go-goroutines   10.058s
*/

// Menggunakan channel sebagai parameter, tidak perlu pakai pointer (&)
func GiveMeResponse(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "Aku Bukan Dia"
}

func TestChannelAsParameter(t *testing.T) {
	channel := make(chan string)
	//jangan lupa close channel
	defer close(channel)

	go GiveMeResponse(channel)

	//menerima data
	data := <-channel
	fmt.Println(data)

	time.Sleep(5 * time.Second)
}

/*
=== RUN   TestChannelAsParameter
Aku Bukan Dia
--- PASS: TestChannelAsParameter (7.00s)
PASS
ok      belajar-go-goroutines   7.187s
*/

//Channel In / Out

// hanya bisa mengirim channel
func OnlyIn(channel chan<- string) {
	time.Sleep(2 * time.Second)
	channel <- "Aku Bukan Dia"
}

// hanya bisa menerima channel
func OnlyOut(channel <-chan string) {
	//menerima data
	data := <-channel
	fmt.Println(data)
}

func TestChannelInOut(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go OnlyIn(channel)
	go OnlyOut(channel)

	time.Sleep(5 * time.Second)
}

// Buffered channel, untuk menerima data lebih dari 1
func TestBufferedChannel(t *testing.T) {
	channel := make(chan string, 3)
	defer close(channel)

	go func() {
		channel <- "Satu"
		channel <- "Dua"
		channel <- "Tiga"
	}()

	go func() {
		fmt.Println(<-channel)
		fmt.Println(<-channel)
		fmt.Println(<-channel)

	}()

	time.Sleep(3 * time.Second)
	fmt.Println("Selesai")
}

/*
=== RUN   TestBufferedChannel
Satu
Dua
Tiga
Selesai
--- PASS: TestBufferedChannel (0.00s)
PASS
ok      belajar-go-goroutines   0.467s
*/

// Range channel, perulangan
func TestRangeChannel(t *testing.T) {
	channel := make(chan string)

	go func() {
		for i := 0; i < 10; i++ {
			channel <- "Perulangan ke " + strconv.Itoa(i)
		}
		close(channel)
	}()

	for data := range channel {
		fmt.Println("Menerima data", data)
	}
}

// Select Channel, untuk memilih channel mana yang diterima lebih dahulu
func TestSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("Data dari channel 1", data)
			counter++
		case data := <-channel2:
			fmt.Println("Data dari channel 2", data)
			counter++
		}
		if counter == 2 {
			break
		}
	}
}

// Select Default
func TestSelectDefaultChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("Data dari channel 1", data)
			counter++
		case data := <-channel2:
			fmt.Println("Data dari channel 2", data)
			counter++
		default:
			fmt.Println("Menunggu data")
		}
		if counter == 2 {
			break
		}
	}
}
