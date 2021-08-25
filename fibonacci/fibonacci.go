package fibonacci

import (
	"errors"
	"fmt"
)

func FibonacciSlice(x int, y int) ([]uint64, error) {
	if y-x < 0 {
		return nil, errors.New(fmt.Sprintf("Левая граница x: %v больше чем правая y: %v", x, y))
	}
	var res []uint64

	if x >= 0 && x <= 3 {
		var initValues = [3]uint64{0, 1, 1}
		var initBounds = 3
		if y < 3 {
			initBounds = y + 1
		}
		res = initValues[x:initBounds]
	} else {
		res = make([]uint64, 0)
	}
	var first, second uint64 = 1, 1
	for i := 3; i <= y; i++ {
		var sum = first + second
		first = second
		second = sum
		if x <= i {
			res = append(res, second)
		}
	}
	return res, nil
}
