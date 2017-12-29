package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

var (
	MajorScales []Scale
	MinorScales []Scale
)

type Scale struct {
	Fundamental       string
	Type string
	Notes      []string
}

func NewScale(fundamental, scaleType, notes string) Scale {
	return Scale{
		Fundamental: fundamental,
		Type: scaleType,
		Notes: strings.Split(notes, " "),
	}
}

func (s *Scale) CheckNotes(notes []string) {
	notesFound := 0

	for _, testNote := range notes {
		for _, scaleNote := range s.Notes {
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
	MAJOR_C_FLAT := "Cb Db Eb Fb Gb Ab Bb"
	MAJOR_C := "C D E F G A B"
	MAJOR_C_SHARP := "C# D# E# F# G# A# B#"

	MAJOR_D_FLAT := "Db Eb F Gb Ab Bb C"
	MAJOR_D := "D E F# G A B C#"

	MAJOR_E_FLAT := "Eb F G Ab Bb C D"
	MAJOR_E := "E F# G# A B C D#"

	MAJOR_F := "F G A Bb C D E"
	MAJOR_F_SHARP := "F# G# A# B C# D# E#"

	MAJOR_G_FLAT := "Gb Ab Bb Cb Db Eb F"
	MAJOR_G := "G A B C D E F#"

	MAJOR_A_FLAT := "Ab Bb C Db Eb F G"
	MAJOR_A := "A B C# D E F# G#"

	MAJOR_B_FLAT := "Bb C D Eb F G A"
	MAJOR_B := "B C# D# E F# G# A#"

	MajorScales = append(MajorScales, NewScale("Cb", "Major", MAJOR_C_FLAT))
	MajorScales = append(MajorScales, NewScale("C", "Major", MAJOR_C))
	MajorScales = append(MajorScales, NewScale("C#", "Major", MAJOR_C_SHARP))

	MajorScales = append(MajorScales, NewScale("Db", "Major", MAJOR_D_FLAT))
	MajorScales = append(MajorScales, NewScale("D", "Major", MAJOR_D))

	MajorScales = append(MajorScales, NewScale("Eb", "Major", MAJOR_E_FLAT))
	MajorScales = append(MajorScales, NewScale("E", "Major", MAJOR_E))

	MajorScales = append(MajorScales, NewScale("F", "Major", MAJOR_F))
	MajorScales = append(MajorScales, NewScale("F#", "Major", MAJOR_F_SHARP))

	MajorScales = append(MajorScales, NewScale("Gb", "Major", MAJOR_G_FLAT))
	MajorScales = append(MajorScales, NewScale("G", "Major", MAJOR_G))

	MajorScales = append(MajorScales, NewScale("Ab", "Major", MAJOR_A_FLAT))
	MajorScales = append(MajorScales, NewScale("A", "Major", MAJOR_A))

	MajorScales = append(MajorScales, NewScale("Bb", "Major", MAJOR_B_FLAT))
	MajorScales = append(MajorScales, NewScale("B", "Major", MAJOR_B))

	MINOR_C := "C D Eb F G Ab Bb"
	MINOR_C_SHARP := "C# D# E F# G# A B"

	MINOR_D := "D E F G A Bb C"
	MINOR_D_SHARP := "D# E# F# G# A# B C#"

	MINOR_E_FLAT := "Eb F Gb Ab Bb Cb Db"
	MINOR_E := "E F# G A B C D"

	MINOR_F := "F G Ab Bb C Db Eb"
	MINOR_F_SHARP := "F# G# A B C# D E"

	MINOR_G := "G A Bb C D Eb F"
	MINOR_G_SHARP := "G# A# B C# D# E F#"

	MINOR_A_FLAT := "Ab Bb Cb Db Eb Fb Gb"
	MINOR_A := "A B C D E F G"
	MINOR_A_SHARP := "A# B# C# D# E# F# G#"

	MINOR_B_FLAT := "Bb C Db Eb F Gb Ab"
	MINOR_B := "B C# D E F# G A"

	MinorScales = append(MinorScales, NewScale("C", "Minor", MINOR_C))
	MinorScales = append(MinorScales, NewScale("C#", "Minor", MINOR_C_SHARP))

	MinorScales = append(MinorScales, NewScale("D", "Minor", MINOR_D))
	MinorScales = append(MinorScales, NewScale("D#", "Minor", MINOR_D_SHARP))

	MinorScales = append(MinorScales, NewScale("Eb", "Minor", MINOR_E_FLAT))
	MinorScales = append(MinorScales, NewScale("E", "Minor", MINOR_E))

	MinorScales = append(MinorScales, NewScale("F", "Minor", MINOR_F))
	MinorScales = append(MinorScales, NewScale("F#", "Minor", MINOR_F_SHARP))

	MinorScales = append(MinorScales, NewScale("G", "Minor", MINOR_G))
	MinorScales = append(MinorScales, NewScale("G#", "Minor", MINOR_G_SHARP))

	MinorScales = append(MinorScales, NewScale("Ab", "Minor", MINOR_A_FLAT))
	MinorScales = append(MinorScales, NewScale("A", "Minor", MINOR_A))
	MinorScales = append(MinorScales, NewScale("A#", "Minor", MINOR_A_SHARP))

	MinorScales = append(MinorScales, NewScale("Bb", "Minor", MINOR_B_FLAT))
	MinorScales = append(MinorScales, NewScale("B", "Minor", MINOR_B))
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Invalid number of arguments.")
	}

	initScales()

	for _, scale := range MajorScales {
		scale.CheckNotes(os.Args[1:])
	}

	for _, scale := range MinorScales {
		scale.CheckNotes(os.Args[1:])
	}
}

