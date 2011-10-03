package hmm

import (
	"fmt";
//	"rand";
	"strconv";
)

/*
This file will contain the primary elements of the hidden markov model.
This file, combined with n-grams will form the basis of the voice recognition system.
TODO: please note that the function calls as they exist now probably will not make it into the final version. I will take any requests...
*/

var occurances = make([]float64, 4)
var numruns = 10000.0

func printmodel() {
	for i := 0; i < len(occurances); i++ {
		fmt.Printf("state " + strconv.Itoa(i) + " percentage = " + strconv.Ftoa64((occurances[i] * 100)/numruns, 'f', 5) + "\n")
	}
}

func train() {
	for i := 0.0; i < numruns; i++ {
		occurances[Getnextstate()]++
	}
}

func initialize() {
	//occurances := make([]float64, 4)
	for i := 0; i < len(occurances); i++ {
		occurances[i] = 0.0
	}
}

func Model() {
	initialize()
	train()
	printmodel()
}
