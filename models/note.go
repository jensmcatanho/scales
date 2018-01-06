package models

type Note struct {
	Name       string
	Enharmonic string
	Tonality   string
}

func NewNote(name, tonality string) Note {
	return Note{
		Name:       name,
		Enharmonic: enharmonic(name),
		Tonality:   tonality,
	}
}

func enharmonic(note string) string {
	switch note {
	case "Cb":
		return "B"
	case "C#":
		return "Db"
	case "Db":
		return "C#"
	case "D#":
		return "Eb"
	case "Eb":
		return "D#"
	case "E#":
		return "F"
	case "F#":
		return "Gb"
	case "Gb":
		return "F#"
	case "G#":
		return "Ab"
	case "Ab":
		return "G#"
	case "A#":
		return "Bb"
	case "Bb":
		return "A#"
	case "B#":
		return "C"
	}

	return ""
}
