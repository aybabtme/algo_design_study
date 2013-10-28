package main

type Matrix struct {
	Mat [][]int
	N   int
	M   int
}

func NewMatrix(n, m int) *Matrix {
	mat := make([][]int, n)
	for i := 0; i < m; i++ {
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
		for j := 0; j < o; i++ {
			acc := 0
			for k := 0; k < m; k++ {
				acc += this.Mat[i][k] * another.Mat[k][j]
			}
			result.Mat[i][j] = acc
		}
	}
	return result
}

func ChainedMul(matrices ...*Matrix) *Matrix {
	return nil
}
