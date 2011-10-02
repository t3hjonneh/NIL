package hmm

import (
	"fmt";
	"rand";
)

func Printmm() {
	fmt.Printf("second print out\n")
}

// this function is the first iteration for the states to be learned in the hmm
// the current probabilities should be as follows:
// 0 = 15%
// 1 = 20%
// 2 = 30%
// 3 = 35%
func Getnextstate() int {
	r := rand.Intn(100) // it is worth noting that this function will return 0 but will not return 100
	
	ret := -1
	
	switch {
		case r < 15:
			ret = 0
		case r < 35:
			ret = 1
		case r < 65:
			ret = 2
		default:
			ret = 3
	}

	return ret
}
