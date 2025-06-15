package rotate

func Reverse(arr *[4][4]int) [4][4]int {
	var tmp [4][4]int
	for i := range 4 {
		for j := range 4 {
			tmp[i][j] = arr[3-i][3-j]
		}
	}
	return tmp
}
func TurnLeft(arr *[4][4]int) [4][4]int {
	var tmp [4][4]int
	for i := range 4 {
		for j := range 4 {
			tmp[i][j] = arr[3-j][3-i]
		}
	}
	return tmp
}
func TurnRight(arr *[4][4]int) [4][4]int {
	var tmp = Reverse(arr)
	tmp = TurnLeft(&tmp)
	return tmp

}
