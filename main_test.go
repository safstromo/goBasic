package main

import "testing"

func TestGetRace(t *testing.T) {
	var eggRace = "eggRace"

	result, err := getRace(eggRace)

	if result != EggRace || err != nil {
		t.Errorf("Invalid result, got '%s', want '%s'", result, eggRace)
	}

}
