package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Sensor struct {
	ID    int
	Name  string
	Value string
}

func main() {
	sensors := []Sensor{
		{ID: 1, Name: "airTemp", Value: "NULL"},
		{ID: 2, Name: "airPressure", Value: "NULL"},
		{ID: 7, Name: "precipitation", Value: "NULL"},
		{ID: 11, Name: "windSpeed", Value: "NULL"},
		{ID: 12, Name: "windDirection", Value: "NULL"},
		{ID: 13, Name: "humidity", Value: "NULL"},
		{ID: 14, Name: "dewPoint", Value: "NULL"},
		{ID: 15, Name: "soilMoisture", Value: "NULL"},
		{ID: 22, Name: "cloudCover", Value: "NULL"},
	}

	fmt.Println("--- Weather Station ---")

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		input := scanner.Text()

		switch input {
		case "get":
			printState(sensors)

		case "clear":
			clearState(sensors)

		case "exit":
			fmt.Println("Exiting...")
			return

		default:
			updateSensor(sensors, input)
		}
	}
}

func printState(sensors []Sensor) {
	for _, sensor := range sensors {
		fmt.Printf("%s:%s\n", sensor.Name, sensor.Value)
	}
}

func clearState(sensors []Sensor) {
	for i := range sensors {
		sensors[i].Value = "NULL"
	}
}

func updateSensor(sensors []Sensor, input string) {
	parts := strings.SplitN(input, ",", 2)
	if len(parts) != 2 {
		return
	}

	id, err := strconv.Atoi(parts[0])
	if err != nil {
		return
	}

	value := strings.TrimSuffix(parts[1], ".0")

	for i := range sensors {
		if sensors[i].ID == id {
			sensors[i].Value = value
			return
		}
	}
}
