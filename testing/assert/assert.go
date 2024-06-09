package assert

import (
	"reflect"
	"testing"
)

// True ensures that the passed in boolean evaluates to true and immediately fails the test
// otherwise. This accepts variadic arguments that will be immediately passed through to the
// underlying call to `t.Fatal(args...)`
func True(t *testing.T, b bool, args ...any) {
	t.Helper()
	if !b {
		t.Fatal(args...)
	}
}

// False ensures that the passed in boolean evaluates to false and immediately fails the test
// otherwise. This accepts variadic arguments that will be immediately passed through to the
// underlying call to `t.Fatal(args...)`
func False(t *testing.T, b bool, args ...any) {
	t.Helper()
	if b {
		t.Fatal(args...)
	}
}

// Equals ensures that hte actual and expected values are equal and immediately fails the test
// otherwise.
//
// NOTE: This requires the inputs to be comparable. If you need deep comparison of nested fields,
// see `DeepEquals`.
func Equals[T comparable](t *testing.T, act, exp T) {
	t.Helper()
	if act != exp {
		t.Fatalf("Expected:\n\"\"\"\n%v\n\"\"\"\nReceived:\n\"\"\"\n%v\n\"\"\"", exp, act)
	}
}

// DeepEquals immediately fails the currently running test if the actual and expected values are not
// _deeply_ equal (using `reflect.DeepEqual`).
func DeepEquals(t *testing.T, act, exp any) {
	t.Helper()
	if !reflect.DeepEqual(act, exp) {
		t.Fatalf("Expected:\n\"\"\"\n%v\n\"\"\"\nReceived:\n\"\"\"\n%v\n\"\"\"", exp, act)
	}
}

// NotEquals ensures that the actual and expected values are _not_ equal and immediately fails the
// test otherwise.
func NotEquals[T comparable](t *testing.T, act, not_exp T) {
	t.Helper()
	if act == not_exp {
		t.Fatalf("Not expected:\n\"\"\"\n%v\n\"\"\"\nReceived:\n\"\"\"\n%v\n\"\"\"", not_exp, act)
	}
}

// Ok ensures that there is _no_ error and immediately fails the test otherwise.
func Ok(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatal(err)
	}
}

// Err ensures that there _is_ an error and immediately fails the test otherwise.
func Err(t *testing.T, err error) {
	t.Helper()
	if err == nil {
		t.Fatal("Expected error; received nil")
	}
}

func Panic(t *testing.T, fn func()) {
	t.Helper()
	defer func() {
		if r := recover(); r == nil {
			t.Fatal("Expected to panic; no panic occurred")
		}
	}()
	fn()
}
