package models

import (
	"fmt"
	"scales/representations"
	"strings"
)

type Scale struct {
	Fundamental string
	Type        string
	Notes       []Note
}

func NewScaleFromRepresentation(scaleRepresentation representations.Scale) Scale {
	var notes []Note
	tonality := tonality(scaleRepresentation.Type)

	for n, note := range strings.Split(scaleRepresentation.Notes, " ") {
		notes = append(notes, NewNote(note, tonality[n]))
	}

	return Scale{
		Fundamental: scaleRepresentation.Fundamental,
		Type:        scaleRepresentation.Type,
		Notes:       notes,
	}
}

func (s *Scale) CheckNotes(notes []string, enharmonic, tonality bool) {
	notesFound := 0

	for _, testNote := range notes {
		for _, scaleNote := range s.Notes {
			if strings.Compare(scaleNote.Name, testNote) == 0 {
				notesFound++
				break
			} else if enharmonic && strings.Compare(scaleNote.Enharmonic, testNote) == 0 {
				notesFound++
				break
			}
		}
	}

	if notesFound == len(notes) {
		s.print(enharmonic, tonality)
	}
}

func (s *Scale) print(enharmonic, tonality bool) {
	fmt.Printf("%v %v:", s.Fundamental, s.Type)

	for _, note := range s.Notes {
		if enharmonic && len(note.Enharmonic) > 0 {
			if tonality {
				fmt.Printf(" %v/%v%v", note.Name, note.Enharmonic, note.Tonality)
				continue
			}

			fmt.Printf(" %v/%v", note.Name, note.Enharmonic)
			continue

		} else if tonality {
			fmt.Printf(" %v%v", note.Name, note.Tonality)
			continue
		}

		fmt.Printf(" %v", note.Name)
	}

	fmt.Println("\n")
}

func tonality(scaleType string) []string {
	switch scaleType {
	case "Major":
		return []string{"7M", "m7", "m7", "7M", "7", "m7", "m7(b5)"}
	case "Natural Minor":
		return []string{"7M", "m7", "m7", "7M", "7", "m7", "m7(b5)"}
	case "Harmonic Minor":
		return []string{"7M", "m7", "m7", "7M", "7", "m7", "m7(b5)"}
	case "Melodic Minor":
		return []string{"7M", "m7", "m7", "7M", "7", "m7", "m7(b5)"}
	}

	return []string{}
}
