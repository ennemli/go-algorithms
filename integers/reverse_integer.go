package integers

func reverse(x int) int {
	xB, yB := -2147483648, 2147483647
	reversed_int := 0
	for x != 0 {
		digit := x % 10
		reversed_int = reversed_int*10 + digit
		x = int(x / 10)
	}
	if reversed_int <= yB && reversed_int >= xB {
		return reversed_int
	} else {
		return 0
	}
}
