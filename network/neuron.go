// This file is for neurons, that will be used in layers. Neuron shouold have weights, bias terms and activation function.
// Idea for struct Neuron and function Activate is generated by ChatGPT. Further functions are my own doing.
// I have also commented everything manually to understand the code and to make it understandable
package network

import (
	"dnn_golang_training/nnmath"
	"fmt"
)

// Each neuron remembers all the INCOMING weights, and only outputs single value.

// Creating base struct for neuron. Bias term and weights are the parameters
// for each neuron and activation function is the activation function chosen
// for the neuron. Activation function is chosen here, as I want to be able to
// change it for each neuron.

type Neuron struct {
	BiasTerm           float64
	ActivationFunction string
	Weights            []float64
}

// Function to calculate activation for a neuron.
// n *Neuron is a method receiver that can modify and read actual data of a neuron
// gets list of inputs as parameter, this is neurons of previous layer
// (or inputs if its input layer)
// returns single float64 as output of a neuron
func (n *Neuron) Activate(inputs []float64) float64 {
	// Sum of all the inputs
	sum := 0.0
	// Loops through all the weighs in neurons Weights list and calculates
	// weight times input using same index in inputs list. Adds result to sum.
	for i, w := range n.Weights {
		sum += w * inputs[i]
	}
	// Adds bias term to the sum
	sum += n.BiasTerm
	// applies activation function to the sum and returns the value
	activationFunc := nnmath.ActivationFunctions[n.ActivationFunction]
	return activationFunc(sum)
}

// Creating constructor so we can create neurons
func NewNeuron(biasTerm float64, activationFunction string, weights []float64) *Neuron {
	temp := &Neuron{
		BiasTerm:           biasTerm,
		ActivationFunction: activationFunction,
		Weights:            weights,
	}
	return temp
}

// Creating functio for changing weights. This will just set weight as given value, and should not be used
// other than for testing purposes
func (n *Neuron) SetWeight(newWeightNumber int, newWeightValue float64, newBias float64) {
	// Only set weight if new value isn't 0, this allows only setting bias term and keeping weight untouched
	if newWeightValue != 0 {
		n.Weights[newWeightNumber] = newWeightValue
	}
	// Only set bias term if new value isn't 0, this allows only setting weight and keeping bias term untouched
	if newBias != 0 {
		n.BiasTerm = newBias
	}
}

// Method for printing all the weights of single node. This is for debugging and testing purposes and should not
// be used during training or use of model
func (n *Neuron) PrintNeuronsWeights() {
	for i := 0; i < len(n.Weights); i++ {
		fmt.Print(n.Weights[i], "\n")
	}
}
