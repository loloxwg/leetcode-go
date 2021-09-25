package __25


func MaxValue(n int)int{
	v:=[5]int{3,7,11,15,20}
	w:=[]int{1,2,3,4,5}
	dp:=[6][]int{}
	for i := 0; i < len(w); i++ {
		for j := 0; j <=n ; j++ {
			if j<w[i] {
				dp[i+1][j]=dp[i][j]
			}else {
				dp[i+1][j]=max(dp[i][j],dp[i+1][j-w[i]]+v[i])
			}
		}
	}
	return dp[5][n]

}

func max(a, b int) int {
	if a>b{
		return a
	}
	return b
}