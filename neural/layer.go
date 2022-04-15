package neural

import (
	"fmt"
	"math"
	"math/rand"
)

// ACTIVATION FUNCTIONS

func Sigmoid(n float64) (val float64) {
	val = 1 / (1 + math.Exp(-n))
	return
}

// Takes a matrix, assumed as the output, then creates a new matrix and passes it through the soft max
func SoftMax(m Matrix) (result Matrix) {
	var sum float64
	result = NewMatrix(m.rows, m.cols)
	for i,v := range m.values {
		sum += math.Exp(v)
		result.values[i] = v
	}
	result.Map(func(v float64) float64 {
		return math.Exp(v) / sum
	})
	return
}

// Pass in an input matrix to compute the softmax derivative on, returns a new matrix after calculation
// Assumes input is softmax
func SoftMaxDeriv(m Matrix) (result Matrix) {
	result = NewMatrix(m.rows, m.cols)
	for i,v := range m.values {
		result.values[i] = v * (1.0 - v)
	}
	return
}

// params: n float64, isderiv bool
func Relu(k float64) (n float64) {
	n = math.Max(0,k)
	return
}

// Simple line function that uses an approximation of uelers number as its constant
func ELine(x float64) float64 {
	return math.E * x + 1
}

func drelu(k float64) (n float64) {
	if k < 0 {
		n = 0
	} else if k > 1 {
		n = 1
	} else {
		n = rand.Float64()
	}
	return
}

// Given the output values of a function F return the derivative of that function f(x)
// Where x is the output of F(k)
func dsigmoid(k float64) (n float64) {
	return k * (1.0 - k)
}

func dtanh(x float64) float64 {
	var y = 1 / (math.Pow(math.Cosh(x), 2));
	return y;
  }

// This adjust weights ever so slightly
func mutate(x float64) float64 {
	val := rand.Float64()
	if val < .1 {
		val = 1
		if rand.Float32() < .5 {
			val = -1
		}
		return x + val*rand.Float64()*.5
	}else {
		return x
	}
}

func SetActivation(l* Layer, name string) {
	switch name {
	case "Tanh":
		l.activation = math.Tanh
		l.derivative = dtanh
	case "Relu": 
		l.activation = Relu
		l.derivative = drelu
	case "Sigmoid": 
		l.activation = Sigmoid
		l.derivative = dsigmoid
	case "Softmax":
		// Special case
		l.activation = nil
		l.derivative = nil
	}
}

// A layer for a neural network, can feed forward and can propogate backwards
// Product weights by inputs
// Add bias to result
// Pass through activation function
type Layer struct {
	// Weights is node_num x input_num
	// Bias is node_num x 1
	weights, biases, outputs, inputs Matrix
	activation, derivative func(float64) float64
}

func NewLayer(node_num int, activation string) (l Layer) {
	l.weights = NewMatrix(node_num,1) // Update later when gathering inputs
	l.biases = NewMatrix(node_num,1)
	l.biases.Randomize()
	SetActivation(&l, activation)
	fmt.Println("New Layer Created.")
	return
}

// Pass in a layer which is the input to this layer
// This layers inputs are the given layer's outputs
func (l *Layer) Input(in Matrix) {
	l.inputs = in
	if l.weights.cols != in.rows {
		// Weights not set, randomize and set
		l.weights = NewMatrix(l.weights.rows, in.rows)
		l.weights.Randomize()
	}
	l.outputs = Product(l.weights, in)
	l.outputs.Add(l.biases)
	if l.activation != nil {
		l.outputs.Map(l.activation)
	} else {
		l.outputs = SoftMax(l.outputs)
	}
}

// Pass in errors and learning rate to adjust weights and biases
func (l *Layer) Adjust(out_err Matrix, learn_rate float64) {
	weights_t := Transpose(l.weights)
	errors := Product(weights_t,out_err)
	gradients := NewMatrix(0,0)
	if l.derivative != nil {
		gradients = MatrixMap(l.outputs,l.derivative)
	} else {
		gradients = SoftMaxDeriv(l.outputs)
	}
	gradients.Product(errors)
	gradients.Multiply(learn_rate)
	gradients.Map(mutate)

	transposed := Transpose(l.inputs)
	deltas := Product(gradients,transposed)
	l.weights.Add(deltas)
	l.biases.Add(deltas)
}