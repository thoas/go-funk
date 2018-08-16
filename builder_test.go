package funk

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChain(t *testing.T) {
	testCases := []struct {
		In    interface{}
		Panic string
	}{
		// Check with array types
		{In: []int{0, 1, 2}},
		{In: []string{"aaa", "bbb", "ccc"}},
		{In: []interface{}{0, false, "___"}},

		// Check with map types
		{In: map[int]string{0: "aaa", 1: "bbb", 2: "ccc"}},
		{In: map[string]string{"0": "aaa", "1": "bbb", "2": "ccc"}},
		{In: map[int]interface{}{0: 0, 1: false, 2: "___"}},

		// Check with invalid types
		{false, "Type bool is not supported by Chain"},
		{0, "Type int is not supported by Chain"},
		{"nope", "Type string is not supported by Chain"},
	}

	for idx, tc := range testCases {
		t.Run(fmt.Sprintf("test case #%d", idx+1), func(t *testing.T) {
			is := assert.New(t)

			if tc.Panic != "" {
				is.PanicsWithValue(tc.Panic, func() {
					Chain(tc.In)
				})
				return
			}

			chain := Chain(tc.In)
			collection := chain.(*chainBuilder).collection

			is.Equal(collection, tc.In)
		})
	}
}

func TestLazyChain(t *testing.T) {
	testCases := []struct {
		In    interface{}
		Panic string
	}{
		// Check with array types
		{In: []int{0, 1, 2}},
		{In: []string{"aaa", "bbb", "ccc"}},
		{In: []interface{}{0, false, "___"}},

		// Check with map types
		{In: map[int]string{0: "aaa", 1: "bbb", 2: "ccc"}},
		{In: map[string]string{"0": "aaa", "1": "bbb", "2": "ccc"}},
		{In: map[int]interface{}{0: 0, 1: false, 2: "___"}},

		// Check with invalid types
		{false, "Type bool is not supported by LazyChain"},
		{0, "Type int is not supported by LazyChain"},
		{"nope", "Type string is not supported by LazyChain"},
	}

	for idx, tc := range testCases {
		t.Run(fmt.Sprintf("test case #%d", idx+1), func(t *testing.T) {
			is := assert.New(t)

			if tc.Panic != "" {
				is.PanicsWithValue(tc.Panic, func() {
					LazyChain(tc.In)
				})
				return
			}

			chain := LazyChain(tc.In)
			collection := chain.(*lazyBuilder).exec()

			is.Equal(collection, tc.In)
		})
	}
}

func TestLazyChainWith(t *testing.T) {
	testCases := []struct {
		In    func() interface{}
		Panic string
	}{
		// Check with array types
		{In: func() interface{} { return []int{0, 1, 2} }},
		{In: func() interface{} { return []string{"aaa", "bbb", "ccc"} }},
		{In: func() interface{} { return []interface{}{0, false, "___"} }},

		// Check with map types
		{In: func() interface{} { return map[int]string{0: "aaa", 1: "bbb", 2: "ccc"} }},
		{In: func() interface{} { return map[string]string{"0": "aaa", "1": "bbb", "2": "ccc"} }},
		{In: func() interface{} { return map[int]interface{}{0: 0, 1: false, 2: "___"} }},

		// Check with invalid types
		{
			In:    func() interface{} { return false },
			Panic: "Type bool is not supported by LazyChainWith generator"},
		{
			In:    func() interface{} { return 0 },
			Panic: "Type int is not supported by LazyChainWith generator"},
		{
			In:    func() interface{} { return "nope" },
			Panic: "Type string is not supported by LazyChainWith generator"},
	}

	for idx, tc := range testCases {
		t.Run(fmt.Sprintf("test case #%d", idx+1), func(t *testing.T) {
			is := assert.New(t)

			if tc.Panic != "" {
				is.PanicsWithValue(tc.Panic, func() {
					LazyChainWith(tc.In).(*lazyBuilder).exec()
				})
				return
			}

			chain := LazyChainWith(tc.In)
			collection := chain.(*lazyBuilder).exec()

			is.Equal(collection, tc.In())
		})
	}
}
