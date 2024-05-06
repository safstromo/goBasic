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

	parsedPersons := parsePerson(file)

	if len(parsedPersons) != 4 {
		t.Errorf("Failed, number of persons is: %v ", len(parsedPersons))
	}

	parisHiltonId := 4174331
	if _, ok := parsedPersons[parisHiltonId]; !ok {
		t.Errorf("Paris Hilton should exist in the map")
	}

}
