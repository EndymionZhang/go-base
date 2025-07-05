package main

func singleNumber(nums []int) int {
	maps := make(map[int]int)
	for _, i := range nums {
		v, exist := maps[i]
		if exist {
			maps[i] = v + 1
		} else {
			maps[i] = 1
		}
	}

	for k, v := range maps {
		if v == 1 {
			return k
		}
	}
	return 0
}

// 任何数和0做异或运算，结果仍然是原来的数，即 0 ^ num = num
// 任何数和它自身做异或运算，结果是0，即 num ^ num = 0
// 异或运算满足交换律和结合律，即 a ^ b ^ a = b ^ a ^ a = b ^ (a ^ a) = b ^ 0 = b
func singleNumberV2(nums []int) int {
	single := 0
	for _, num := range nums {
		single ^= num
	}
	return single
}
