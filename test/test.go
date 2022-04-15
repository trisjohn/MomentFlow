package main

import (
	"fmt"
	"gui/neural"
)

func main() {
	fmt.Println("BEGIN")

	n := neural.NewNeuralNetwork(2,2,1, 0.23)
	train_data := neural.TrainingData_xor()
	for i := 0; i < 1000; i++ {
		v := neural.RandomItem(train_data)
		n.Train(v.Inputs, v.Targets)
	}

	fmt.Println(0, n.FeedForward(neural.NewFArr(0,0)))
	fmt.Println(1, n.FeedForward(neural.NewFArr(0,1)))
	fmt.Println(1, n.FeedForward(neural.NewFArr(1,0)))
	fmt.Println(0, n.FeedForward(neural.NewFArr(1,1)))
}