func twoSum(nums []int, target int) []int {

	indices := [][]int{}

	for i := 0; i < len(nums); i++ {
		for j := i; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				if i != j {
					indices = append(indices, []int{i, j})
				}
			}
		}
	}

	fmt.Println(indices)
	return indices[0]
}