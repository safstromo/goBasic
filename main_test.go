package main

import (
	"os"
	"testing"
)

func TestGetRaceSuccess(t *testing.T) {
	eggRace := "eggRace"
	thousandMeters := "1000m"
	sackRace := "sackRace"

	eggResult, err := parseRace(eggRace)

	if eggResult != EggRace || err != nil {
		t.Errorf("Invalid result, got '%s', want '%s'", eggResult, eggRace)
	}

	thousandResult, err := parseRace(thousandMeters)

	if thousandResult != Race1000m || err != nil {
		t.Errorf("Invalid result, got '%s', want '%s'", thousandResult, eggRace)
	}

	sackResult, err := parseRace(sackRace)

	if sackResult != SackRace || err != nil {
		t.Errorf("Invalid result, got '%s', want '%s'", sackResult, eggRace)
	}

}

func TestGetRaceFail(t *testing.T) {

	result, err := parseRace("1234")

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

	parsedPersons := parsePersons(file)

	if len(parsedPersons) != 3 {
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

	duration, err := parseRaceTime(start, end)
	if err != nil {
		t.Fatal(err)
	}

	if duration.String() != "1h1m1s" {
		t.Errorf("Invalid result, duration is: %v, expected 1h1m1s", duration)
	}

	start2 := "01:01:01"
	end2 := "24:02:01"

	duration2, err := parseRaceTime(start2, end2)

	if err == nil {
		t.Errorf("Invalid result, duration should be an error but returned %v", duration2)
	}

}

func TestValitateName(t *testing.T) {
	name1 := "Test name"
	name2 := "Test n8me"
	name3 := ""
	name4 := "T!st"

	var err error

	name1, err = validateName(name1)

	if err != nil {
		t.Errorf("Invalid result expected %v, got: %v", name1, err)
	}

	_, err = validateName(name2)

	if err == nil {
		t.Errorf("Invalid result expected %v, got: %v", name2, err)
	}

	_, err = validateName(name3)

	if err == nil {
		t.Errorf("Invalid result expected %v, got: %v", name3, err)
	}

	_, err = validateName(name4)

	if err == nil {
		t.Errorf("Invalid result expected %v, got: %v", name4, err)
	}
}
