# Secret Word CLI Game

![game screenshot](http://i.critrussell.net/uvkRh3gYBcZayK8.png)

## Build

`cd cmd/wordle && go build .`

## Installation

`cd cmd/wordle && go install .`

## Dependencies

- [github.com/fatih/color]()
- [github.com/manifoldco/promptui]()

# Wordle Solver

In development

# Wordle Library


Installation: `go get -u github.com/crit/wordle-go/pkg/wordle`

Usage:

```go
package main

import (
	"fmt"

	"github.com/crit/wordle-go/pkg/wordle"
)

func main() {
	// compare answer against a guess to get a result set for all characters.
	result := wordle.Compare("peach", "party") // wordle.Result{Correct, Elsewhere, Incorrect, Incorrect, Incorrect}
	fmt.Println(result)
}
```
