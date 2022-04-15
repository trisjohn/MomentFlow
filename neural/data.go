package neural

import (
	"fmt"
	"math/rand"
)

// Holds data for a nueral network
type Data struct {
	Inputs, Targets []float64
}

// Return a new float64 array filled with whatever
func NewFArr(f... float64) (arr []float64) {
	arr = append(arr, f...)
	return
}

// Grabs a random item from an array of Data
func RandomItem(a []Data) (item Data) {
	item = a[rand.Int31n(int32(len(a)))]
	return
}

// Return a new data piece to be used in a neural network
// Wants: input [] float64 and output [] float64. Output can be left blank
func NewData(data... []float64) (d Data) {
	if len(data) > 2 {
		fmt.Println("Wrong number of arguments provided to new data wants up to 2")
		return
	}
	d.Inputs = data[0]
	if len(data) > 1 {
		d.Targets = data[1]
	}
	return
}

// Return a hard coded training data set for xor
func TrainingData_xor() (datas []Data) {
	datas = append(datas, NewData(NewFArr(0,1),NewFArr(1)))
	datas = append(datas, NewData(NewFArr(1,0),NewFArr(1)))
	datas = append(datas, NewData(NewFArr(0,0),NewFArr(0)))
	datas = append(datas, NewData(NewFArr(1,1),NewFArr(0)))
	return
}