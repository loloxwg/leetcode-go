package __20

//34. 在排序数组中查找元素的第一个和最后一个位置
func searchRange(nums []int, target int) []int {
	left, right := 0, len(nums)-1
	res := []int{}

	for left <= right {
		for ; left < len(nums); left++ {
			if nums[left] == target {
				res = append(res, left)
				break
			}
			for ; ; right-- {
				if nums[right] == target {
					res = append(res, right)
					break
				}
			}
			if right < left {
				return []int{-1, -1}
			}
		}

	}
	return res

}
