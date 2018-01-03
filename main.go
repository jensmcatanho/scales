package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
)

var (
	MajorScales        []Scale
	NaturalMinorScales []Scale
)

type Scale struct {
	Fundamental string `json: fundamental`
	Type        string `json: type`
	Notes       string `json: notes`
}

func NewScale(fundamental, scaleType, notes string) Scale {
	return Scale{
		Fundamental: fundamental,
		Type:        scaleType,
		Notes:       notes,
	}
}

func (s *Scale) CheckNotes(notes []string) {
	notesFound := 0

	for _, testNote := range notes {
		for _, scaleNote := range strings.Split(s.Notes, " ") {
			if strings.Compare(scaleNote, testNote) == 0 {
				notesFound++
				break
			}
		}
	}

	if notesFound == len(notes) {
		s.print()
	}
}

func (s *Scale) print() {
	fmt.Printf("%v %v: %v\n", s.Fundamental, s.Type, s.Notes)
}

func initScales() {
	configFile, err := os.Open("config.json")
	if err != nil {
		log.Fatal(err)
	}

	decoder := json.NewDecoder(configFile)
	scales := []Scale{}
	err = decoder.Decode(&scales)
	if err != nil {
		log.Fatal(err)
	}

	for _, scale := range scales {
		switch scale.Type {
		case "Major":
			MajorScales = append(MajorScales, scale)
		case "Natural Minor":
			NaturalMinorScales = append(NaturalMinorScales, scale)
		}
	}
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Invalid number of arguments.")
	}

	initScales()

	fmt.Println("--- Major Scales ---")
	for _, scale := range MajorScales {
		scale.CheckNotes(os.Args[1:])
	}

	fmt.Println("--- Natural Minor Scales ---")
	for _, scale := range NaturalMinorScales {
		scale.CheckNotes(os.Args[1:])
	}
}
