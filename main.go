package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Race string

const (
	Race1000m Race = "1000m"
	EggRace   Race = "eggRace"
	SackRace  Race = "sackRace"
)

type person struct {
	Name      string
	Id        int
	StartTime string
	EndTime   string
	RaceType  Race
}

func NewPerson(Name string, Id int, StartTime string, EndTime string, RaceType Race) person {
	return person{
		Name,
		Id,
		StartTime,
		EndTime,
		RaceType,
	}
}

func main() {
	file, err := os.Open("race-results.txt")

	if err != nil {
		fmt.Println(err)
		return
	}

	participants := parsePerson(file)

	fmt.Println(participants)
	fmt.Println(len(participants))
}

func parsePerson(file *os.File) []person {
	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)

	var persons []person

	lineNumber := 0

	for scanner.Scan() {
		lineNumber++
		split := strings.Split(scanner.Text(), ",")

		if len(split) < 4 {
			fmt.Printf("The line is not in the correct format: '%v'\n", split)
			continue
		}

		id, err := strconv.Atoi(split[1])

		if err != nil {
			fmt.Println(fmt.Errorf("Unable to parse Id: '%v' on line %v", split[1], lineNumber))
			continue
		}

		race, err := getRace(split[4])

		if err != nil {
			fmt.Printf("%v on line %v\n", err, lineNumber)
			continue
		}

		person := NewPerson(split[0], id, split[2], split[3], race)

		persons = append(persons, person)

	}

	file.Close()
	return persons
}

func getRace(raceType string) (Race, error) {
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
