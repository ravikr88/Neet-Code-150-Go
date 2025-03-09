package validanagram

func ValidAnagram(word, test string) bool {
	arr := [26]int{}

	if len(word) != len(test) {
		return false
	}

	for i := range len(word) {
		arr[word[i]-'a']++
		arr[test[i]-'a']--
	}

	for _, val := range arr {
		if val != 0 {
			return false
		}
	}
	return true
}
