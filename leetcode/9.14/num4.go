package __14

import __15 "studyGo/leetcode/9.15"

//作者：xiao-hao-8o
//链接：https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/solution/gohua-dong-chuang-kou-xun-zhao-zi-fu-chu-eis5/
//来源：力扣（LeetCode）
//著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

func lengthOfLongestSubstring(s string) int {
	strlen := len(s)
	if strlen == 0 {
		return 0
	}
	left, right, ans := 0, 0, 0
	m := map[byte]int{}
	for right < strlen {
		if _, ok := m[s[right]]; !ok {
			m[s[right]] = right
		} else {
			if m[s[right]]+1 >= left {
				left = m[s[right]] + 1
			}
			m[s[right]] = right //!!
		}
		ans = __15.Max(right-left+1, ans)
		right++
	}
	return ans
}
