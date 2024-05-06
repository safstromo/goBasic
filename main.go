package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type RaceType string

const (
	Race1000m RaceType = "1000m"
	EggRace   RaceType = "eggRace"
	SackRace  RaceType = "sackRace"
)

type person struct {
	Name  string
	Races []race
}

type race struct {
	RaceType  RaceType
	StartTime string
	EndTime   string
}

func NewPerson(name string) *person {
	return &person{
		Name:  name,
		Races: make([]race, 0),
	}
}

func NewRace(RaceType RaceType, StartTime string, EndTime string) *race {
	return &race{
		RaceType,
		StartTime,
		EndTime,
	}
}

func main() {
	file, err := os.Open("race-results.txt")

	if err != nil {
		fmt.Println(err)
		return
	}

	participants := parsePerson(file)

	for _, person := range participants {
		fmt.Println(person)
	}
}

func parsePerson(file *os.File) map[int]person {
	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)

	persons := make(map[int]person)

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

		person, ok := persons[id]

		if ok {
			newRace := NewRace(race, split[2], split[3])
			person.Races = append(person.Races, *newRace)
			persons[id] = person
		} else {

			newPerson := NewPerson(split[0])
			newRace := NewRace(race, split[2], split[3])
			newPerson.Races = append(newPerson.Races, *newRace)
			persons[id] = *newPerson
		}

	}

	file.Close()
	return persons
}

func getRace(raceType string) (RaceType, error) {
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
