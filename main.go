package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type RaceType string

const (
	Race1000m RaceType = "1000m"
	EggRace   RaceType = "eggRace"
	SackRace  RaceType = "sackRace"
)

type person struct {
	Id    int
	Name  string
	Races []race
}

func (p person) getTotalTime() (time.Duration, error) {

	if len(p.Races) != 3 {
		return 0, fmt.Errorf("%v did not participate in all 3 races.", p.Name)
	}

	sum := 0 * time.Second

	for _, race := range p.Races {
		sum += race.FinalTime
	}

	return sum, nil
}

type race struct {
	RaceType  RaceType
	StartTime string
	EndTime   string
	FinalTime time.Duration
}

func NewPerson(name string, id int) *person {
	return &person{
		Id:    id,
		Name:  name,
		Races: make([]race, 0),
	}
}

func NewRace(RaceType RaceType, StartTime string, EndTime string, FinalTime time.Duration) *race {
	return &race{
		RaceType,
		StartTime,
		EndTime,
		FinalTime,
	}
}

func main() {
	file, err := os.Open("race-results.txt")

	if err != nil {
		fmt.Println(err)
		return
	}

	participants := parsePerson(file)
	printWinners(participants)

}

func printWinners(participants map[int]person) {

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

	for _, winner := range winners {
		totalTime, _ := winner.getTotalTime()
		averageTime := totalTime / 3

		fmt.Println("\n-------------------------------------------------")
		fmt.Printf("Winner is %v with ID:%v \n", winner.Name, winner.Id)
		fmt.Printf("Total time %v\n", totalTime.String())
		fmt.Printf("Average time %v\n\n", averageTime.String())
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
			log.Printf("The line is not in the correct format: '%v'\n", split)
			continue
		}

		id, err := strconv.Atoi(split[1])

		if err != nil {
			log.Println(fmt.Errorf("Unable to parse Id: '%v' on line %v", split[1], lineNumber))
			continue
		}

		race, err := getRace(split[4])

		if err != nil {
			log.Printf("%v on line %v\n", err, lineNumber)
			continue
		}

		finalTime, err := getRaceTime(split[2], split[3])

		if err != nil {
			log.Printf("%v in line %v\n", err, lineNumber)
			continue
		}

		person, ok := persons[id]

		if ok {
			newRace := NewRace(race, split[2], split[3], finalTime)
			person.Races = append(person.Races, *newRace)
			persons[id] = person
		} else {
			newPerson := NewPerson(split[0], id)
			newRace := NewRace(race, split[2], split[3], finalTime)
			newPerson.Races = append(newPerson.Races, *newRace)
			persons[id] = *newPerson
		}

	}

	file.Close()
	return persons
}

func getRaceTime(start string, end string) (time.Duration, error) {
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
