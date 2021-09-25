package __25

import  "studyGo/ulit"

func rob(nums []int) int {
	len:=len(nums)
	if len == 0{
		return 0
	}
	if len==1 {
		return nums[0]
	}
	dp:=make([]int,len)
	dp[0]=nums[0]
	dp[1]=ulit.Max(nums[0],nums[1])
	for i := 2; i <len; i++ {
		dp[i]=ulit.Max(dp[i-1],dp[i-2]+nums[i])
	}
	return dp[len]
}

//滚动数组
func rob_rr(nums []int) int {
	len:=len(nums)
	if len == 0{
		return 0
	}
	if len==1 {
		return nums[0]
	}

	first :=nums[0]
	second :=ulit.Max(nums[0],nums[1])
	for i := 2; i <len; i++ {
		first,second=second,max(second,first+nums[i])
	}
	return second
}

