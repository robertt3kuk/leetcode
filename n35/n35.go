package n35

func searchInsert(nums []int, target int) int {
	high := len(nums) - 1
	low := 0
	mid := high + low/2
	if nums[low] >= target {
		return low
	}
	for {
		if high <= low {
			break
		}
		if target == nums[mid] {
			return mid
		}
		mid = high + low/2
		if target > nums[mid] {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return high + 1
}
