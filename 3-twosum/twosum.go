package twosum

func Twosum_bf(nums []int, target int) []int {

	// var ans []int

	for i := range len(nums) - 1 {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}

	return nil //ans
}

func Twosum_HashMap(arr []int, target int) []int {
	// var ans []int

	hashMap := make(map[int]int)

	for ix, num := range arr {
		if val, ok := hashMap[num]; ok {
			// if num is present, this means its corresponding number's index is val
			return []int{val, ix}
		}
		hashMap[target-num] = ix // store corresponding number, with its index as value
	}
	return nil

}
