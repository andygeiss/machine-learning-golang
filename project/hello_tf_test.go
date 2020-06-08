package main_test

import (
	tf "github.com/tensorflow/tensorflow/tensorflow/go"
	"github.com/tensorflow/tensorflow/tensorflow/go/op"
	"os"
	"testing"
)

func setup() {
	_ = os.Setenv("TF_CPP_MIN_LOG_LEVEL", "3")
}

func Benchmark_TensorFlow_NewScope(b *testing.B) {
	setup()
	for i := 0; i < b.N; i++ {
		op.NewScope()
	}
}

func Benchmark_TensorFlow_Scope_Finalize(b *testing.B) {
	setup()
	s := op.NewScope()
	for i := 0; i < b.N; i++ {
		_, _ = s.Finalize()
	}
}

func Benchmark_TensorFlow_NewSession(b *testing.B) {
	setup()
	s := op.NewScope()
	c := op.Const(s, "Hello from TensorFlow version "+tf.Version())
	graph, _ := s.Finalize()

	for i := 0; i < b.N; i++ {
		sess, _ := tf.NewSession(graph, nil)
		_, _ = sess.Run(nil, []tf.Output{c}, nil)
	}
}

func Benchmark_TensorFlow_Session_Run(b *testing.B) {
	setup()
	s := op.NewScope()
	c := op.Const(s, "Hello from TensorFlow version "+tf.Version())
	graph, _ := s.Finalize()
	sess, _ := tf.NewSession(graph, nil)

	for i := 0; i < b.N; i++ {
		_, _ = sess.Run(nil, []tf.Output{c}, nil)
	}
}
