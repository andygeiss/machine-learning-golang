package main_test

import (
	tf "github.com/tensorflow/tensorflow/tensorflow/go"
	"github.com/tensorflow/tensorflow/tensorflow/go/op"
	"golang.org/x/exp/rand"
	"gonum.org/v1/gonum/mat"
	"os"
	"testing"
)

func Benchmark_TensorFlow_Mat_Sub(b *testing.B) {
	_ = os.Setenv("TF_CPP_MIN_LOG_LEVEL", "3")
	root := op.NewScope()
	X := op.Placeholder(root.SubScope("input"), tf.Float, op.PlaceholderShape(tf.MakeShape(1000, 1000)))
	Y := op.Placeholder(root.SubScope("input"), tf.Float, op.PlaceholderShape(tf.MakeShape(1000, 1000)))
	product := op.MatMul(root, X, Y)
	graph, _ := root.Finalize()
	sess, _ := tf.NewSession(graph, &tf.SessionOptions{})
	x, _ := tf.NewTensor(randomMatrixFloat64(1000, 1000))
	y, _ := tf.NewTensor(randomMatrixFloat64(1000, 1000))
	for i := 0; i < b.N; i++ {
		_, _ = sess.Run(map[tf.Output]*tf.Tensor{X: x, Y: y}, []tf.Output{product}, nil)
	}
}

func Benchmark_GoNum_Mat_Sub(b *testing.B) {
	x := mat.NewDense(1000, 1000, randomVectorFloats64(1000000))
	y := mat.NewDense(1000, 1000, randomVectorFloats64(1000000))
	z := mat.NewDense(1000, 1000, randomVectorFloats64(1000000))
	for i := 0; i < b.N; i++ {
		z.Zero()
		z.Sub(x, y)
	}
}

func randomVectorFloats64(size int) []float64 {
	out := make([]float64, size)
	for i := 0; i < size; i++ {
		out[i] = rand.Float64()
	}
	return out
}

func randomMatrixFloat64(r, c int) [][]float64 {
	out := make([][]float64, r)
	for i := 0; i < r; i++ {
		out[i] = make([]float64, c)
		for j := 0; j < c; j++ {
			out[i][j] = rand.Float64()
		}
	}
	return out
}
