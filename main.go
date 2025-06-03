// This class is entry point for the app. It has training and evaluation.
package main

func main() {
	// Creating a simple neural network model
	//neuralNetwork := modelArchitectures.FirstModel()

	// Adjusting single weights and biases of neural network
	// neuralNetwork.AdjustWeights(2, 1, 1, 2.5, 2)
	// neuralNetwork.AdjustWeights(2, 1, 2, 2.5, 2)

	// Testing loading a file
	//byteText, err := LoadFileIfExists("models/FirstModel")
	//if err != nil {
	//	fmt.Print(err)
	//}

	// Making a network out of loaded []byte
	// neuralNetworkModel := network.LoadNeuralNetworkFromJSON(byteText)

	//Loading a dataset
	//numbers := data.LoadDataset("data/digits_dataset.json")

	// Testing network with loaded dataset
	//neuralNetworkModel.Predict(numbers[0].NumberVector)
}
