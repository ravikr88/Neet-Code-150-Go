package containsduplicate

func ContainsDuplicate(nums []int) bool {
	hashSet := make(map[int]struct{})

	for _, num := range nums {
		if _, exists := hashSet[num]; exists {
			return true
		}
		hashSet[num] = struct{}{}
	}

	return false
}
