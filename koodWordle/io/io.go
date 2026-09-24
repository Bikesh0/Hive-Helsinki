package io

import (
	"bufio"
	"encoding/csv"
	"os"
	"strconv"

	"koodWordle/model"
)

func LoadWords(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var words []string

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		word := scanner.Text()

		if word != "" {
			words = append(words, word)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return words, nil
}

func SaveGame(result model.GameResult) error {
	file, err := os.OpenFile(
		"stats.csv",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0644,
	)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	record := []string{
		result.Username,
		result.SecretWord,
		strconv.Itoa(result.Attempts),
		result.Result,
	}

	return writer.Write(record)
}

func LoadStats(username string) ([]model.GameResult, error) {
	file, err := os.Open("stats.csv")
	if err != nil {
		if os.IsNotExist(err) {
			return []model.GameResult{}, nil
		}

		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)

	var results []model.GameResult

	for {
		record, err := reader.Read()

		if err != nil {
			break
		}

		if len(record) != 4 {
			continue
		}

		attempts, err := strconv.Atoi(record[2])
		if err != nil {
			continue
		}

		if record[0] == username {
			results = append(results, model.GameResult{
				Username:   record[0],
				SecretWord: record[1],
				Attempts:   attempts,
				Result:     record[3],
			})
		}
	}

	return results, nil
}