package main

func longestCommonPrefix(strs []string) string {
	if len(strs) == 1 {
		return strs[0]
	}
	runes := []byte{}
	fistStr := strs[0]
	for i := 0; i < len(fistStr); i++ {
		for j := 1; j < len(strs); j++ {
			if i >= len(strs[j]) || strs[j][i] != fistStr[i] {
				return string(runes)
			}
		}
		runes = append(runes, fistStr[i])
	}
	return ""
}
