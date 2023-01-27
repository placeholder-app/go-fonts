package a

import (
	"testing"
)

type TestCase struct {
	a, b interface{}
}

var equalsCases = []TestCase{
	{1, 1},
	{"a", "a"},
	{true, true},
	{false, false},
	{nil, nil},
}

var notEqualsCases = []TestCase{
	{1, 2},
	{"a", "b"},
	{true, false},
	{false, true},
	{nil, 1},
}

func TestEquals(t *testing.T) {
	for _, c := range equalsCases {
		Equals(t, c.a, c.b)
	}
}

func TestNotEquals(t *testing.T) {
	for _, c := range notEqualsCases {
		NotEquals(t, c.a, c.b)
	}
}

func TestShouldFailTest(t *testing.T) {
	ShouldFailTest(t, func(mockTest *testing.T) {
		Equals(mockTest, 1, 2)
	})
}

func TestShouldFailTestFail(t *testing.T) {
	ShouldFailTest(t, func(mockTest *testing.T) {
		ShouldFailTest(mockTest, func(subMockTest *testing.T) {
			Equals(subMockTest, 1, 1)
		})
	})
}

func TestEqualsFail(t *testing.T) {
	for _, c := range notEqualsCases {
		ShouldFailTest(t, func(mockTest *testing.T) {
			Equals(mockTest, c.a, c.b)
		})
	}
}

func TestNotEqualsFail(t *testing.T) {
	for _, c := range equalsCases {
		ShouldFailTest(t, func(mockTest *testing.T) {
			NotEquals(mockTest, c.a, c.b)
		})
	}
}

func TestThrows(t *testing.T) {
	Throws(t, func() {
		panic("test")
	})
}

func TestThrowsFail(t *testing.T) {
	ShouldFailTest(t, func(mockTest *testing.T) {
		Throws(mockTest, func() {})
	})
}
