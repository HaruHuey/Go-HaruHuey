package main

import (
	"fmt"
)

func main() {
	fmt.Printf("Hello, Go!\nGo!")

	// int / 8 / 16 / 32 / 64
	intdata := 7

	// float / 32 / 64
	var floatdata float32 = 2.8

	stringdata := "HaruHuey Go"

	// int 형 리스트
	intlist := []int{57, 23, 15}

	stringlist := []string{
		"HaruHuey",
		"TypeVariable",
		"Go lang",
	}

	fmt.Println(intdata)
	fmt.Println(floatdata)
	fmt.Println(stringdata)
	fmt.Println(intlist[1])
	fmt.Println(intlist[0], intlist[1], intlist[2])
	fmt.Println(stringlist[0], stringlist[1], stringlist[2])
}
