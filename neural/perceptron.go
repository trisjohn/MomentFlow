package neural

import "math/rand"

func sign(n float64) int {
	if n >= 0{
		return 1
	} else {
		return -1
	}
}

type Perceptron struct {
	weights []float64
	bias, lr float64
}

func (p* Perceptron) Guess(inputs []float64) int {
	sum := 0.0
	for i,v := range p.weights {
		sum += inputs[i] * v
	}
	output := sign(sum)
	return output
}

func (p* Perceptron) GuessY(x float64) float64 {
	w0,w1,w2 := p.weights[0], p.weights[1], p.weights[2]
	return -(w2/w1)*p.bias - (w0/w1) * x
}

func(p* Perceptron) Train(inputs []float64, target int) {
	guess := p.Guess(inputs)
	err := guess - target

	for i := range p.weights {
		p.weights[i] = float64(err) * inputs[i] * p.lr
	}
}

func NewPerceptron(n int, learning_rate float64) Perceptron{
	p := Perceptron{}
	p.weights = make([]float64, 2)
	p.lr = learning_rate
	p.bias = 1
	for i := range p.weights{
		p.weights[i] = rand.Float64() * 2 - 1
	}
	
	return p
}