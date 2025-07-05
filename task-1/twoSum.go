package main

func twoSum(nums []int, target int) []int {
	maps := make(map[int]int)
	for i, v := range nums {
		index, exist := maps[target-v]
		if exist {
			return []int{i, index}
		}
		maps[v] = i
	}
	return nil
}
