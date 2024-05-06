package main

import (
	"os"
	"testing"
)

func TestGetRaceSuccess(t *testing.T) {
	eggRace := "eggRace"
	thousandMeters := "1000m"
	sackRace := "sackRace"

	eggResult, err := getRace(eggRace)

	if eggResult != EggRace || err != nil {
		t.Errorf("Invalid result, got '%s', want '%s'", eggResult, eggRace)
	}

	thousandResult, err := getRace(thousandMeters)

	if thousandResult != Race1000m || err != nil {
		t.Errorf("Invalid result, got '%s', want '%s'", thousandResult, eggRace)
	}

	sackResult, err := getRace(sackRace)

	if sackResult != SackRace || err != nil {
		t.Errorf("Invalid result, got '%s', want '%s'", sackResult, eggRace)
	}

}

func TestGetRaceFail(t *testing.T) {

	result, err := getRace("1234")

	if result != "" && err == nil {
		t.Errorf("Invalid result, got '%s' '%s' , want 'Unable to parse Race type", result, err)
	}

}

func TestParsePerson(t *testing.T) {
	file, err := os.Open("testFile")

	if err != nil {
		t.Fatal(err)
	}

	defer file.Close()

	firstPerson := NewPerson(
		"Paris Hilton",
		4174331,
		"11:54:45",
		"11:59:53",
		SackRace,
	)

	parsedPersons := parsePerson(file)

	if len(parsedPersons) != 4 {
		t.Errorf("Failed, number of persons is: %v ", len(parsedPersons))
	}

	if parsedPersons[0] != firstPerson {
		t.Errorf("First person parsed is incorrect, got %v , expected %v", parsedPersons[0], firstPerson)
	}

}
