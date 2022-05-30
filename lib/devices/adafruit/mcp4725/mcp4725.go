package mcp4725

// https://github.com/adafruit/Adafruit_MCP4725/blob/master/Adafruit_MCP4725.h

// #define MCP4725_I2CADDR_DEFAULT (0x62) ///< Default i2c address
// #define MCP4725_CMD_WRITEDAC (0x40)    ///< Writes data to the DAC
// #define MCP4725_CMD_WRITEDACEEPROM                                             \
//   (0x60) ///< Writes data to the DAC and the EEPROM (persisting the assigned
//          ///< value after reset)

// /**************************************************************************/
// /*!
//     @brief  Class for communicating with an MCP4725 DAC
// */
// /**************************************************************************/
// class Adafruit_MCP4725 {
// public:
//   Adafruit_MCP4725();
//   bool begin(uint8_t i2c_address = MCP4725_I2CADDR_DEFAULT,
//              TwoWire *wire = &Wire);
//   bool setVoltage(uint16_t output, bool writeEEPROM,
//                   uint32_t dac_frequency = 400000);

// private:
//   Adafruit_I2CDevice *i2c_dev = NULL;
// };

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"periph.io/x/conn/v3/i2c"
	"periph.io/x/conn/v3/physic"
)

const MaxVoltage = 0x0FFF
const MinVoltage = 0x0000

const DefaultAddress uint16 = 0x62
const CommandDontPersistWrite byte = 0x40
const CommandPersistWrite byte = 0x60

type Mcp4725 struct {
	Address uint16
	Bus     i2c.Bus
}

func (m Mcp4725) SetVoltage(level int16, persist bool, busFrequency physic.Frequency) error {
	var err error

	log.WithField("level", level).Debug("Preparing packet to set voltage output")
	packet := preparePacket(level, persist)

	log.Debug("Preparing to set voltage")
	err = m.setVoltage(packet, busFrequency)

	return err
}

func (m Mcp4725) setVoltage(packet [3]byte, busFrequency physic.Frequency) error {
	// set I2C frequency
	log.WithField("address", fmt.Sprintf("0x%x", m.Address)).Debug("Preparing to write to I2C bus")
	m.Bus.SetSpeed(busFrequency)
	err := m.Bus.Tx(m.Address, packet[:], nil)

	if err != nil {
		log.WithError(err).Error("Failed to write to I2C bus")
		return err
	}

	log.WithField("address", fmt.Sprintf("0x%x", m.Address)).Debug("Wrote to I2C bus")
	// set I2c frequency back to whatever it was before?

	return err
}

func preparePacket(level int16, persist bool) [3]byte {
	var packet [3]byte

	log.Debug("Preparing packet")

	if persist {
		packet[0] = CommandPersistWrite
	} else {
		packet[0] = CommandDontPersistWrite
	}

	packet[1] = byte(level / 16)
	packet[2] = byte((level % 16) << 4)

	log.WithField("packet", packet).Debug("Packet prepared")

	return packet
}

// bool Adafruit_MCP4725::setVoltage(uint16_t output, bool writeEEPROM,
// 	uint32_t i2c_frequency) {
// i2c_dev->setSpeed(i2c_frequency); // Set I2C frequency to desired speed

// uint8_t packet[3];

// if (writeEEPROM) {
// packet[0] = MCP4725_CMD_WRITEDACEEPROM;
// } else {
// packet[0] = MCP4725_CMD_WRITEDAC;
// }
// packet[1] = output / 16;        // Upper data bits (D11.D10.D9.D8.D7.D6.D5.D4)
// packet[2] = (output % 16) << 4; // Lower data bits (D3.D2.D1.D0.x.x.x.x)

// if (!i2c_dev->write(packet, 3)) {
// return false;
// }

// i2c_dev->setSpeed(100000); // reset to arduino default
// return true;
// }
