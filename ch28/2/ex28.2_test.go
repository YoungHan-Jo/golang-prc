package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFibonacci1(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(0, fibonacci1(-1), "fibonacci1(-1) should be 0")
	assert.Equal(0, fibonacci1(0), "fibonacci1(0) should be 0")
	assert.Equal(1, fibonacci1(1), "fibonacci1(1) should be 1")
	assert.Equal(2, fibonacci1(3), "fibonacci1(3) should be 2")
	assert.Equal(233, fibonacci1(13), "fibonacci1(13) should be 233")
}

func TestFibonacci2(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(0, fibonacci2(-1), "fibonacci2(-1) should be 0")
	assert.Equal(0, fibonacci2(0), "fibonacci2(0) should be 0")
	assert.Equal(1, fibonacci2(1), "fibonacci2(1) should be 1")
	assert.Equal(2, fibonacci2(3), "fibonacci2(3) should be 2")
	assert.Equal(233, fibonacci2(13), "fibonacci2(13) should be 233")
}

func BenchmarkFibonacci1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fibonacci1(20)
	}
}

func BenchmarkFibonacci2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fibonacci2(20)
	}
}

func TestAtoi1(t *testing.T) {
	n, err := Atoi("1")
	if err != nil {
		t.Fail()
	}
	if n != 1 {
		t.Fail()
	}
}

func TestAtoi4(t *testing.T) {
	n, err := Atoi("    1234")
	if err != nil {
		t.Fail()
	}
	if n != 1234 {
		t.Fail()
	}
}

func TestAtoi5(t *testing.T) {
	n, err := Atoi("- 1234")
	if err != nil {
		t.Fail()
	}
	if n != -1234 {
		t.Fail()
	}
}
