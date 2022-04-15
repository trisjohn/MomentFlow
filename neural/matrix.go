package neural

import (
	"fmt"
	"math"
	"math/rand"
)

type Matrix struct {
	rows, cols int
	values []float64
}

func (m Matrix) String() string {
	s := fmt.Sprintf("Matrix %vx%v : [  \n", m.rows, m.cols)
	for i := 0; i < m.rows; i++ {
		r := "[ "
		for j := 0; j < m.cols; j++ {
			sVal := fmt.Sprintf("%v, ", m.At(i,j))
			r += sVal
		}
		s += r + "]\n"
	}
	s += "   ]\n"
	return s
}

func(m *Matrix) matrixEqualDim(other Matrix) bool {
	return other.rows == m.rows && other.cols == m.cols
}

func matrixError(args... interface{}) {
	msg := "Error within Matrix struct: "
	for _,v := range args {
		msg += fmt.Sprintf("%v ", v)
	}
	panic(msg)
}

func NewMatrix(rows, cols int) (m Matrix) {
	m = Matrix{}
	m.rows = rows
	m.cols = cols
	total := rows * cols
	m.values = make([]float64, total)
	return
}

func MatrixFromArray(arr []float64) (m Matrix) {
	m = NewMatrix(len(arr),1)
	m.values = arr
	return
}

func (m *Matrix) At(row,col int) float64 {
	if row >= m.rows || col >= m.cols {
		matrixError("Values (", m.values, ") out of bounds",row, "rows=", m.rows,col, "cols=", m.cols)
		return 0.0
	}
	return m.values[col+row*m.cols]
}

func (m *Matrix) Map(f func(float64) float64) {
	for i,v := range m.values {
		m.values[i] = f(v)
	}
}
func MatrixMap(m Matrix,f func(float64) float64) Matrix {
	for i,v := range m.values {
		m.values[i] = f(v)
	}
	return m
}

func Transpose(m Matrix ) (result Matrix) {
	result = NewMatrix(m.cols, m.rows)
	for i := 0; i < len(m.values); i++ {
		result.values[i] = m.At(i/m.cols,i%m.cols)
	}
	return
}

func (m *Matrix) ToArray() (result []float64) {
	return m.values
}

// Returns true only if each element is exactly equal to the other's
func (m *Matrix) Equals(other Matrix) bool {
	if !m.matrixEqualDim(other) {
		return false
	}
	for i := 0; i < m.rows; i++ {
		for j := 0; j < m.cols; j++ {
			if m.At(i,j) != other.At(i,j) {
				return false
			}
		}
	}
	return true
}

// Pass in an integer to fill each element in the array
func (m *Matrix) FillInt(n int) {
	for i := range m.values {
		m.values[i] = float64(n)
	}
}

// Pass in an float64 to fill each element in the array
func (m *Matrix) FillFloat(n float64) {
	for i := range m.values {
		m.values[i] = n
	}
}

// Pass in an int for a base, else leave empty to randomize between 0 and 1
func (m *Matrix) Randomize(args... int) {
	t := m.rows * m.cols
	if len(args) == 0 {
		for i := 0; i < t; i++ {
			m.values[i] = 1 - 2 * rand.Float64()
		}
	} else {
		for i := 0; i < t; i++ {
			m.values[i] = math.RoundToEven(rand.Float64() * float64(args[0]))
		}
	}
}

// Return a new matrix that is a - b
func Subtract(a,b Matrix) (result Matrix) {
	if !a.matrixEqualDim(b) {
		matrixError("A's Dimensions do not match b's", a,b)
	}
	result = NewMatrix(a.rows, a.cols)
	for i := range result.values {
		result.values[i] = a.values[i] - b.values[i]
	}
	return
}


// Return a new matrix that is the dot product of a and b, a cols # must equal b rows #
func Product(a,b Matrix) (result Matrix) {
	// fmt.Println("Multiplying: ", a , " BY ", b)
	if a.cols != b.rows {
		matrixError("a's # of cols must match b's # of rows for dot product",a.cols, b.rows)
	}

	result = NewMatrix(a.rows, b.cols)
	
	for i := range result.values {
		sum := 0.0
		for j := 0; j < a.cols; j++ {
			x := a.At(i/result.cols,j)
			y := b.values[i/result.rows+j*b.cols]
			// fmt.Println(i,":",j,"->",x,"x",y,x*y)
			sum += x * y
		}
		// fmt.Println("result=",sum)
		result.values[i] = sum
		// fmt.Println(i, " >> ", result)
	}
	return
}

// Scalar operation
func (m *Matrix) Multiply(n float64) {
	for i := range m.values {
		m.values[i] *= n
	}
}

// In place, multiply m element-wise by other, replacing m with the new product
func (m *Matrix) Product(other Matrix) {
	if !m.matrixEqualDim(other) {
		matrixError("Matrix dimensions do not match for elementwise product.",m, other)
	}
	for i,v := range m.values {
		m.values[i] = v * other.values[i]
	}
}

func (m *Matrix) Add(n interface{}) {
	switch k := n.(type) {
	case float64:
		for i,v := range m.values {
			m.values[i] = v + k
		}
	case Matrix:
		// Element-wise
		if !k.matrixEqualDim(k) {
			matrixError("Error when adding matrices, sizes do not match:",m,k)
		}
		for i,v := range k.values {
			m.values[i] += v
		}
	}
}
