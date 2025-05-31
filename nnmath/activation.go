// This file is for activation functions
// All the code is made manually and I dont need external help for this file as I wrote
// my bachelors thesis about deep neural networks with focus on activation functions.
package nnmath

func ReLU(x float64) float64 {
	if x < 0 {
		return 0
	} else {
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
