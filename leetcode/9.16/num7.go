package __16

func maxSlidingWindow(nums []int, k int) []int {

	if k == 1 {
		return nums
	}
	queue := []int{0}
	var res = []int{}
	for i := 1; i < len(nums); i++ {
		if i-queue[0] >= k {
			queue = queue[1:]
		}
	}
	return res

}
