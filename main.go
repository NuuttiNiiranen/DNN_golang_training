// This class is entry point for the app. It has training and evaluation.
package main

import "dnn_golang_training/models"

func main() {
	// Creating a simple neural network model
	neuralNetwork := models.FirstModel()
	// Running the first model with input of 64 1s
	neuralNetwork.Predict([]float64{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1})
}
