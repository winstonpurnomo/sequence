package sequence_test

import (
	"errors"
	"fmt"
	"strings"
	"testing"

	"github.com/winstonpurnomo/sequence"
)

func TestMap(t *testing.T) {
	tests := map[string]struct {
		input    []interface{}
		fn       func(interface{}) interface{}
		expected []interface{}
	}{
		"doubleInts": {
			input:    []interface{}{1, 2, 3},
			fn:       func(n interface{}) interface{} { return n.(int) * 2 },
			expected: []interface{}{2, 4, 6},
		},
		"toUpperStrings": {
			input:    []interface{}{"a", "b", "c"},
			fn:       func(s interface{}) interface{} { return strings.ToUpper(s.(string)) },
			expected: []interface{}{"A", "B", "C"},
		},
		"identity": {
			input:    []interface{}{1, "a", 3.14},
			fn:       func(n interface{}) interface{} { return n },
			expected: []interface{}{1, "a", 3.14},
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			result := sequence.Map(tt.input, tt.fn)
			for i, v := range result {
				if v != tt.expected[i] {
					t.Errorf("got %v, want %v", result, tt.expected)
					break
				}
			}
		})
	}
}

func TestTryMap(t *testing.T) {
	tests := map[string]struct {
		input    []interface{}
		fn       func(interface{}) (interface{}, error)
		expected []interface{}
		err      error
	}{
		"doubleInts": {
			input:    []interface{}{1, 2, 3},
			fn:       func(n interface{}) (interface{}, error) { return n.(int) * 2, nil },
			expected: []interface{}{2, 4, 6},
			err:      nil,
		},
		"squareIntsWithError": {
			input:    []interface{}{1, 2, 3},
			fn:       func(n interface{}) (interface{}, error) { return 0, errors.New("test error") },
			expected: nil,
			err:      fmt.Errorf("tryMap(): %w", errors.New("test error")),
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			result, err := sequence.TryMap(tt.input, tt.fn)
			if !errors.Is(err, tt.err) {
				t.Errorf("got error %v, want %v", err, tt.err)
			}
			if tt.err == nil {
				for i, v := range result {
					if v != tt.expected[i] {
						t.Errorf("got %v, want %v", result, tt.expected)
						break
					}
				}
			}
		})
	}
}

func TestCollectMap(t *testing.T) {
	tests := map[string]struct {
		input    []interface{}
		fn       func(interface{}) (interface{}, error)
		expected []interface{}
		errs     []error
	}{
		"doubleInts": {
			input:    []interface{}{1, 2, 3},
			fn:       func(n interface{}) (interface{}, error) { return n.(int) * 2, nil },
			expected: []interface{}{2, 4, 6},
			errs:     nil,
		},
		"squareIntsWithMixedErrors": {
			input: []interface{}{1, 2, 3},
			fn: func(n interface{}) (interface{}, error) {
				if n == 2 {
					return 0, errors.New("test error")
				}
				return n.(int) * n.(int), nil
			},
			expected: []interface{}{1, 9},
			errs:     []error{errors.New("test error")},
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			result, errs := sequence.CollectMap(tt.input, tt.fn)
			for i, v := range result {
				if v != tt.expected[i] {
					t.Errorf("got %v, want %v", result, tt.expected)
					break
				}
			}
			if len(errs) != len(tt.errs) {
				t.Errorf("got errors %v, want %v", errs, tt.errs)
			}
		})
	}
}

func TestCompactMap(t *testing.T) {
	tests := map[string]struct {
		input    []interface{}
		fn       func(interface{}) *interface{}
		expected []interface{}
	}{
		"nonNilResults": {
			input: []interface{}{1, 2, nil, 3},
			fn: func(n interface{}) *interface{} {
				if n == nil {
					return nil
				}
				return &n
			},
			expected: []interface{}{1, 2, 3},
		},
		"allNilResults": {
			input:    []interface{}{nil, nil, nil},
			fn:       func(n interface{}) *interface{} { return nil },
			expected: []interface{}{},
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			result := sequence.CompactMap(tt.input, tt.fn)
			for i, v := range result {
				if v != tt.expected[i] {
					t.Errorf("got %v, want %v", result, tt.expected)
					break
				}
			}
		})
	}
}

func TestReduce(t *testing.T) {
	tests := map[string]struct {
		input    []interface{}
		into     interface{}
		reducer  func(interface{}, interface{}) interface{}
		expected interface{}
	}{
		"sumInts": {
			input:    []interface{}{1, 2, 3},
			into:     0,
			reducer:  func(acc, n interface{}) interface{} { return acc.(int) + n.(int) },
			expected: 6,
		},
		"concatStrings": {
			input:    []interface{}{"a", "b", "c"},
			into:     "",
			reducer:  func(acc, s interface{}) interface{} { return acc.(string) + s.(string) },
			expected: "abc",
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			result := sequence.Reduce(tt.input, tt.into, tt.reducer)
			if result != tt.expected {
				t.Errorf("got %v, want %v", result, tt.expected)
			}
		})
	}
}
