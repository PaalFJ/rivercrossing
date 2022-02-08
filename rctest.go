package main

import (
	"strings"
	"testing"
)

// Tester hva som er igjen på vestsiden av verdenen.
// String.contains ser om stringen inneholder spesifisert "string", hvis den gjør det med en eller flere variabler utløses en
// error melding.
func TestInPutBoatWest(t *testing.T) {

	PutInBoatWest("MSK")
	PutInBoatWest("Korn")
	ViewState()

	if !strings.Contains(GetWest(), "MSK") {
		if strings.Contains(GetWest(), "Kylling") && strings.Contains(GetWest(), "Korn") {
			t.Errorf("Illegal state, Kylling spiser Korn")
		}
		if strings.Contains(GetWest(), "Kylling") && strings.Contains(GetWest(), "Rev") {
			t.Errorf("Illegal state, Rev spiser Kylling")
		}
	}

}

// tester hva som er i båten, om den blir overbelastet, om noe spiser noe.
func TestGetBoat(t *testing.T) {

	if strings.Contains(GetBoat(), "Kylling") && strings.Contains(GetBoat(), "MSK") && strings.Contains(GetBoat(), "Rev") {
		t.Errorf("Illegal state, Overbelastet båt")
	}
	if strings.Contains(GetBoat(), "Korn") && strings.Contains(GetBoat(), "MSK") && strings.Contains(GetBoat(), "Rev") {
		t.Errorf("Illegal state, Overbelastet båt")
	}
	if strings.Contains(GetBoat(), "Korn") && strings.Contains(GetBoat(), "MSK") && strings.Contains(GetBoat(), "Kylling") {
		t.Errorf("Illegal state, Overbelastet båt")
	}
	if strings.Contains(GetBoat(), "Kylling") && strings.Contains(GetBoat(), "Rev") {
		t.Errorf("Illegal state, Rev spiser Kylling")
	}
	if strings.Contains(GetBoat(), "MSK") && strings.Contains(GetBoat(), "Korn") {
		t.Errorf("Illegal state, Rev spiser Kylling")
	}
	if strings.Contains(GetBoat(), "Korn") && strings.Contains(GetBoat(), "Kylling") {
		t.Errorf("Illegal state, Kylling spiser Korn")
	}
}

