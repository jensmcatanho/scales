package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"scales/models"
	"scales/representations"

	"github.com/urfave/cli"
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
	initScales()

	app := cli.NewApp()
	app.Flags = []cli.Flag {
		cli.BoolFlag{
			Name: "maj, M",
			Usage: "list only the major scales",
		},
		cli.BoolFlag{
			Name: "min, m",
			Usage: "list only the natural minor scales",
		},
		cli.BoolFlag{
			Name: "enharmonic, e",
			Usage: "list notes from the scales with their enharmonics",
		},
	}
	
	app.Action = func(c *cli.Context) error {
		if c.Bool("maj") {
			fmt.Println("--- Major Scales ---")
			for _, scale := range majorScales {
				scale.CheckNotes(c.Args(), c.Bool("enharmonic"))
			}

			return nil
			
		} else if c.Bool("min") {
			fmt.Println("--- Natural Minor Scales ---")
			for _, scale := range naturalMinorScales {
				scale.CheckNotes(c.Args(), c.Bool("enharmonic"))
			}
			
			return nil

		} else if len(c.Args()) < 1 {
			cli.ShowAppHelp(c)
			return nil

		}

		fmt.Println("--- Major Scales ---")
		for _, scale := range majorScales {
			scale.CheckNotes(c.Args(), c.Bool("enharmonic"))
		}
	
		fmt.Println("--- Natural Minor Scales ---")
		for _, scale := range naturalMinorScales {
			scale.CheckNotes(c.Args(), c.Bool("enharmonic"))
		}
		
		return nil
	}

	app.Run(os.Args)

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
