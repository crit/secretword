package secretword

// HistoryPrinter is the interface used to print the results contained in History.
type HistoryPrinter func(Result, string)

// History contains all the guesses and results for a round of wordle.
type History struct {
	guesses []string
	results []Result
}

// Append adds a guess and result to the history.
func (h *History) Append(res Result, guess string) {
	h.guesses = append(h.guesses, guess)
	h.results = append(h.results, res)
}

// Print takes a HistoryPrinter and outputs the recorded guesses and results.
func (h *History) Print(printer HistoryPrinter) {
	for i, guess := range h.guesses {
		printer(h.results[i], guess)
	}
}
