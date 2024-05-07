package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
	"unicode"
)

func main() {
	file, err := os.Open("race-results.txt")

	if err != nil {
		log.Println(err)
		return
	}

	participants := parsePersons(file)

	file.Close()

	winners := calculateWinners(participants)

	printWinners(winners)
}

func parsePersons(file *os.File) map[int]person {
	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)

	persons := make(map[int]person)

	lineNumber := 0

	for scanner.Scan() {
		lineNumber++
		split := strings.Split(scanner.Text(), ",")

		if len(split) < 4 {
			log.Printf("The line is not in the correct format: '%v'\n", split)
			continue
		}

		name, err := validateName(split[0])

		if err != nil {
			log.Printf("%v on line %v", err, lineNumber)
			continue
		}

		id, err := strconv.Atoi(split[1])

		if err != nil {
			log.Println(fmt.Errorf("Unable to parse Id: '%v' on line %v", split[1], lineNumber))
			continue
		}

		race, err := parseRace(split[4])

		if err != nil {
			log.Printf("%v on line %v\n", err, lineNumber)
			continue
		}

		finalTime, err := parseRaceTime(split[2], split[3])

		if err != nil {
			log.Printf("%v in line %v\n", err, lineNumber)
			continue
		}

		person, ok := persons[id]

		if ok {
			newRace := NewRace(race, split[2], split[3], finalTime)
			person.addRace(newRace)
			persons[id] = person
		} else {
			newPerson := NewPerson(name, id)
			newRace := NewRace(race, split[2], split[3], finalTime)
			newPerson.addRace(newRace)
			persons[id] = *newPerson
		}
	}
	return persons
}

func validateName(name string) (string, error) {

	if strings.TrimSpace(name) == "" {
		return "", fmt.Errorf("Error parsing name, name cannot be empty")
	}

	for _, char := range name {
		if !unicode.IsLetter(char) && !unicode.IsSpace(char) {
			return "", fmt.Errorf("Error parsing name, it can only contain letters and spaces. '%v'", name)
		}

	}
	return name, nil
}

func parseRaceTime(start string, end string) (time.Duration, error) {
	startTime, err := time.Parse("15:04:05", start)
	if err != nil {
		return 0, fmt.Errorf("Error parsing startTime: %v", err)
	}

	endTime, err := time.Parse("15:04:05", end)
	if err != nil {
		return 0, fmt.Errorf("Error parsing endTime: %v", err)
	}

	finalTime := endTime.Sub(startTime)

	if finalTime <= 0 {
		return 0, fmt.Errorf("Error parsing finalTime is less than 0: '%v'", err)
	}

	return finalTime, nil
}

func parseRace(raceType string) (RaceType, error) {
	switch raceType {
	case "1000m":
		return Race1000m, nil
	case "eggRace":
		return EggRace, nil
	case "sackRace":
		return SackRace, nil
	}

	return "", fmt.Errorf("Unable to parse race string: '%v'", raceType)
}

func calculateWinners(participants map[int]person) []*person {

	var winners []*person
	var winnerTime time.Duration

	for _, person := range participants {

		personTime, _ := person.getTotalTime()

		if len(person.Races) == 3 {
			if winners == nil || personTime <= winnerTime {
				if personTime < winnerTime {
					winners = nil
				}
				winners = append(winners, &person)
				winnerTime = personTime

			}
		}
	}
	return winners
}

func printWinners(winners []*person) {

	for _, winner := range winners {
		totalTime, _ := winner.getTotalTime()
		averageTime := totalTime / 3

		fmt.Println("\n-------------------------------------------------")
		fmt.Printf("Winner is %v with ID:%v \n", winner.Name, winner.Id)
		fmt.Printf("Total time %v\n", totalTime.String())
		fmt.Printf("Average time %v\n\n", averageTime.String())
	}
}
