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

	var thousandMetersRace []person
	var eggRace []person
	var stackRace []person

	for _, participant := range participants {

		switch participant.RaceType {

		case Race1000m:
			if idAndNameIsUnique(thousandMetersRace, participant) {
				thousandMetersRace = append(thousandMetersRace, participant)
			}
		case EggRace:
			if idAndNameIsUnique(eggRace, participant) {
				eggRace = append(eggRace, participant)
			}
		case SackRace:
			if idAndNameIsUnique(stackRace, participant) {
				stackRace = append(stackRace, participant)
			}
		}
	}

	fmt.Println(thousandMetersRace)
	fmt.Println(eggRace)
	fmt.Println(stackRace)
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

		if !idExistInRace(persons, id, race) {
			person := NewPerson(split[0], id, split[2], split[3], race)
			persons = append(persons, person)
		}

	}

	file.Close()
	return persons
}

func idAndNameIsUnique(personSlice []person, personToAdd person) bool {
	for _, person := range personSlice {
		if person.Id == personToAdd.Id && person.Name == personToAdd.Name {
			fmt.Printf("Unable to parse person, Id or name is not unique: Id:'%v' Name:'%v' this exist: %v \n", personToAdd.Id, personToAdd.Name, person)
			return false
		}
	}
	return true
}

func idExistInRace(personSlice []person, id int, race Race) bool {
	for _, person := range personSlice {
		if person.Id == id && person.RaceType == race {
			fmt.Printf("A person with id: '%v' already exist %v \n", id, person)
			return true
		}
	}
	return false
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
