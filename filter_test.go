package sequence_test

import (
	"reflect"
	"strings"
	"testing"

	"github.com/winstonpurnomo/sequence"
)

func TestFilter(t *testing.T) {
	tests := map[string]struct {
		input    []interface{}
		test     func(interface{}) bool
		expected []interface{}
	}{
		"filterEvenNumbers": {
			input:    []interface{}{1, 2, 3, 4, 5},
			test:     func(n interface{}) bool { return n.(int)%2 == 0 },
			expected: []interface{}{2, 4},
		},
		"filterStringsWithA": {
			input:    []interface{}{"apple", "banana", "avocado"},
			test:     func(s interface{}) bool { return strings.Contains(s.(string), "a") },
			expected: []interface{}{"banana", "avocado"},
		},
		"filterNonNil": {
			input:    []interface{}{1, nil, 3, nil, 5},
			test:     func(n interface{}) bool { return n != nil },
			expected: []interface{}{1, 3, 5},
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			result := sequence.Filter(tt.input, tt.test)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("got %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestFirst(t *testing.T) {
	tests := map[string]struct {
		input    []interface{}
		test     func(interface{}) bool
		expected interface{}
	}{
		"firstEvenNumber": {
			input:    []interface{}{1, 3, 4, 6},
			test:     func(n interface{}) bool { return n.(int)%2 == 0 },
			expected: 4,
		},
		"firstStringWithB": {
			input:    []interface{}{"apple", "banana", "avocado"},
			test:     func(s interface{}) bool { return strings.Contains(s.(string), "b") },
			expected: "banana",
		},
		"firstNonNil": {
			input:    []interface{}{nil, nil, 3, 5},
			test:     func(n interface{}) bool { return n != nil },
			expected: 3,
		},
		"noMatch": {
			input:    []interface{}{1, 3, 5},
			test:     func(n interface{}) bool { return n.(int) > 5 },
			expected: nil,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			result := sequence.First(tt.input, tt.test)
			if result != tt.expected {
				t.Errorf("got %v, want %v", result, tt.expected)
			}
		})
	}
}
