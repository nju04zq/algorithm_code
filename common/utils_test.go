package utils

import "fmt"
import "testing"

func TestHelloWorld(t *testing.T) {
	a := MakeRandArray()
	fmt.Println(a)
	b := MakeRandSortedArray()
	fmt.Println(b)
	vals := []string{"1", "2", "3", "#", "#", "4", "5", "#", "6", "7"}
	root := makeTree(vals)
	fmt.Println(vals)
	dumpTree(root)
}
