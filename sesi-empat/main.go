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
	var wg sync.WaitGroup

	switch arg := os.Args[1]; arg {
	case "arranged":
		runArraged(&wg)
	case "random":
		runRandom(&wg)
	default:
		fmt.Println("Pilih random atau arranged.")
	}

	wg.Wait()
}

func runArraged(wg *sync.WaitGroup) {
	ch1 := make(chan interface{})
	ch2 := make(chan interface{})

	for i := 1; i < 5; i++ {
		wg.Add(2)

		go func() {
			defer wg.Done()
			data := []string{"coba1", "coba2", "coba3"}
			ch1 <- data
		}()

		go func(index int) {
			defer wg.Done()
			result1 := <-ch1
			printData(result1, index)
			data := []string{"bisa1", "bisa2", "bisa3"}
			ch2 <- data
		}(i)

		result2 := <-ch2
		printData(result2, i)
	}
}

func runRandom(wg *sync.WaitGroup) {
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
}
