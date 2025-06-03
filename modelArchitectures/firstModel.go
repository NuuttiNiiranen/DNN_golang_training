package modelArchitectures

import (
	"dnn_golang_training/network"
	"math/rand"
)

func FirstModel() *network.NeuralNetwork {
	// Input is 64 numbers for this test (8 by 8 image)
	inputSize := 64

	// Defining neurons for the first hidden layer
	firstHiddenLayer := createLayer(inputSize, "relu", -1, 1, network.NewLayer([]*network.Neuron{}), true, 64)

	// Creating second hidden layer
	secondHiddenLayer := createLayer(32, "relu", -1, 1, firstHiddenLayer, false, 0)

	// Creating output layer
	outputLayer := createLayer(10, "relu", -1, 1, secondHiddenLayer, false, 0)

	// Defining neural network from layers I just made
	neuralNetwork := network.NewNeuralNetwork([]*network.Layer{firstHiddenLayer, secondHiddenLayer, outputLayer})

	return neuralNetwork
}

// A method to create a slice with given capacity of randomized weights between -1 and 1.
// TODO change limits to parameters instead of hard-coded, cba rn as Im at work
func randomizeWeights(numberOfWeights int) []float64 {
	weights := make([]float64, numberOfWeights)
	for i := range weights {
		weights[i] = 2*rand.Float64() - 1
	}

	return weights
}

// A method for creating a layer with given size and activation function. This function defines connections
// between layers, so knownledge of which was the last layer is needed to create next layer,
// this limitation has been skipped with boolean. I had great idea for this as shower thought, but I forgot it
// (hopefully temporarily), and will fix later to get rid of one parameter.
// replace previouslayer with previouslayer.length so length is calculated in calling method, not here
func createLayer(numberOfNeurons int, activationFunction string, min float64, max float64, previousLayer *network.Layer, isFirstLayer bool, inputSize int) *network.Layer {
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
