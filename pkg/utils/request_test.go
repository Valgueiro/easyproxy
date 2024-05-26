package utils

import (
	"fmt"
	"net/http"
	"slices"
	"sort"
	"testing"

	"github.com/google/go-cmp/cmp"
)

var sortedStrings cmp.Option = cmp.Transformer("Sort", func(in []string) []string {
	out := append([]string(nil), in...) // Copy input to avoid mutating it
	sort.Strings(out)
	return out
})

func Test_MergeHeaders(t *testing.T) {
	a := http.Header{
		"bla":        []string{"bla"},
		"duplicated": []string{"a"},
	}

	b := http.Header{
		"foo":        []string{"bar"},
		"duplicated": []string{"b"},
	}

	result := MergeHeaders(a, b)

	// When a key is only in one side, it should get the value from that side
	if diff := cmp.Diff(result.Get("bla"), a.Get("bla")); diff != "" {
		t.Error(diff)
	}

	if diff := cmp.Diff(result.Get("foo"), b.Get("foo")); diff != "" {
		t.Error(diff)
	}

	// it should merge when the key is on both sides
	expected := fmt.Sprintf("%s%s", a.Get("duplicated"), b.Get("duplicated"))
	if diff := cmp.Diff(result.Get("duplicated"), expected, sortedStrings); diff != "" {
		t.Error(diff)
	}

	// it should not add any other key then the one in both headers

	keys := make([]string, 0, len(result))
	for k := range result {
		keys = append(keys, k)
	}

	expected2 := make([]string, 0)
	for k := range a {
		expected2 = append(expected2, k)
	}
	for k := range b {
		if !slices.Contains(expected2, k) {
			expected2 = append(expected2, k)
		}
	}

	if diff := cmp.Diff(keys, expected2, sortedStrings); diff != "" {
		t.Error(diff)
	}
}
