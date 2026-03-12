package decorator

import (
	"bytes"
	"go/format"
	"testing"
)

// TestGenericTypeAliases verifies the correct decoration and round-trip restoration
// of Go 1.24 Generic Type Aliases (e.g., type A[T any] = B[T]).
//
// This validates that the combination of TypeParams and Assign (aliasing) co-existing
// on a TypeSpec does not cause decoration conflicts or restoration ordering issues.
func TestGenericTypeAliases(t *testing.T) {
	tests := []struct {
		name string
		code string
	}{
		{
			name: "basic-generic-alias",
			code: `package main

type Vector[T any] = []T`,
		},
		{
			name: "multiple-type-params",
			code: `package main

type MapAlias[K comparable, V any] = map[K]V`,
		},
		{
			name: "complex-constraints",
			code: `package main

type Number[T int | float64] = T`,
		},
		{
			name: "decorated-type-params",
			// Verifies comments attached to TypeParams do not bleed into the Assign token
			code: `package main

type List /*a*/ [T any] = /*b*/ /*c*/ []T`,
		},
		{
			name: "decorated-assignment",
			// Verifies strict separation between the generic list and the right-hand side type
			code: `package main

// Start
type Alias /*1*/ [T any] /*2*/ = /*3*/ Target[T] // End`,
		},
		{
			name: "multiline-generic-alias",
			code: `package main

type BigAlias[
	T any,
	U comparable,
] = struct {
	a T
	b U
}`,
		},
		{
			name: "interface-constraint-alias",
			code: `package main

type Stringer = interface {
	String() string
}

type Ptr[T Stringer] = *T`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Ensure source is formatted for consistent comparison
			expectBytes, err := format.Source([]byte(test.code))
			if err != nil {
				t.Fatalf("format.Source error: %v", err)
			}
			expected := string(expectBytes)

			// Parse using dst decorator
			f, err := Parse(expected)
			if err != nil {
				t.Fatalf("decorator.Parse error: %v", err)
			}

			// Restore to AST and print
			// We use a fresh restorer to ensure no decoration state leakage
			r := NewRestorer()
			buf := &bytes.Buffer{}
			if err := r.Fprint(buf, f); err != nil {
				t.Fatalf("restorer.Fprint error: %v", err)
			}

			got := buf.String()

			// Normalize whitespace for robust comparison
			if normalize(got) != normalize(expected) {
				t.Errorf("Round-trip mismatch.\nExpected:\n%s\nGot:\n%s\nDiff:\n%s", expected, got, diff(normalize(expected), normalize(got)))
			}
		})
	}
}
