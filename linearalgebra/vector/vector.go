package vector

import (
	"github.com/marti700/veritas/commons"
)

type Vector struct {
	Size int
	Data []float64
}

func NewVector(v []float64) Vector {
	return Vector{
		Size: len(v),
		Data: v,
	}
}

// calculates the dot produt between two vectors
func (v Vector) DotProduct(v1 Vector) float64 {
	var sumbuffer []float64

	for i := range v.Data {
		sumbuffer = append(sumbuffer, v.Data[i]*v1.Data[i])
	}

	return commons.Sum(sumbuffer)
}

// Creates a vector of the specified lenght which contents are ones
func Ones(length int) Vector {
	temp := make([]float64, length)
	for i := 0; i < length; i++ {
		temp[i] = 1
	}
	return NewVector(temp)
}
