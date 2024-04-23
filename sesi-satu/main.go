package main

import "fmt"

// Buatlah looping berkondisi dengan variable n sebagai penentu maksimal number loopingnya,
// Setiap looping dengan angka kelipatan 3 maka diganti dengan kata “Fizz”
// Setiap looping dengan angka kelipatan 5 maka diganti dengan kata “Buzz”
// Setiap looping dengan angka kelipatan 3 dan 5 maka diganti dengan kata “FizzBuzz”

func main() {
	var input int16

	fmt.Print("Masukkan batas loop = ")
	fmt.Scanln(&input)
	fmt.Println()

	for i := 1; i <= int(input); i++ {
		switch {
		case i%3 == 0 && i%5 == 0:
			fmt.Println("FizzBuzz")
		case i%3 == 0:
			fmt.Println("Fizz")
		case i%5 == 0:
			fmt.Println("Buzz")
		default:
			fmt.Println(i)
		}

	}
}
