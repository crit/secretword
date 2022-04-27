package secretword

import "strings"

// Current game constants
const (
	CharLimit = 5
	TurnLimit = 6
)

// Status is the state of each letter in a guess.
type Status int

const (
	Correct Status = iota
	Incorrect
	Elsewhere
)

// Result is all the letter statuses in a guess.
type Result []Status

// Is checks to see if one result equals another result.
func (r Result) Is(other Result) bool {
	if len(r) != len(other) {
		return false
	}

	for i, status := range r {
		if status != other[i] {
			return false
		}
	}

	return true
}

// Basic result sets for easy comparison.
var (
	Success = Result{Correct, Correct, Correct, Correct, Correct}
	Failure = Result{Incorrect, Incorrect, Incorrect, Incorrect, Incorrect}
)

// Compare calculates the Result of a guess against an answer.
func Compare(answer, guess string) (res Result) {
	if len(answer) != CharLimit || len(guess) != CharLimit {
		return Failure
	}

	if answer == guess {
		return Success
	}

	for i := 0; i < CharLimit; i++ {
		res = append(res, check(i, answer, guess))
	}

	return res
}

// check is a helper function that compares the letter at a specific position in the answer and guess.
func check(position int, answer, guess string) Status {
	a, g := answer[position], guess[position]

	if a == g {
		return Correct
	}

	if strings.Contains(answer, string(g)) {
		return Elsewhere
	}

	return Incorrect
}
