package main

import (
	"fmt"

	"test/swipe"
	check "test/utils"
)

func main() {
	arr := [4][4]int{
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	}

	for range 3 {
		check.RandGen(&arr)
	}
	var opt string
	for {
		check.Display(&arr)
		fmt.Println("D: RIGHT")
		fmt.Println("A: LEFT")
		fmt.Println("W: UP")
		fmt.Println("S: DOWN")
		fmt.Println("0: EXIT")
		fmt.Println("Nhap nuoc di: ")
		fmt.Scan(&opt)

		if opt == "0" {
			break
		} else if opt == "d" {
			swipe.Right(&arr)
		} else if opt == "a" {
			swipe.Left(&arr)
		} else if opt == "w" {
			swipe.Up(&arr)
		} else if opt == "s" {
			swipe.Down(&arr)
		}
		if check.IsOver(&arr) {
			fmt.Println("WINNER")
		}
		check.RandGen(&arr)
	}

}
