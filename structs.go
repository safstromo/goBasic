package main

import (
	"fmt"
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
