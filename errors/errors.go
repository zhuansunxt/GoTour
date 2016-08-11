// Re-implement Sqrt funciton to support returning error when argument is
// invalid.
// The purpose of this practice is to get familiar with built-in interface
// error, and how to make use of it to indicate error message.

// An interesting question for this part of practice: a call to fmt.Sprint(e) 
// inside the Error method will send the program into an infinite loop. You can
// avoid this by converting e first: fmt.Sprint(float64(e)). Why?
// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - 
// The answer is, when you invoke fmt.Sprint(e), e is of type ErrNegativeSqrt, 
// which only implements error interface (Error()), but not Stringer interface
// (String()). So system is trying to convert e to type string, which only can 
// be achieved by involing Error() method in error interface, causing a recursive
// method calling without a ending point. The program will crash when it's out 
// of memory.

package main

import (
	"fmt"
	"math"
)

type ErrNegetiveSqrt float64

func (e ErrNegetiveSqrt) Error() string {
	// Why below code doesn't work?
	// return fmt.Sprint("cannot Sqrt negative number ", e)
	// See explanation at the beginning comment section.
	
	return fmt.Sprint("cannot Sqrt negative number ", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return -1, ErrNegetiveSqrt(x)
	}

	init_guess := float64(1)
	conv_threshold := float64(0.001)
	delta := float64(1<<31)
	root := float64(0)
	for ; delta > conv_threshold ; {
		root = init_guess - (math.Pow(init_guess, 2) - x) / (2 * init_guess)
		delta = math.Abs(root - init_guess)
		init_guess = root
	}
	return root, nil;
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}

