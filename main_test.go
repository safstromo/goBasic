package main

import "testing"

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
