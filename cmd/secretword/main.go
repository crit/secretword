package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"math/rand"
	"strings"
	"time"

	"github.com/crit/secretword/pkg/secretword"
	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
)

// Letter printers
var correct = color.New(color.BgGreen, color.FgBlack).PrintfFunc()
var elsewhere = color.New(color.BgHiYellow, color.FgBlack).PrintfFunc()
var incorrect = color.New(color.BgRed, color.FgWhite).PrintfFunc()

// Win/Loss printers
var win = color.New(color.Bold, color.BgGreen, color.FgBlack).PrintlnFunc()
var lose = color.New(color.Bold, color.BgRed, color.FgWhite).PrintlnFunc()

func main() {
	answer := randomAnswer(readLines(answerData))
	var history secretword.History
	var winner bool

	fmt.Println("Welcome to Secret Word")
	fmt.Println("Green (match), Yellow (elsewhere), Red (nowhere)")

	for i := secretword.TurnLimit; i > 0; i-- {
		fmt.Println("")

		// prepare prompt: "Guesses Left 6: "
		prompt := promptui.Prompt{
			Label:    fmt.Sprintf("Guesses Left %d", i),
			Validate: validWord,
		}

		// display prompt to user
		guess, err := prompt.Run()
		if err != nil {
			log.Fatal(err)
		}

		// test guess against answer
		res := secretword.Compare(answer, guess)

		// record guess and show the user
		history.Append(res, guess)
		history.Print(printResult)

		// Success ends the loop and the game.
		if res.Is(secretword.Success) {
			winner = true
			break
		}
	}

	fmt.Println("")

	// Print results of the game.
	if winner {
		win("You Won!")
	} else {
		lose("The correct answer was:", answer)
	}
}

// readLines reads in a string and separates by new line.
func readLines(data string) []string {
	buf := bytes.NewReader([]byte(data))

	var lines []string
	scanner := bufio.NewScanner(buf)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}

	return lines
}

// randomAnswer pulls a random entry in the provided slice.
func randomAnswer(lines []string) string {
	s := rand.NewSource(time.Now().Unix())
	r := rand.New(s)
	return lines[r.Intn(len(lines))]
}

// validWord makes sure that the prompt gets something we can work with.
func validWord(word string) error {
	if len(word) == secretword.CharLimit {
		return nil
	}

	return fmt.Errorf("guess must be %d characters long", secretword.CharLimit)
}

// printResult formats the wordle result for display in the terminal using the letter printers.
func printResult(res secretword.Result, guess string) {
	letters := toLetters(guess)

	for i := 0; i < len(guess); i++ {
		switch res[i] {
		case secretword.Correct:
			correct("[%s]", letters[i])
		case secretword.Elsewhere:
			elsewhere("[%s]", letters[i])
		case secretword.Incorrect:
			incorrect("[%s]", letters[i])
		}
	}

	fmt.Println("")
}

// toLetters splits the word into its letters and upper cases them. Intended for Printf interfaces.
func toLetters(word string) (letters []any) {
	for _, b := range word {
		letters = append(letters, strings.ToUpper(string(b)))
	}

	return letters
}
