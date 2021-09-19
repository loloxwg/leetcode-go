package __15

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMaxProfits(t *testing.T) {
	nums := []int{7, 1, 5, 3, 6, 4}
	var got = MaxProfit(nums)
	want := 4
	err := reflect.DeepEqual(got, want)
	if err != true {
		fmt.Printf("got=%v want=%v", got, want)
	}

}
