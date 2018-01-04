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

	for _, note := range strings.Split(scaleRepresentation.Notes, " ") {
		notes = append(notes, NewNote(note))
	}

	return Scale{
		Fundamental: scaleRepresentation.Fundamental,
		Type:        scaleRepresentation.Type,
		Notes:       notes,
	}
}

func (s *Scale) CheckNotes(notes []string) {
	notesFound := 0

	for _, testNote := range notes {
		for _, scaleNote := range s.Notes {
			if strings.Compare(scaleNote.Name, testNote) == 0 {
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
	fmt.Printf("%v %v:", s.Fundamental, s.Type,)

	for _, note := range s.Notes {
		fmt.Printf(" %v", note.Name)
	}

	fmt.Println()
}