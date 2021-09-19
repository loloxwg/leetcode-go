package studyGo

import (
	"fmt"
	"studyGo/leetcode/9.15"
)

func main() {
	nums := []int{1, 2, 3, 3, 4}
	got := __15.RemoveDuplicates(nums)
	want := 4

	fmt.Printf("got=%v want=%v", got, want)

}
