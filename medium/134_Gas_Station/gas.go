package main

import "fmt"
import "math/rand"
import "time"

func canCompleteCircuit(gas []int, cost []int) int {
	n := len(gas)
	if n == 0 {
		return -1
	}
	f, i, j := 0, 0, 0
	for i < n {
		f += (gas[j] - cost[j])
		//fmt.Printf("#i %d, j %d, f %d\n", i, j, f)
		j++
		if j >= n {
			j -= n
		}
		if f >= 0 && j == i {
			return i
		}
		if f >= 0 {
			continue
		}
		for i < n {
			w := gas[i] - cost[i]
			f -= w
			i++
			//fmt.Printf("$i %d, j %d, f %d\n", i, j, f)
			if w < 0 && f >= 0 {
				break
			}
			if i == j {
				break
			}
		}
	}
	return -1
}

func bf(gas, cost []int) int {
	//fmt.Printf("%v, %v\n", gas, cost)
	n := len(gas)
	var i, j, f int
	for i = 0; i < n; i++ {
		f, j = 0, i
		for {
			f += (gas[j] - cost[j])
			//fmt.Printf("i %d, j %d, f %d\n", i, j, f)
			j++
			if j >= n {
				j -= n
			}
			if f < 0 || j == i {
				break
			}
		}
		if f >= 0 && j == i {
			return i
		}
	}
	return -1
}

func MakeRandArray() []int {
	maxLen, maxElement := 10, 20
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	len := r.Int() % maxLen
	a := make([]int, len)
	for i := 0; i < len; i++ {
		a[i] = r.Int()%maxElement - 10
	}
	return a
}

func testGas() {
	gas := MakeRandArray()
	cost := make([]int, len(gas))
	//fmt.Printf("%v\n", gas)
	res := canCompleteCircuit(gas, cost)
	//fmt.Printf("%v, get %d\n", gas, res)
	ans := bf(gas, cost)
	//fmt.Printf("%v, get %d\n", gas, ans)
	if res != ans {
		panic(fmt.Sprintf("%v, get %d, ans %d", gas, res, ans))
	}
}

func main() {
	for i := 0; i < 10000; i++ {
		testGas()
	}
}
