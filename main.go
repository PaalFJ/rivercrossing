package main

import (
	"fmt"
	"strings"
)

// string varriabler av "verdenen"
var west = "(Kylling, Korn, MSK, Rev ---W___"
var east = "___E--- )"
var boat = "\\__/"

func main() {
	ViewState()
	PutInBoatWest("Kylling")
	PutInBoatWest("MSK")
	ViewState()
	CrossWaterEast("Kylling")
	ViewState()

}

/* Skriver ut bilde av verdenen */
func ViewState() {
	fmt.Println(west + boat + east)
}

// flytter objekter(items) fra vestside til båt.
func PutInBoatWest(item string) {
	switch item {
	case "Kylling":
		west = strings.ReplaceAll(west, "Kylling", "")               // ser etter lengeden på string og kutter opp
		boat = boat[:len(boat)-3] + " Kylling " + boat[len(boat)-3:] // sånn at 3 siste tegne i stringen blir før kylling
		break                                                        // og 3 siste blir etter kylling.
	case "MSK":
		west = strings.ReplaceAll(west, "MSK", "")
		boat = boat[:len(boat)-3] + " MSK " + boat[len(boat)-3:]
	case "Rev":
		west = strings.ReplaceAll(west, "Rev", "")
		boat = boat[:len(boat)-3] + " Rev " + boat[len(boat)-3:]
	case "Korn":
		west = strings.ReplaceAll(west, "Korn", "")
		boat = boat[:len(boat)-3] + " Korn " + boat[len(boat)-3:]
	}
}

// Flytter objekter(items) fra båt til østsiden
func CrossWaterEast(item string) {
	switch item {
	case "Korn":
		boat = strings.ReplaceAll(boat, "Korn", "")
		east = east[:len(east)-1] + " Korn " + east[len(east)-1:]
		break
	case "Rev":
		boat = strings.ReplaceAll(boat, "Rev", "")
		east = east[:len(east)-1] + " Rev " + east[len(east)-1:]
		break
	case "MSK":
		boat = strings.ReplaceAll(boat, "MSK", "")
		east = east[:len(east)-1] + " MSK " + east[len(east)-1:]
		break
	case "Kylling":
		boat = strings.ReplaceAll(boat, "Kylling", "")
		east = east[:len(east)-1] + " Kylling " + east[len(east)-1:]
		break
	}
}

// Henter ut west verdien for å bruke i test på rc.test.go
func GetWest() string {
	return west
}

func GetBoat() string {
	return boat
}

/**
// Ved evt implementering av putInBoatEast, vil også GetEast være en måte å teste varriabler.
func GetEast() string {
	return east
}
*/

