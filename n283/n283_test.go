package n283

import (
	"fmt"
	"log"
	"reflect"
	"testing"
)

func err(init, nums, want []int) string {
	str := fmt.Sprintf("wanted: %v but got %v initial is %v", want, nums, init)
	return str
}

func TestFirst(t *testing.T) {
	nums := []int{0, 1, 0, 3, 12}
	var init []int
	init = append(init, nums...)
	want := []int{1, 3, 12, 0, 0}
	moveZeroes(nums)
	if !reflect.DeepEqual(want, nums) {
		log.Fatalln(err(init, nums, want))
	}
}
func TestSecond(t *testing.T) {
	nums := []int{0}
	want := []int{0}
	moveZeroes(nums)
	if !reflect.DeepEqual(want, nums) {
		log.Fatalln(err(want, nums, want))
	}
}
func TestThrid(t *testing.T) {
	nums := []int{0, 0, 1}
	want := []int{1, 0, 0}
	var init []int
	init = append(init, nums...)
	moveZeroes(nums)
	if !reflect.DeepEqual(want, nums) {
		log.Fatalln(err(init, nums, want))
	}
}
