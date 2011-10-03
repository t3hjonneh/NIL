package main

import (
	"fmt";
	"hmm";
	"strconv";
)

/*
For this main function we will be making calls to other functions which can be used to test individual components.
The main function should be sectioned off into two parts. The top part should be used to call test functions. The bottom part should be used for code for the project as a whole.
This setup will allow us to quickly and easily comment out unwanted code to test/build only the relevant code.
Either section may be commented out by removing the first / at the beginning of the seciton.
Likewise the code can be restored by replacing the first / at the beginning of the section (that is to say, /* for commented out and //* for commented in).
Naming convention: Test functions should have "test_" at the front of them.
*/


func test_atomaton() {
	counters := make([]float64, 4)
	numruns := 10000.0
	
	for i := 0; i < len(counters); i++ {
		counters[i] = 0.0
	}

	for i := 0.0; i < numruns; i++ {
		counters[hmm.Getnextstate()]++
	}

	for i := 0; i < len(counters); i++ {
		fmt.Printf("state " + strconv.Itoa(i) + " percentage = " + strconv.Ftoa64((counters[i] * 100)/numruns, 'f', 5) + "\n")
	}
}

func main() {
//********************************************
// this is the section for test code

	test_atomaton()

//********************************************/
//********************************************
// this is for the final project

//********************************************/
}
