package __20

//33. 搜索旋转排序数组
func search(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := 1 + (r-1)>>1
		if nums[mid] == target {
			return mid
		} else if nums[mid] >= nums[0] {
			//左边有序
			if nums[0] <= target && target <= nums[mid] {
				r = mid - 1
			} else {
				l = mid + 1
			}
		} else {
			if nums[mid] <= target && target <= nums[len(nums)-1] {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}

	}
	return -1
}
