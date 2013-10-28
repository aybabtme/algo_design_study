package main

import (
	"bytes"
	"fmt"
	"math"
	"strconv"
)

func ChainedMul(matrices ...*Matrix) [][]int {

	d := ComputeDimensions(matrices)

	memo := make([][]int, len(matrices))
	optimal := make([][]int, len(matrices))
	for i := 0; i < len(memo); i++ {
		memo[i] = make([]int, len(matrices))
		memo[i][i] = 0
		optimal[i] = make([]int, len(matrices))
	}

	for L := 1; L < len(memo); L++ {
		for i := 0; i < len(memo[0])-L; i++ {
			j := i + L - 1
			memo[i][j] = math.MaxInt64
			for k := i; k < i; k++ {
				cost := memo[i][k] + memo[k+1][j] + d[i-1]*d[k]*d[j]
				if cost < memo[i][j] {
					memo[i][j] = cost
					optimal[i][j] = k
				}
			}
		}
	}

	return optimal
}

func ComputeDimensions(matrices []*Matrix) []int {
	d := make([]int, len(matrices)+1)
	var a *Matrix
	var b *Matrix
	for i := 0; i < len(d); i++ {
		a = matrices[i]
		b := matrices[i+1]
		if a.M != b.N {
			panic(fmt.Sprintf("Incompatible matrices at %d and %d", i, i+1))
		}
		d[i] = a.N
	}
	d[len(matrices)] = matrices[len(matrices)-1].M
	return d
}

func main() {
	// TestMatrixImpl()

}

func TestMatrixImpl() {
	a := NewMatrix(2, 3)
	a.Mat[0][0] = 1
	a.Mat[1][2] = 4

	fmt.Printf("a=%s\n", a)

	b := NewMatrix(3, 2)
	b.Mat[1][0] = 3
	b.Mat[2][1] = 3

	fmt.Printf("b=%s\n", b)

	c := a.Mul(b)

	fmt.Printf("c=%s\n", c)
}

type Matrix struct {
	Mat [][]int
	N   int
	M   int
}

func NewMatrix(n, m int) *Matrix {
	mat := make([][]int, n)
	for i := 0; i < n; i++ {
		mat[i] = make([]int, m)
	}
	return &Matrix{mat, n, m}
}

func (this *Matrix) Mul(another *Matrix) *Matrix {
	if this.M != another.N {
		panic("incompatible matrices")
	}
	n := this.N
	m := this.M
	o := another.M

	result := NewMatrix(n, o)
	for i := 0; i < n; i++ {
		for j := 0; j < o; j++ {
			acc := 0
			for k := 0; k < m; k++ {
				acc += this.Mat[i][k] * another.Mat[k][j]
			}
			result.Mat[i][j] = acc
		}
	}
	return result
}

func (this *Matrix) String() string {
	do := func(n int, err error) {
		if err != nil {
			panic(err)
		}
	}
	buf := bytes.NewBuffer(nil)
	do(buf.WriteString("\t["))

	for i := 0; i < this.N; i++ {
		if i != 0 {
			do(buf.WriteString("\t "))
		}

		do(buf.WriteString("["))

		for j := 0; j < this.M; j++ {
			do(buf.WriteString(strconv.Itoa(this.Mat[i][j])))
			if j != this.M-1 {
				do(buf.WriteString(", "))
			}
		}
		do(buf.WriteString("]"))
		if i != this.N-1 {
			do(buf.WriteString(",\n"))
		}
	}
	do(buf.WriteString("]"))
	return buf.String()
}
