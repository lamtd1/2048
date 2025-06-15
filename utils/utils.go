package check

import (
	"fmt"
	"math/rand"
)

func IsOver(arr *[4][4]int) bool {
	for i := range 4 {
		for j := range 4 {
			if arr[i][j] == 2048 {
				return true
			}
		}
	}
	return false
}
func Display(arr *[4][4]int) {
	border := "+-----+-----+-----+-----+"
	for i := 0; i < 4; i++ {
		fmt.Println(border)
		for j := 0; j < 4; j++ {
			if arr[i][j] == 0 {
				fmt.Printf("|%5s", "   ")
			} else {
				fmt.Printf("|%5d", arr[i][j])
			}
		}
		fmt.Println("|")
	}
	fmt.Println(border)
}

func RandGen(arr *[4][4]int) {
	for {
		i := rand.Intn(4)
		j := rand.Intn(4)

		if arr[i][j] == 0 {
			if rand.Intn(2) == 0 {
				arr[i][j] = 4
			} else {
				arr[i][j] = 2
			}
			break
		} else {
			continue
		}
	}

}
