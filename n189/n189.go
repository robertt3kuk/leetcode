package n189

/*
	 func rotate(nums []int, k int) {
		for k != 0 {
			Rot(nums)
			k--
		}

}

	func Rot(nums []int) {
		var n []int
		n = append(n, nums...)
		for i := 0; i < len(n); i++ {
			if i == 0 {
				nums[0] = n[len(n)-1]
				continue
			}

			nums[i] = n[i-1]
		}

}
*/
func rotate(nums []int, k int) {
	var n []int
	for k > len(nums) {
		k -= len(nums)
	}

	n = append(n, nums...)
	for i := 0; i < len(n); i++ {
		if i-k < 0 {
			nums[i] = n[i+len(n)-k]
		} else {
			nums[i] = n[i-k]
		}

	}

}
