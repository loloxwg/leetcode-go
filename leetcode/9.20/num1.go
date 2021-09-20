package __20

import "sort"

//34. 在排序数组中查找元素的第一个和最后一个位置
func searchRange(nums []int, target int) []int {
	leftmost := sort.SearchInts(nums, target)
	if leftmost == len(nums) || nums[leftmost] != target {
		return []int{-1, -1}
	}
	rightmost := sort.SearchInts(nums, target+1) - 1
	return []int{leftmost, rightmost}
}

func searchRange2(nums []int, target int) []int {
	n := len(nums)
	l := 0
	r := n - 1
	targetMid := -1
	start := 0
	end := 0

	for l <= r {
		mid := (l + r) / 2
		if nums[mid] == target {
			targetMid = mid
			break
		}
		//target 在 mid-r之间, l = mid +1
		if nums[mid] < target && target <= nums[r] {
			l = mid + 1
		} else {
			//target 在 l - mid 之间，r = mid -1
			r = mid - 1
		}

	}

	if targetMid == -1 {
		return []int{-1, -1}
	} else {
		//找target的上下边界
		for i := targetMid; i >= 0; i-- {
			if nums[i] == target {
				start = i

			} else {
				break
			}
		}
		for i := targetMid; i < n; i++ {
			if nums[i] == target {
				end = i

			} else {
				break
			}
		}
		return []int{start, end}
	}
}
