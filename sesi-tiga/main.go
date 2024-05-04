package main

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
)

// Buat logic untuk menampilkan/generate data peserta dari cli by name, contoh cli : go run main.go Fitri
// Buatlah looping yang memproses slice of string dari beberapa nama peserta,
// dimana jika nama yg di ketik pada cli sama dengan yang ada di data peserta
// maka akan dapat index dari data pesertanya dan di append pada data baru untuk ditampilkan pada CLI atau terminal.
// Jika data nama atau absen yang diinput pada CLI tidak tersedia pada program,
// misal go run main.go Fitri atau go run main.go 1 maka bisa dihandle dengan pesan "Data dengan nama/absen tsb tidak tersedia".

type Friend struct {
	Nama      string
	Alamat    string
	Pekerjaan string
	Alasan    string
}

var friends = []Friend{
	{"Agresia", "Jl. Sudirman No. 1", "Developer", "Tertarik dengan performa Golang"},
	{"Annisa", "Jl. Setiabudi No. 2", "Developer", "Ingin mempelajari Golang"},
	{"Dwi", "Jl. Karet No. 3", "Data Scientist", "Ingin switching karir"},
	{"Farida", "Jl. Kuningan No. 4", "Software Engineer", "Ingin migrasi ke Golang"},
	{"Najya", "Jl. Matraman No. 5", "DevOps Engineer", "Biar keren"},
	{"Rahma", "Jl. BNI No. 6", "Game Developer", "Hype"},
	{"Rika", "Jl. Mandiri No. 7", "Backend Engineer", "Bahasanya relatif mudah"},
}

func findFriendByName(name string) (Friend, int) {
	for index, friend := range friends {
		if strings.EqualFold(friend.Nama, name) {
			return friend, index
		}
	}
	return Friend{}, -1
}

func findFriendById(id int) (Friend, int) {
	if id >= 0 && id < len(friends) {
		return friends[id], id
	}

	return Friend{}, -1
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Tolong masukkan nama atau nomor absen\nContoh: 'go run main.go Fitri' atau 'go run main.go 2'")
		return
	}

	arg := os.Args[1]

	var friend Friend
	var order int

	if input, err := strconv.Atoi(arg); err == nil {
		friend, order = findFriendById(input)
	} else {
		friend, order = findFriendByName(arg)
	}

	result1 := reflect.Indirect(reflect.ValueOf(friend))
	result2 := reflect.ValueOf(friend)

	if order != -1 {
		fmt.Println("ID: ", order)

		for i := 0; i < result2.NumField(); i++ {
			field := result1.Type().Field(i).Name
			value := reflect.ValueOf(friend).Field(i)

			fmt.Printf("%s: %s\n", field, value)
		}
	} else {
		fmt.Println("Data dengan nama/absen tsb tidak tersedia")
	}
}
