package n283

func moveZeroes(nums []int) {
	last := len(nums) - 1
	for i := 0; i < last; i++ {
		if nums[i] == 0 {
			for d := i; d < last; d++ {
				if d < last {
					nums[d] = nums[d+1]
				}
			}
			nums[last] = 0
			last--
			i--
		}
	}

}
