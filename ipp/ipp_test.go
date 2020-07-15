package ipp

import "testing"
import "fmt"

func TestIpp(t *testing.T) {

	fmt.Printf("MKL version:%s\n", Version())

	data := make([]float32, 10)
	Zero(data)
	AddC_I(data, 1.2)

	fmt.Println(data)
	n := AddC(data, 1)
	n = AddC(n, 1)
}
