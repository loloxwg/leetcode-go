package __15

import (
	"fmt"
	"reflect"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	nums := []int{1, 2, 3, 3, 4}
	got := RemoveDuplicates(nums)
	want := 4
	err := reflect.DeepEqual(got, want)
	if err != true {
		fmt.Printf("got=%v want=%v", got, want)
	}

}
