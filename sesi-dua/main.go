package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Buatlah looping dengan variable yang berisi string suatu kalimat dan pecahlah kalimat tersebut menjadi 1 huruf per 1 kata

func main() {
	fmt.Print("Masukkan kalimat = ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	fmt.Println()

	for _, letter := range input {
		if letter != ' ' {
			fmt.Println(string(letter))
		} else {
			fmt.Println()
		}
	}

	result := map[string]int{}
	for _, char := range input {
		result[string(char)]++
	}

	fmt.Println(result)
}
