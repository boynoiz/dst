package decorator

import (
	"bytes"
	"go/format"
	"testing"
)

// TestRangeIterators verifies that the decorator correctly handles Go 1.23+ range iterator syntax,
// ensuring comments and whitespace around the iterator expression are preserved during round-trips.
func TestRangeIterators(t *testing.T) {
	tests := []struct {
		name string
		code string
	}{
		{
			name: "basic-func-iterator",
			code: `package main

func main() {
	for x := range seq {
		print(x)
	}
}`,
		},
		{
			name: "iterator-call",
			code: `package main

func main() {
	for k, v := range slices.All(s) {
		print(k, v)
	}
}`,
		},
		{
			name: "integer-range-go1.22",
			code: `package main

func main() {
	for i := range 10 {
		print(i)
	}
}`,
		},
		{
			name: "complex-inline-comments",
			code: `package main

func main() {
	// Start loop
	for /*a*/ i /*b*/, /*c*/ v := /*d*/ range /*e*/ iter( /*f*/ ) /*g*/ {
		print(i, v)
	}
}`,
		},
		{
			name: "multiline-iterator-call",
			code: `package main

func main() {
	for i := range merge(
		seq1,
		seq2,
	) {
		print(i)
	}
}`,
		},
		{
			name: "iterator-with-block-spacing",
			code: `package main

func main() {
	for x := range seq {

		print(x)

	}
}`,
		},
		{
			name: "range-decorations-iter",
			// Test explicit attachment points for RangeStmt with an iterator expression
			// The X expression here is a CallExpr
			code: `package main

func main() {
	for i := range /*Range*/ slices.All(s) /*X*/ {
	}
}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// standard formatting ensures consistent whitespace for comparison
			formatted, err := format.Source([]byte(test.code))
			if err != nil {
				t.Fatalf("format.Source failed: %s", err)
			}
			expectedCode := string(formatted)

			// Parse the code using the decorator
			f, err := Parse(expectedCode)
			if err != nil {
				t.Fatalf("Parse failed: %s", err)
			}

			// Restore the file to AST and print it
			// We use a clean FileRestorer to simulate a full round-trip
			r := NewRestorer()
			buf := &bytes.Buffer{}
			if err := r.Fprint(buf, f); err != nil {
				t.Fatalf("Fprint failed: %s", err)
			}

			// Normalize strings to ignore insignificant whitespace differences if necessary,
			// though typical round-trip should be exact for formatted code.
			got := buf.String()
			
			// Use the diff helper assumed to be present in the package from other tests
			// (e.g. wrapper around diffmatchpatch as seen in existing tests like decorator_fragment_test.go)
			if normalize(got) != normalize(expectedCode) {
				t.Errorf("Mismatch in round-trip.\nExpected:\n%s\nGot:\n%s\nDiff:\n%s", expectedCode, got, diff(normalize(expectedCode), normalize(got)))
			}
		})
	}
}