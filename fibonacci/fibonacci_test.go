package fibonacci

import (
	"errors"
	"reflect"
	"testing"
)

var expectingFibonacciArray = []uint64{0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377, 610, 987, 1597, 2584, 4181, 6765, 10946, 17711, 28657, 46368, 75025, 121393, 196418, 317811, 514229, 832040, 1346269, 2178309, 3524578, 5702887, 9227465, 14930352, 24157817, 39088169, 63245986, 102334155}
var expectingFibonacciArray1 = []uint64{0, 1}
var expectingFibonacciArray2 = []uint64{1, 1}
var expectingFibonacciArray3 = []uint64{5}
var expectingFibonacciArray4 = []uint64{21, 34, 55}
var expectingFibonacciArray5 = []uint64{21, 34, 55, 89, 144, 233, 377, 610}

func getSec(x uint64, posX int, y uint64, posY int) *FibonacciInitSection {
	return &FibonacciInitSection{Fn: FibonacciNumber{Value: x, Position: posX}, Sn: FibonacciNumber{Value: y, Position: posY}}
}

func compareWithSec(x int, y int, compareSlice []uint64, fis *FibonacciInitSection) error {
	fibSlice, err := FibonacciSlice(x, y, fis)
	if err != nil {
		return err
	}
	if reflect.DeepEqual(fibSlice, compareSlice) == false {
		return errors.New("fail test")
	}
	return nil
}

func TestFibonacci(t *testing.T) {

	if err := compareWithSec(0, 40, expectingFibonacciArray, nil); err != nil {
		t.Fatal(err.Error())
	}

	if err := compareWithSec(0, 1, expectingFibonacciArray1, nil); err != nil {
		t.Fatal(err.Error())
	}

	if err := compareWithSec(1, 2, expectingFibonacciArray2, nil); err != nil {
		t.Fatal(err.Error())
	}

	if err := compareWithSec(5, 5, expectingFibonacciArray3, nil); err != nil {
		t.Fatal(err.Error())
	}

	if err := compareWithSec(8, 10, expectingFibonacciArray4, nil); err != nil {
		t.Fatal(err.Error())
	}

	// 				With init values
	// ------------------------------------------------------------------------------

	sec1 := getSec(0, 0, 1, 1)
	sec2 := getSec(21, 8, 34, 9)
	sec3 := getSec(1, 1, 1, 1)
	sec4 := getSec(3, 4, 5, 5)
	sec5 := getSec(0, 0, 5, 5)

	if err := compareWithSec(0, 40, expectingFibonacciArray, sec1); err != nil {
		t.Fatal(err.Error())
	}
	if err := compareWithSec(0, 40, expectingFibonacciArray, sec2); err != nil {
		t.Fatal(err.Error())
	}
	if err := compareWithSec(0, 40, expectingFibonacciArray, sec3); err != nil {
		t.Fatal(err.Error())
	}
	if err := compareWithSec(0, 40, expectingFibonacciArray, sec4); err != nil {
		t.Fatal(err.Error())
	}
	if err := compareWithSec(0, 40, expectingFibonacciArray, sec5); err != nil {
		t.Fatal(err.Error())
	}

	if err := compareWithSec(8, 15, expectingFibonacciArray5, sec2); err != nil {
		t.Fatal(err.Error())
	}

	if err := compareWithSec(5, 5, expectingFibonacciArray3, sec2); err != nil {
		t.Fatal(err.Error())
	}
}
