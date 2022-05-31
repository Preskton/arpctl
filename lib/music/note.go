package music

type Note struct {
	Name                string
	Letter              string
	Number              uint
	EnharmonicSharpName string
	EnharmonicFlatName  string
	Frequency           float64
	Wavelength          float64
	// TODO voltage value ranges for MCP4725 --> move this to dedicated classes/maps
	// this class shouldn't know these outputs exist
	// these are OK to be int16s b/c that's the top range of the MCP4725
	Castor    uint16
	Pollux    uint16
	Werkstatt uint16
}

var AllNotes []Note

func GetNoteByName(name string) *Note {
	for _, note := range AllNotes {
		if isNameMatch(name, &note) {
			return &note
		}
	}

	return nil
}

func isNameMatch(name string, note *Note) bool {
	return note != nil && note.Name == name
}

func isEnharmonicMatch(name string, note *Note) bool {
	return note != nil && (note.EnharmonicSharpName == name || note.EnharmonicFlatName == name)
}

func GetNoteByGeneralSearch(target string) *Note {
	for _, note := range AllNotes {
		if isNameMatch(target, &note) {
			return &note
		}

		if isEnharmonicMatch(target, &note) {
			return &note
		}
	}

	return nil
}
