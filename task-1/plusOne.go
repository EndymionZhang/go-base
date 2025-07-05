package main

func plusOne(digits []int) []int {
	length := len(digits)
	plus := 1
	for i := length - 1; i >= 0; i-- {
		digits[i] += plus
		if digits[i] == 10 {
			digits[i] = 0
			plus = 1
		} else {
			plus = 0
			break
		}
	}
	if plus == 1 {
		digits = append([]int{1}, digits...)
	}
	return digits
}
