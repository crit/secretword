# Secret Word CLI Game

![game screenshot](http://i.critrussell.net/uvkRh3gYBcZayK8.png)

## Build

`cd cmd/secretword && go build .`

## Installation

`cd cmd/secretword && go install .`

## Dependencies

- [github.com/fatih/color]()
- [github.com/manifoldco/promptui]()

# Secret Word Solver

In development

# Secret Word Library


Installation: `go get -u github.com/crit/secretword/pkg/wordle`

Usage:

```go
package main

import (
	"fmt"

	"github.com/crit/secretword/pkg/secretword"
)

func main() {
	// compare answer against a guess to get a result set for all characters.
	result := secret.Compare("peach", "party") // wordle.Result{Correct, Elsewhere, Incorrect, Incorrect, Incorrect}
	fmt.Println(result)
}
```
