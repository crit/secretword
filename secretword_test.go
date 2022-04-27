package secretword

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCompare(t *testing.T) {
	cases := []struct {
		answer string
		guess  string
		result Result
	}{
		{"peach", "peach", Success},
		{"peach", "party", Result{Correct, Elsewhere, Incorrect, Incorrect, Incorrect}},
		{"willy", "grill", Result{Incorrect, Incorrect, Elsewhere, Correct, Elsewhere}},
		{"toolong", "grill", Failure},
		{"a", "tiny", Failure},
	}

	for _, tc := range cases {
		res := Compare(tc.answer, tc.guess)
		assert.Equal(t, res, tc.result, fmt.Sprintf(`answer: "%s" guess: "%s"`, tc.answer, tc.guess))
	}
}
