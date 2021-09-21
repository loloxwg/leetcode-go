package __21

//寻找旋转排序数组中的最小值 二分
func findMin(nums []int) int {
	left, right := 0, len(nums)-1

	for left < right {
		mid := (right-left)>>1 + left
		if nums[mid] > nums[right] {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return nums[left]
}
