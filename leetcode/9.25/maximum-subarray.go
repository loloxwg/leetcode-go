package __25

import "studyGo/ulit"

//最大子序列和

func maxSubArray(nums []int)int{
	cur,max:=nums[0],nums[0]
	for i := 0; i <len(nums) ; i++ {
		cur=ulit.Max(nums[i],cur+nums[i])
		max=ulit.Max(max,cur)
	}
	return max

}
