// This class is entry point for the app. It has training and evaluation.
package main

import (
	"dnn_golang_training/network"
	"dnn_golang_training/nnmath"
)

// I am defining each neuron in the network manually at this stage as I want to see
// that everything works. I will later implement methods to fill layers faster and easier.

func main() {
	// Defining neurons for input layer
	neuron1 := network.NewNeuron(1, nnmath.ReLU, []float64{1, 1})
	neuron2 := network.NewNeuron(1, nnmath.ReLU, []float64{1, 1})
	neuron3 := network.NewNeuron(1, nnmath.ReLU, []float64{1, 1})
	neuron4 := network.NewNeuron(1, nnmath.ReLU, []float64{1, 1})
	neuron5 := network.NewNeuron(1, nnmath.ReLU, []float64{1, 1})
	neuron6 := network.NewNeuron(1, nnmath.ReLU, []float64{1, 1})

	// Defining neurons for the first hidden layer
	hiddenNeuron1 := network.NewNeuron(1, nnmath.ReLU, []float64{1, 1})
	hiddenNeuron2 := network.NewNeuron(1, nnmath.ReLU, []float64{1, 1})
	hiddenNeuron3 := network.NewNeuron(1, nnmath.ReLU, []float64{1, 1})
	hiddenNeuron4 := network.NewNeuron(1, nnmath.ReLU, []float64{1, 1})

	// Defining input layer
	inputLayer := network.NewLayer([]network.Neuron{*neuron1, *neuron2, *neuron3, *neuron4, *neuron5, *neuron6})
	firstHiddenLayer := network.NewLayer([]network.Neuron{*hiddenNeuron1, *hiddenNeuron2, *hiddenNeuron3, *hiddenNeuron4})

	// Defining neural network
	neuralNetwork := network.NewNeuralNetwork([]network.Layer{*inputLayer, *firstHiddenLayer})

	neuralNetwork.PrintLayers()

}
