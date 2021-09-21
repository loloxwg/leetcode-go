package __21

//典型的动态规划题目
//确定初始条件: f(1) = 1, f(2) = 2
//确定状态方程: f(n) = f(n-1) + f(n-2)
//确定最终答案: f(n) 即为最终的答案

//1: 1        1
//2: 2
//3: 3		1+1+1 ,1+2  2+1
//4: 5        1+1+1+1 ,2+1+1,1+2+1,1+1+2，2+2
//f(n)=f(n-1)+f(n-2)
func climbStairs(n int) int {
	if n <= 2 {
		return n
	}
	before, after := 1, 2
	for i := 3; i <= n; i++ {
		before, after = after, before+after
	}
	return after
}
