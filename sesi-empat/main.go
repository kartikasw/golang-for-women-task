package main

import (
	"fmt"
	"os"
	"sync"
)

// Buatlah 2 concurrency dengan GOROUTINE dengan 2 data yang berbeda dengan tipe data interface keduanya.
// Lakukan looping kedua GOROUTINE sebanyak 8 kali
// dalam total dimana 4 kali looping interface1 dan 4 kali looping interface2 tersebut lalu tampilkan kedua prosesnya menggunakan Println
// Buatlah 2 kodingan dengan kondisi :
// GOROUTINE keduanya menampilkan secara acak
// GOROUTINE keduanya menampilkan secara rapih

func printData(data interface{}, index int) {
	fmt.Println(data, index)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Pilih random atau arranged.")
		return
	}

	switch arg := os.Args[1]; arg {
	case "arranged":
		runArraged()
	case "random":
		runRandom()
	default:
		fmt.Println("Pilih random atau arranged.")
	}
}

func runArraged() {
	ch1 := make(chan interface{})
	ch2 := make(chan interface{})

	for i := 1; i < 5; i++ {
		go func() {
			data := []string{"coba1", "coba2", "coba3"}
			ch1 <- data
		}()

		go func(index int) {
			result1 := <-ch1
			printData(result1, index)
			data := []string{"bisa1", "bisa2", "bisa3"}
			ch2 <- data
		}(i)

		result2 := <-ch2
		printData(result2, i)
	}
}

func runRandom() {
	var wg sync.WaitGroup

	for i := 1; i < 5; i++ {
		wg.Add(2)

		go func(index int) {
			defer wg.Done()
			data := []string{"coba1", "coba2", "coba3"}
			printData(data, index)
		}(i)

		go func(index int) {
			defer wg.Done()
			data := []string{"bisa1", "bisa2", "bisa3"}
			printData(data, index)
		}(i)
	}

	wg.Wait()
}
