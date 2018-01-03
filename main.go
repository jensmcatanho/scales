package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"scales/models"
	"scales/representations"
)

var (
	majorScales         []models.Scale
	naturalMinorScales  []models.Scale
/*
	harmonicMinorScales []models.Scale
	melodicMinorScales  []models.Scale
*/
)

func initScales() {
	configFile, err := os.Open("config.json")
	if err != nil {
		log.Fatal(err)
	}

	decoder := json.NewDecoder(configFile)
	scales := []representations.Scale{}
	err = decoder.Decode(&scales)
	if err != nil {
		log.Fatal(err)
	}

	for _, scale := range scales {
		switch scale.Type {
		case "Major":
			majorScales = append(majorScales, models.NewScaleFromRepresentation(scale))
		case "Natural Minor":
			naturalMinorScales = append(naturalMinorScales, models.NewScaleFromRepresentation(scale))
/*
		case "Harmonic Minor":
			harmonicMinorScales = append(harmonicMinorScales, models.NewScaleFromRepresentation(scale))
		case "Melodic Minor":
			melodicMinorScales = append(melodicMinorScales, models.NewScaleFromRepresentation(scale))
*/
		}
	}
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Invalid number of arguments.")
	}

	initScales()

	fmt.Println("--- Major Scales ---")
	for _, scale := range majorScales {
		scale.CheckNotes(os.Args[1:])
	}

	fmt.Println("--- Natural Minor Scales ---")
	for _, scale := range naturalMinorScales {
		scale.CheckNotes(os.Args[1:])
	}
	/*
		fmt.Println("--- Harmonic Minor Scales ---")
		for _, scale := range harmonicMinorScales {
			scale.CheckNotes(os.Args[1:])
		}

		fmt.Println("--- Melodic Minor Scales ---")
		for _, scale := range melodicMinorScales {
			scale.CheckNotes(os.Args[1:])
		}
	*/
}
