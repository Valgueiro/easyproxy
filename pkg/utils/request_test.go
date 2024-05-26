package utils

import (
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
	expected := a[http.CanonicalHeaderKey("bla")]
	if diff := cmp.Diff(result[http.CanonicalHeaderKey("bla")], expected); diff != "" {
		t.Error(diff)
	}

	expected = b[http.CanonicalHeaderKey("foo")]
	if diff := cmp.Diff(result[http.CanonicalHeaderKey("foo")], expected); diff != "" {
		t.Error(diff)
	}

	// it should merge when the key is on both sides
	expected = append(a[http.CanonicalHeaderKey("duplicated")], b[http.CanonicalHeaderKey("duplicated")]...)
	if diff := cmp.Diff(result[http.CanonicalHeaderKey("duplicated")], expected, sortedStrings); diff != "" {
		t.Error(diff)
	}

	// it should not add any other key then the one in both headers

	keys := make([]string, 0, len(result))
	for k := range result {
		keys = append(keys, k)
	}

	expected = make([]string, 0)
	for k := range a {
		expected = append(expected, k)
	}
	for k := range b {
		if !slices.Contains(expected, k) {
			expected = append(expected, k)
		}
	}

	if diff := cmp.Diff(keys, expected, sortedStrings); diff != "" {
		t.Error(diff)
	}
}
