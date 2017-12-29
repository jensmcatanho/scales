package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type Scale struct {
	Name string
	Notes []string
}

func NewScale(name string, notes string) Scale {
	return Scale{
		Name:    name,
		Notes: strings.Split(notes, " "),
	}
}

func (s *Scale) CheckNotes(notes []string) {
	notesFound := 0

	for _, testNote := range(notes) {
		for _,  scaleNote:= range(s.Notes) {
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
	fmt.Printf("%v: %v\n", s.Name, s.Notes)
}

func main() {
	var scales []Scale

	if len(os.Args) < 2 {
		log.Fatal("Invalid number of arguments.")
	}

	MAJOR_C := "C D E F G A B"
	MAJOR_D := "D E F# G A B C#"
	MAJOR_E := "E F# G# A B C D#"
	MAJOR_F := "F G A Bb C D E"
	MAJOR_G := "G A B C D E F#"
	MAJOR_A := "A B C# D E F# G#"
	MAJOR_B := "B C# D# E F# G# A#"


	scales = append(scales, NewScale("C Major", MAJOR_C))
	scales = append(scales, NewScale("D Major", MAJOR_D))
	scales = append(scales, NewScale("E Major", MAJOR_E))
	scales = append(scales, NewScale("F Major", MAJOR_F))
	scales = append(scales, NewScale("G Major", MAJOR_G))
	scales = append(scales, NewScale("A Major", MAJOR_A))
	scales = append(scales, NewScale("B Major", MAJOR_B))

	for _, scale := range(scales) {
		scale.CheckNotes(os.Args[1:])
	}
}
