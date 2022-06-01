# `arpctl`

A command line utility for ARPeggiation ConTroL.

## Usage

###  `gpio`

- `list` - list all gpio pins on the current device
- `pulse` - simple `pulse` command to toggle a specific pin

###  `i2c`

Note: some features will eventually moved to `dac`.

- `list` - enumerate all I2C busses & pins on the current device
- `test` - runs a test over a duration and steps up voltage on an Adafruit MCP4725 DAC.