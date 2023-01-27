package a

import (
	"testing"
)

func Equals(t *testing.T, a, b interface{}) {
	if a != b {
		t.Errorf("expected %v to equal %v", a, b)
	}
}

func NotEquals(t *testing.T, a, b interface{}) {
	if a == b {
		t.Errorf("expected %v to not equal %v", a, b)
	}
}

func Throws(t *testing.T, f func()) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("expected function to throw")
		}
	}()

	f()
}

func ShouldFailTest(t *testing.T, f func(mockTest *testing.T)) {
	mockTest := new(testing.T)
	f(mockTest)

	if !mockTest.Failed() {
		t.Error("expected a failure, but test exited successfully")
	}
}
