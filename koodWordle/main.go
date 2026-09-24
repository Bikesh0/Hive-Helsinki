

	/*
	words := loadwords()
	 // Entry point, processes arguments, starts game
    index := getWordIndex()
	words := loadwords()
	secretword := words[index]
	username := getUsername()
	result := playGame(secretWord)
	saveResult(username, secretWord,result)
	showStats(username) */


	package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	index, ok := getWordIndex()
	if !ok {
		return
	}

	words, ok := loadWords()
	if !ok {
		return
	}

	if index < 0 || index >= len(words) {
		return
	}

	secretWord := words[index]

	scanner := bufio.NewScanner(os.Stdin)

	username, ok := getUsername(scanner)
	if !ok {
		return
	}

	playGame(secretWord, username, scanner)
}

func getWordIndex() (int, bool) {

	if len(os.Args) < 2 {
		return 0, false
	}

	index, err := strconv.Atoi(os.Args[1])
	if err != nil {
		return 0, false
	}

	return index, true
}

func getUsername(scanner *bufio.Scanner) (string, bool) {

	fmt.Println("Enter your username:")

	if !scanner.Scan() {
		return "", false
	}

	return scanner.Text(), true
}

func playGame(secretWord string, username string, scanner *bufio.Scanner) bool {

	maxAttempts := 6

	for attempts := 0; attempts < maxAttempts; attempts++ {

		fmt.Println("Enter your guess:")

		if !scanner.Scan() {
			return false
		}

		guess := scanner.Text()

		// Wordle game logic will go here.

		if guess == secretWord {
			fmt.Println("Correct!")
			return true
		}

		fmt.Printf("Attempts remaining: %d\n", maxAttempts-(attempts+1))
	}

	return false
}