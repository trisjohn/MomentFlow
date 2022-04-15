package neural

import (
	"fmt"
	// "math"
	"math/rand"
	"time"
)

type NeuralNetwork struct {
	inputs, hiddens, outputs []float64
	learn_rate float64
	weights_ih, weights_ho, bias_h, bias_o Matrix
}

func NewNeuralNetwork(inputN, hiddenN, outputN int, learnrate float64) (N NeuralNetwork) {
	N = NeuralNetwork{}
	rand.Seed(time.Now().UnixNano())
	N.inputs = make([]float64, inputN)
	N.hiddens = make([]float64, hiddenN)
	N.outputs = make([]float64, outputN)
	N.learn_rate = learnrate
	N.weights_ih = NewMatrix(hiddenN, inputN)
	N.weights_ho = NewMatrix(outputN, hiddenN)
	N.bias_h = NewMatrix(hiddenN, 1)
	N.bias_o = NewMatrix(outputN, 1)
	N.weights_ih.FillFloat(.1)
	N.weights_ho.FillFloat(.1)
	N.bias_o.FillFloat(.1)
	N.bias_h.FillFloat(.1)
	// N.weights_ih.Randomize()
	// N.weights_ho.Randomize()
	// N.bias_h.Randomize()
	// N.bias_o.Randomize()
	fmt.Println("New Neural Network Created.")
	return
}

func (n *NeuralNetwork) FeedForward(input_array []float64)  []float64 {
	inputs := MatrixFromArray(input_array)
	hidden := Product(n.weights_ih, inputs)
	hidden.Add(n.bias_h)
	hidden.Map(Relu)
	output := Product(n.weights_ho, hidden)
	output.Add(n.bias_o)
	// For multi category classifcations
	// output = SoftMax(output)
	// General Use
	// output.Map(math.Tanh)
	output.Map(Relu)
	return output.ToArray()
}

func (n *NeuralNetwork) Train(inputs_arr, targets_arr []float64) {
	inputs := MatrixFromArray(inputs_arr)
	targets := MatrixFromArray(targets_arr)
	hidden := Product(n.weights_ih, inputs)
	hidden.Add(n.bias_h)
	hidden.Map(Relu)
	
	outputs := Product(n.weights_ho, hidden)
	outputs.Add(n.bias_o)
	// outputs.Map(Sigmoid)
	outputs.Map(Relu)
	// outputs.Map(math.Tanh)
	fmt.Println("OUTPUTS:",outputs)

	output_errors := Subtract(targets, outputs)
	// Calculate Gradient
	// gradients := MatrixMap(outputs,dsigmoid)
	// gradients := MatrixMap(outputs,dtanh)
	gradients := MatrixMap(outputs,drelu)
	gradients.Product(output_errors)
	gradients.Multiply(n.learn_rate)
	fmt.Println("GRADIENTS:",gradients)

	// // Calculate deltas
	hidden_t := Transpose(hidden)
	w_ho_deltas := Product(gradients,hidden_t)
	
	// // Adjust the weights by deltas
	n.weights_ho.Add(w_ho_deltas)
	n.bias_o.Add(gradients)

	// // Calculate the hidden layer errors
	who_t := Transpose(n.weights_ho)
	hidden_errors := Product(who_t, output_errors)

	// // Calculate hiden gradient
	hidden_gradient := MatrixMap(hidden,drelu)
	hidden_gradient.Product(hidden_errors)
	hidden_gradient.Multiply(n.learn_rate)

	// // Calculate input to hidden deltas
	inputs_t := Transpose(inputs)

	weight_ih_deltas := Product(hidden_gradient, inputs_t)

	n.weights_ih.Add(weight_ih_deltas)
	n.bias_h.Add(hidden_gradient)
}

// A Basic Neural Network with one input layer, one hidden layer, and one output layer
// If number of out puts is len 3 or more, output layer has a softmax activation function
// Otherwise output layer uses a simple linregression
func NewBaseModel(in_num, hidden_num, out_num int, learn_rate float64) (n NeuralNetwork) {
	n = NewNeuralNetwork(in_num,hidden_num,out_num,learn_rate)
	return
}