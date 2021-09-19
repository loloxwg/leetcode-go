package __16

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	vales := []int{}
	for head != nil {
		vales = append(vales, head.Val)
		head = head.Next
	}
	start, end := 0, len(vales)
	for start < len(vales) {
		if vales[start] != vales[end-1] {
			return false
		}
		start++
		end--
	}
	return true
}
