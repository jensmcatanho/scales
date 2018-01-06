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
	majorScales        []models.Scale
	naturalMinorScales []models.Scale
	harmonicMinorScales []models.Scale
	melodicMinorScales  []models.Scale
)

func loadScales(fp ...string) (scales []representations.Scale, err error) {
	for _, file := range fp {
		scalesRepresentation := []representations.Scale{}
		var configFile *os.File

		configFile, err = os.Open(file)
		if err != nil {
			return
		}

		decoder := json.NewDecoder(configFile)
		err = decoder.Decode(&scalesRepresentation)
		if err != nil {
			return
		}

		scales = append(scales, scalesRepresentation...)
	}

	return
}

func initScales() {

	scales, err := loadScales(
		"config/major.json",
		"config/naturalMinor.json",
		"config/harmonicMinor.json",
		"config/melodicMinor.json",
	)
	if err != nil {
		log.Fatal(err)
	}

	for _, scale := range scales {
		switch scale.Type {
		case "Major":
			majorScales = append(majorScales, models.NewScaleFromRepresentation(scale))
		case "Natural Minor":
			naturalMinorScales = append(naturalMinorScales, models.NewScaleFromRepresentation(scale))
		case "Harmonic Minor":
			harmonicMinorScales = append(harmonicMinorScales, models.NewScaleFromRepresentation(scale))
		case "Melodic Minor":
			melodicMinorScales = append(melodicMinorScales, models.NewScaleFromRepresentation(scale))
		}
	}
}

func main() {
	initScales()

	app := cli.NewApp()
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "maj, M",
			Usage: "list only the major scales",
		},
		cli.BoolFlag{
			Name:  "min, m",
			Usage: "list only the natural minor scales",
		},
		cli.BoolFlag{
			Name: "tonality, t",
			Usage: "list notes from scales",
		},
		cli.BoolFlag{
			Name:  "enharmonic, e",
			Usage: "list notes from the scales with their enharmonics",
		},
	}

	app.Action = func(c *cli.Context) error {
		if c.Bool("maj") {
			fmt.Println("--- Major ---")
			for _, scale := range majorScales {
				scale.CheckNotes(c.Args(), c.Bool("enharmonic"), c.Bool("tonality"))
			}

			return nil

		} else if c.Bool("min") {
			fmt.Println("--- Natural Minor ---")
			for _, scale := range naturalMinorScales {
				scale.CheckNotes(c.Args(), c.Bool("enharmonic"), c.Bool("tonality"))
			}

			return nil

		} else if len(c.Args()) < 1 {
			cli.ShowAppHelp(c)
			return nil

		}

		fmt.Println("--- Major ---")
		for _, scale := range majorScales {
			scale.CheckNotes(c.Args(), c.Bool("enharmonic"), c.Bool("tonality"))
		}

		fmt.Println("--- Natural Minor ---")
		for _, scale := range naturalMinorScales {
			scale.CheckNotes(c.Args(), c.Bool("enharmonic"), c.Bool("tonality"))
		}

		fmt.Println("--- Harmonic Minor ---")
		for _, scale := range harmonicMinorScales {
			scale.CheckNotes(c.Args(), c.Bool("enharmonic"), c.Bool("tonality"))
		}

		fmt.Println("--- Melodic Minor ---")
		for _, scale := range melodicMinorScales {
			scale.CheckNotes(c.Args(), c.Bool("enharmonic"), c.Bool("tonality"))
		}

		return nil
	}

	app.Run(os.Args)
}
