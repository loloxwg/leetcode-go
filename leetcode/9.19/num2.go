package __19

import (
	"strconv"
)

func maximumSwap(num int) int {
	//数字转换为字符数组便于后续处理
	nums := []byte(strconv.Itoa(num))
	//递减单调栈(动态的)
	stk := []int{}

	for i := range nums {
		for len(stk) > 0 && nums[i] > nums[stk[len(stk)-1]] {
			stk = stk[:len(stk)-1]
		}
		stk = append(stk, i)
	}
	// 从前往后找第一个较小的数字
	for i := 0; i < len(stk); i++ {
		if stk[i] != i { //下标和递减栈中不一致，即出现该位置的数字比后面的数字小的情况
			little := i                                                     // 先记录下需要交换的较小的数字下标
			for ; (i+1) < len(stk) && nums[stk[i]] == nums[stk[i+1]]; i++ { //找到需要交换数字的最大下标
			}
			nums[little], nums[stk[i]] = nums[stk[i]], nums[little] //交换数字
			break

		}
	}
	ans, _ := strconv.Atoi(string(nums)) //转换成int返回
	return ans

}
