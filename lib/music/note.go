package music

type Note struct {
	Name                string
	Letter              string
	AbsoluteIndex       uint
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

func something() {

}
