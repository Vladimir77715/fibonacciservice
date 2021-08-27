package fibonacci

import (
	"errors"
	"fmt"
)

type FibonacciInitSection struct {
	Fn FibonacciNumber
	Sn FibonacciNumber
}

type FibonacciNumber struct {
	Value    uint64
	Position int
}

func FibonacciSlice(x int, y int, fib *FibonacciInitSection) ([]uint64, error) {
	var size = y - x + 1
	if y-x < 0 {
		return nil, errors.New(fmt.Sprintf("Левая граница x: %v больше чем правая y: %v", x, y))
	}
	var res = make([]uint64, size)

	resIndex := 0
	var first, second uint64 = 1, 1
	var initPos = 3

	if x >= 0 && x <= 3 {
		var initValues = [3]uint64{0, 1, 1}
		var initBounds = 3
		if y < 3 {
			initBounds = y + 1
		}
		subInit := initValues[x:initBounds]
		lastI := 0
		for i, _ := range subInit {
			res[i] = subInit[i]
			lastI = i
		}
		resIndex = lastI + 1
	}

	if fib != nil {
		if fib.Fn.Position == x && fib.Sn.Position == x+1 {
			first, second = fib.Fn.Value, fib.Sn.Value
			if len(res) == 1 {
				res[0] = first
			} else {
				res[0], res[1] = first, second
			}
			initPos = fib.Sn.Position + 1
			resIndex = 2
		} else if fib.Fn.Position == x-1 && fib.Sn.Position == x {
			first, second = fib.Sn.Value, fib.Fn.Value+fib.Sn.Value
			if len(res) == 1 {
				res[0] = first
			} else {
				res[0], res[1] = first, second
			}
			initPos = fib.Sn.Position + 1
			resIndex = 2
		} else if fib.Fn.Position == x-2 && fib.Sn.Position == x-1 {
			first = fib.Fn.Value + fib.Sn.Value
			second = first + fib.Sn.Value
			if len(res) == 1 {
				res[0] = first
			} else {
				res[0], res[1] = first, second
			}
			initPos = fib.Sn.Position + 1
			resIndex = 2
		}
	}

	for i := initPos; i <= y && resIndex < len(res); i++ {
		var sum = first + second
		first = second
		second = sum
		if x <= i {
			res[resIndex] = second
			resIndex++
		}
	}
	return res, nil
}
