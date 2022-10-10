package n189

import (
	"fmt"
	"log"
	"reflect"
	"testing"
)

func err(init, nums, want []int, k int) string {
	str := fmt.Sprintf("wanted %v but got %v rotation with %d steps initial is %v", want, nums, k, init)
	return str
}

func TestFirst(t *testing.T) {
	var init []int
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	init = append(init, nums...)
	k := 3
	want := []int{5, 6, 7, 1, 2, 3, 4}
	rotate(nums, k)
	if !reflect.DeepEqual(nums, want) {
		log.Fatalln(err(init, nums, want, k))
	}
}
func TestSecond(t *testing.T) {

	var init []int
	nums := []int{-1, -100, 3, 99}
	k := 2
	init = append(init, nums...)
	want := []int{3, 99, -1, -100}
	rotate(nums, k)
	if !reflect.DeepEqual(nums, want) {
		log.Fatalln(err(init, nums, want, k))
	}
}
func TestThrid(t *testing.T) {
	var init []int
	nums := []int{1, 2}
	k := 3
	init = append(init, nums...)
	want := []int{2, 1}
	rotate(nums, k)
	if !reflect.DeepEqual(nums, want) {
		log.Fatalln(err(init, nums, want, k))
	}

}
