package test_test

import (
	"fmt"
	"testing"
)
func TestAdd(t *testing.T) {
	num := add(1)
	t.Logf("test success")
	fmt.Println(num)

}
func add(num int, ) int {
	num += 1
	return num
}