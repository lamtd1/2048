package swipe

import "test/rotate"

func Right(arr *[4][4]int) {
	for i := range 4 {
		newRow := make([]int, 0)
		for j := 3; j >= 0; j-- {
			if arr[i][j] != 0 {
				newRow = append(newRow, arr[i][j])
			}
		}
		merged := make([]int, 0)
		skip := false
		for j := 0; j < len(newRow); j++ {
			if skip {
				skip = false
				continue
			}
			if j+1 < len(newRow) && newRow[j] == newRow[j+1] {
				merged = append(merged, newRow[j]*2)
				skip = true
			} else {
				merged = append(merged, newRow[j])
			}
		}
		k := 0
		for j := 3; j >= 0; j-- {
			if k < len(merged) {
				arr[i][j] = merged[k]
				k++
			} else {
				arr[i][j] = 0
			}
		}
	}
}

func Left(arr *[4][4]int) {
	var tmp = rotate.Reverse(arr)
	Right(&tmp)
	*arr = rotate.Reverse(&tmp)

}

func Up(arr *[4][4]int) {
	var tmp = rotate.TurnLeft(arr)
	Right(&tmp)
	*arr = rotate.TurnLeft(&tmp)
}

func Down(arr *[4][4]int) {
	var tmp = rotate.TurnRight(arr)
	Right(&tmp)
	*arr = rotate.TurnRight(&tmp)
}
