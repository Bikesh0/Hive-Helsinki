package game

import (
	"bufio"
	"fmt"
	"strings"

	"koodWordle/io"
	"koodWordle/model"
)

const (
	Green  = "\033[32m"
	Yellow = "\033[33m"
	White  = "\033[37m"
	Reset  = "\033[0m"
)

const maxAttempts = 6

func Play(
	secretWord string,
	username string,
	scanner *bufio.Scanner,
) {
	secretWord = strings.ToLower(strings.TrimSpace(secretWord))

	attempts := 0

	for attempts < maxAttempts {
		fmt.Println("Enter your guess:")

		if !scanner.Scan() {
			return
		}

		guess := strings.ToLower(strings.TrimSpace(scanner.Text()))

		if len(guess) != 5 {
			continue
		}

		attempts++

		printFeedback(guess, secretWord)

		if guess == secretWord {
			fmt.Printf("Attempts remaining: %d\n", maxAttempts-attempts)

			saveResult(username, secretWord, attempts, "win")

			showStatsOption(username, scanner)
			return
		}

		fmt.Printf("Attempts remaining: %d\n", maxAttempts-attempts)

		if attempts == maxAttempts {
			fmt.Printf("The word was: %s\n", strings.ToUpper(secretWord))

			saveResult(username, secretWord, attempts, "loss")

			showStatsOption(username, scanner)
			return
		}
	}
}

func printFeedback(guess string, secretWord string) {
	result := make([]string, 5)

	used := make([]bool, 5)

	// First pass: correct letters in the correct position.
	for i := 0; i < 5; i++ {
		if guess[i] == secretWord[i] {
			result[i] = Green + strings.ToUpper(string(guess[i])) + Reset
			used[i] = true
		}
	}

	// Second pass: correct letters in the wrong position.
	for i := 0; i < 5; i++ {
		if result[i] != "" {
			continue
		}

		found := false

		for j := 0; j < 5; j++ {
			if !used[j] && guess[i] == secretWord[j] {
				found = true
				used[j] = true
				break
			}
		}

		if found {
			result[i] = Yellow + strings.ToUpper(string(guess[i])) + Reset
		} else {
			result[i] = White + strings.ToUpper(string(guess[i])) + Reset
		}
	}

	fmt.Print("Feedback: ")

	for _, letter := range result {
		fmt.Print(letter)
	}

	fmt.Println()

	fmt.Println("Remaining letters:", remainingLetters(guess, secretWord))
}

func remainingLetters(guess string, secretWord string) string {
	incorrect := make(map[rune]bool)

	for _, letter := range guess {
		letter = []rune(strings.ToUpper(string(letter)))[0]

		if !strings.ContainsRune(
			strings.ToUpper(secretWord),
			letter,
		) {
			incorrect[letter] = true
		}
	}

	var result []string

	for letter := 'A'; letter <= 'Z'; letter++ {
		if !incorrect[letter] {
			result = append(result, string(letter))
		}
	}

	return strings.Join(result, " ")
}

func saveResult(
	username string,
	secretWord string,
	attempts int,
	result string,
) {
	gameResult := model.GameResult{
		Username:   username,
		SecretWord: secretWord,
		Attempts:   attempts,
		Result:     result,
	}

	_ = io.SaveGame(gameResult)
}

func showStatsOption(username string, scanner *bufio.Scanner) {
	fmt.Println("Do you want to see your stats? (yes/no):")

	if !scanner.Scan() {
		return
	}

	answer := strings.ToLower(strings.TrimSpace(scanner.Text()))

	if answer == "yes" {
		showStats(username)
	}
}

func showStats(username string) {
	results, err := io.LoadStats(username)
	if err != nil {
		return
	}

	gamesPlayed := len(results)
	gamesWon := 0
	totalAttempts := 0

	for _, result := range results {
		if result.Result == "win" {
			gamesWon++
		}

		totalAttempts += result.Attempts
	}

	average := float64(0)

	if gamesPlayed > 0 {
		average = float64(totalAttempts) / float64(gamesPlayed)
	}

	fmt.Printf("Stats for %s:\n", username)
	fmt.Printf("Games played: %d\n", gamesPlayed)
	fmt.Printf("Games won: %d\n", gamesWon)
	fmt.Printf("Average attempts per game: %.2f\n", average)
	fmt.Println("Press Enter to exit...")

	// The caller owns the scanner, so this function does not
	// read another line here.
}