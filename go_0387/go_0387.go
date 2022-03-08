package go0387

func FirstUniqChar(s string) int {
	result := make([]int, 26)
	index := 0
	for i := range s {
		index = int(s[i]) - 97
		result[index]++
	}
	for i := range s {
		index = int(s[i]) - 97
		if result[index] == 1 {
			return i
		}
	}
	return -1
}

func firstUniqChar1(s string) int {
	mapResult := make(map[byte]int)
	for i := range s {
		mapResult[s[i]]++
	}
	for i := range s {
		if mapResult[s[i]] == 1 {
			return i
		}
	}
	return -1
}
