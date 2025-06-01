// This file is for activation functions
// All the code is made manually and I dont need external help for this file as I wrote
// my bachelors thesis about deep neural networks with focus on activation functions.
package nnmath

import "fmt"

var ActivationFunctions = map[string]func(float64) float64{
	"relu":   ReLU,
	"linear": linear,
	"lrelu":  LReLU,
}

func ReLU(x float64) float64 {
	if x < 0 {
		fmt.Print("Relu called with value: ", x, "\n")
		return 0
	} else {
		fmt.Print("Relu called with value: ", x, "\n")
		return x
	}

}

func linear(x float64) float64 {
	return x
}

func LReLU(x float64) float64 {
	if x < 0 {
		return 0.1 * x
	} else {
		return x
	}
}
