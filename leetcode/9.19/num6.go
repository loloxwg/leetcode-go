package __19

/*1137. 第 N 个泰波那契数*/
//泰波那契序列Tn定义如下:
//
//T0 = 0, T1 = 1, T2 = 1, 且在 n >= 0
//的条件下 Tn+3 = Tn + Tn+1 + Tn+2

func tribonacci(n int) int {
	if n == 0 {
		return 0
	}
	if n <= 2 {
		return 1
	}
	p, q, r, s := 0, 0, 1, 1
	for i := 3; i <= n; i++ {
		p = q
		q = r
		r = s
		s = p + q + r
	}
	return s

}
func tribonacci2(n int) int {
	if n == 0 {
		return 0
	}
	if n <= 2 {
		return 1
	}
	return tribonacci2(n-1) + tribonacci2(n-2) + tribonacci2(n-3)

}
