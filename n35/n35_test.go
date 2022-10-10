package n35

import (
	"fmt"
	"log"
	"testing"
)

func Fail(want, ret int) string {
	str := fmt.Sprintf("want %d but got: %d", want, ret)
	return str
}

func TestFirst(t *testing.T) {
	nums := []int{1, 3, 5, 6}
	target := 5
	want := 2
	ret := searchInsert(nums, target)
	if ret != want {
		log.Fatalf(Fail(want, ret))
	}
}
func TestSecond(t *testing.T) {
	nums := []int{1, 3, 5, 6}
	target := 2
	want := 1
	ret := searchInsert(nums, target)
	if ret != want {
		log.Fatalf(Fail(want, ret))
	}
}
func TestThird(t *testing.T) {
	nums := []int{1, 3, 5, 6}
	target := 7
	want := 4
	ret := searchInsert(nums, target)
	if ret != want {
		log.Fatalf(Fail(want, ret))
	}
}
