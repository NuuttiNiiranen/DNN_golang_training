package models

import (
	"dnn_golang_training/network"
	"dnn_golang_training/nnmath"
	"math/rand"
)

func FirstModel() *network.NeuralNetwork {
	// Input is 64 numbers for this test (8 by 8 image)
	inputSize := 64

	// Defining neurons for the first hidden layer using a loop and randomized weights and biases. Not using the createLayer method here as it requires previous layer as parameter
	firstHiddenLayer := createLayer(inputSize, nnmath.ReLU, -1, 1, network.NewLayer([]*network.Neuron{}), true, 64)

	// Creating second hidden layer using function
	secondHiddenLayer := createLayer(32, nnmath.ReLU, -1, 1, firstHiddenLayer, false, 0)

	// Creating output layer using function
	outputLayer := createLayer(10, nnmath.ReLU, -1, 1, secondHiddenLayer, false, 0)

	// Defining neural network
	neuralNetwork := network.NewNeuralNetwork([]*network.Layer{firstHiddenLayer, secondHiddenLayer, outputLayer})

	return neuralNetwork
}

// A method to create a slice with given capacity of randomized weights between -1 and 1.
func randomizeWeights(numberOfWeights int) []float64 {
	weights := make([]float64, numberOfWeights)
	for i := range weights {
		weights[i] = 2*rand.Float64() - 1
	}

	return weights
}

// A method for creating a layer with given size and activation function. This function defines connections
// between layers, so knownledge of which was the last layer is needed to create next layer.
func createLayer(numberOfNeurons int, activationFunction func(float64) float64, min float64, max float64, previousLayer *network.Layer, isFirstLayer bool, inputSize int) *network.Layer {
	neuronList := []*network.Neuron{}
	if !isFirstLayer {
		for i := 0; i < numberOfNeurons; i++ {
			neuronList = append(neuronList, network.NewNeuron(min+rand.Float64()*(max-min), activationFunction, randomizeWeights(previousLayer.Length())))
		}
	} else {
		for i := 0; i < numberOfNeurons; i++ {
			neuronList = append(neuronList, network.NewNeuron(min+rand.Float64()*(max-min), activationFunction, randomizeWeights(inputSize)))
		}
	}
	layer := network.NewLayer(neuronList)
	return layer
}
