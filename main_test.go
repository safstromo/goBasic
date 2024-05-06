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

func TestGetRaceTime(t *testing.T) {

	start := "01:01:01"
	end := "02:02:02"

	duration, err := getRaceTime(start, end)
	if err != nil {
		t.Fatal(err)
	}

	if duration.String() != "1h1m1s" {
		t.Errorf("Invalid result, duration is: %v, expected 1h1m1s", duration)
	}

	start2 := "01:01:01"
	end2 := "24:02:01"

	duration2, err := getRaceTime(start2, end2)

	if err == nil {
		t.Errorf("Invalid result, duration should be an error but returned %v", duration2)
	}

}
