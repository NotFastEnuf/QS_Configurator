package blheli

import (
	"errors"
	"fmt"
	"io"

	"github.com/NotFastEnuf/QS_Configurator/pkg/util"
	log "github.com/sirupsen/logrus"
)

// Interface related only
// establish and test connection to the Interface

// Send Structure
// ESC + CMD PARAM_LEN [PARAM (if len > 0)] CRC16_Hi CRC16_Lo
// Return
// ESC CMD PARAM_LEN [PARAM (if len > 0)] + ACK (uint8_t OK or ERR) + CRC16_Hi CRC16_Lo

const blheliEscapeRemote = byte(0x2E) // '.'
const blheliEscapeLocal = byte(0x2F)  // '/'

type BLHeliCmd byte

// Test Interface still present
const BLHeliCmdInterfaceTestAlive = BLHeliCmd(0x30) // '0' alive
// RETURN: ACK

// get Protocol Version Number 01..255
const BLHeliCmdProtocolGetVersion = BLHeliCmd(0x31) // '1' version
// RETURN: uint8_t VersionNumber + ACK

// get Version String
const BLHeliCmdInterfaceGetName = BLHeliCmd(0x32) // '2' name
// RETURN: String + ACK

//get Version Number 01..255
const BLHeliCmdInterfaceGetVersion = BLHeliCmd(0x33) // '3' version
// RETURN: uint8_t AVersionNumber + ACK

// Exit / Restart Interface - can be used to switch to Box Mode
const BLHeliCmdInterfaceExit = BLHeliCmd(0x34) // '4' exit
// RETURN: ACK

// Reset the Device connected to the Interface
const BLHeliCmdDeviceReset = BLHeliCmd(0x35) // '5' reset
// RETURN: ACK

// Get the Device ID connected
// const BLHeliCmdDeviceGetId = BLHeliCmd(0x36)      //'6' device id removed since 06/106
// RETURN: uint8_t DeviceID + ACK

// Initialize Flash Access for Device connected
const BLHeliCmdDeviceInitFlash = BLHeliCmd(0x37) // '7' init flash access
// RETURN: ACK

// Erase the whole Device Memory of connected Device
const BLHeliCmdDeviceEraseAll = BLHeliCmd(0x38) // '8' erase all
// RETURN: ACK

// Erase one Page of Device Memory of connected Device
const BLHeliCmdDevicePageErase = BLHeliCmd(0x39) // '9' page erase
// PARAM: uint8_t APageNumber
// RETURN: ACK

// Read to Buffer from Device Memory of connected Device // Buffer Len is Max 256 Bytes
// BuffLen = 0 means 256 Bytes
const BLHeliCmdDeviceRead = BLHeliCmd(0x3A) // ':' read Device
// PARAM: uint8_t ADRESS_Hi + ADRESS_Lo + BuffLen[0..255]
// RETURN: PARAM: uint8_t ADRESS_Hi + ADRESS_Lo + BUffLen + Buffer[0..255] ACK

// Write to Buffer for Device Memory of connected Device // Buffer Len is Max 256 Bytes
// BuffLen = 0 means 256 Bytes
const BLHeliCmdDeviceWrite = BLHeliCmd(0x3B) // ';' write
// PARAM: uint8_t ADRESS_Hi + ADRESS_Lo + BUffLen + Buffer[0..255]
// RETURN: ACK

// Set C2CK low infinite ) permanent Reset state
const BLHeliCmdDeviceC2CkLow = BLHeliCmd(0x3C) // '<'
// RETURN: ACK

// Read to Buffer from Device Memory of connected Device //Buffer Len is Max 256 Bytes
// BuffLen = 0 means 256 Bytes
const BLHeliCmdDeviceReadEEprom = BLHeliCmd(0x3D) // '=' read Device
// PARAM: uint8_t ADRESS_Hi + ADRESS_Lo + BuffLen[0..255]
// RETURN: PARAM: uint8_t ADRESS_Hi + ADRESS_Lo + BUffLen + Buffer[0..255] ACK

// Write to Buffer for Device Memory of connected Device // Buffer Len is Max 256 Bytes
// BuffLen = 0 means 256 Bytes
const BLHeliCmdDeviceWriteEEprom = BLHeliCmd(0x3E) // '>' write
// PARAM: uint8_t ADRESS_Hi + ADRESS_Lo + BUffLen + Buffer[0..255]
// RETURN: ACK

// Set Interface Mode
const BLHeliCmdInterfaceSetMode = BLHeliCmd(0x3F) // '?'
// #define imC2 0
// #define imSIL_BLB 1
// #define imATM_BLB 2
// #define imSK 3
// PARAM: uint8_t Mode
// RETURN: ACK or ACK_I_INVALID_CHANNEL

//Write to Buffer for Verify Device Memory of connected Device //Buffer Len is Max 256 Bytes
//BuffLen = 0 means 256 Bytes
const BLHeliCmdDeviceVerify = BLHeliCmd(0x40) //'@' write
//PARAM: uint8_t ADRESS_Hi + ADRESS_Lo + BUffLen + Buffer[0..255]
//RETURN: ACK

type BLHeliAck byte

const (
	BLHeliAckOk              BLHeliAck = 0x00
	BLHeliAckIUnknownError   BLHeliAck = 0x01
	BLHeliAckIInvalidCmd     BLHeliAck = 0x02
	BLHeliAckIInvalidCrc     BLHeliAck = 0x03
	BLHeliAckIVerifyError    BLHeliAck = 0x04
	BLHeliAckDInvalidCommand BLHeliAck = 0x05
	BLHeliAckDCommandFailed  BLHeliAck = 0x06
	BLHeliAckDUnknownError   BLHeliAck = 0x07
	BLHeliAckIInvalidChannel BLHeliAck = 0x08
	BLHeliAckIInvalidParam   BLHeliAck = 0x09
	BLHeliAckDGeneralError   BLHeliAck = 0x0F
)

var ackToString = map[BLHeliAck]string{
	BLHeliAckOk:              "BLHeliAckOk",
	BLHeliAckIUnknownError:   "BLHeliAckIUnknownError",
	BLHeliAckIInvalidCmd:     "BLHeliAckIInvalidCmd",
	BLHeliAckIInvalidCrc:     "BLHeliAckIInvalidCrc",
	BLHeliAckIVerifyError:    "BLHeliAckIVerifyError",
	BLHeliAckDInvalidCommand: "BLHeliAckDInvalidCommand",
	BLHeliAckDCommandFailed:  "BLHeliAckDCommandFailed",
	BLHeliAckDUnknownError:   "BLHeliAckDUnknownError",
	BLHeliAckIInvalidChannel: "BLHeliAckIInvalidChannel",
	BLHeliAckIInvalidParam:   "BLHeliAckIInvalidParam",
	BLHeliAckDGeneralError:   "BLHeliAckDGeneralError",
}

type BLHeliMode byte

const (
	BLHeliModeSiLC2  BLHeliMode = 0
	BLHeliModeSiLBLB BLHeliMode = 1
	BLHeliModeAtmBLB BLHeliMode = 2
	BLHeliModeAtmSK  BLHeliMode = 3
	BLHeliModeARMBLB BLHeliMode = 4
)

type BLHeliResponse struct {
	CMD    BLHeliCmd
	ADDR   uint16
	ACK    BLHeliAck
	PARAMS []byte
}

type BLHeliSettings struct {
	MAIN_REVISION            uint8
	SUB_REVISION             uint8
	LAYOUT_REVISION          uint8
	P_GAIN                   uint8
	I_GAIN                   uint8
	GOVERNOR_MODE            uint8
	LOW_VOLTAGE_LIMIT        uint8
	MOTOR_GAIN               uint8
	MOTOR_IDLE               uint8
	STARTUP_POWER            uint8
	PWM_FREQUENCY            uint8
	MOTOR_DIRECTION          uint8
	INPUT_PWM_POLARITY       uint8
	MODE                     uint16
	PROGRAMMING_BY_TX        uint8
	REARM_AT_START           uint8
	GOVERNOR_SETUP_TARGET    uint8
	STARTUP_RPM              uint8
	STARTUP_ACCELERATION     uint8
	VOLT_COMP                uint8
	COMMUTATION_TIMING       uint8
	DAMPING_FORCE            uint8
	GOVERNOR_RANGE           uint8
	STARTUP_METHOD           uint8
	PPM_MIN_THROTTLE         uint8
	PPM_MAX_THROTTLE         uint8
	BEEP_STRENGTH            uint8
	BEACON_STRENGTH          uint8
	BEACON_DELAY             uint8
	THROTTLE_RATE            uint8
	DEMAG_COMPENSATION       uint8
	BEC_VOLTAGE              uint8
	PPM_CENTER_THROTTLE      uint8
	SPOOLUP_TIME             uint8
	TEMPERATURE_PROTECTION   uint8
	LOW_RPM_POWER_PROTECTION uint8
	PWM_INPUT                uint8
	PWM_DITHER               uint8
	BRAKE_ON_STOP            uint8
	LED_CONTROL              uint8

	LAYOUT string
	MCU    string
	NAME   string
}

type BLHeliProtocol struct {
	rw io.ReadWriter
}

func NewBLHeliProtocol(rw io.ReadWriter) (*BLHeliProtocol, error) {
	return &BLHeliProtocol{
		rw: rw,
	}, nil
}

func (p *BLHeliProtocol) readBlheli() (*BLHeliResponse, error) {
	buf := make([]byte, 512)
	length := 0

	{
		n, err := p.rw.Read(buf)
		if err != nil {
			return nil, err
		}
		length += n
	}

	if buf[0] != blheliEscapeRemote {
		return nil, errors.New("invalid magic")
	}

	addr := uint16(buf[3]) | uint16(buf[2])<<8
	cmd, paramLen := BLHeliCmd(buf[1]), uint16(buf[4])
	if paramLen == 0 {
		paramLen = 256
	}
	size := int(5 + paramLen + 1 + 2)
	for length != size {
		n, err := p.rw.Read(buf[length:size])
		if err != nil {
			return nil, err
		}
		if n == 0 {
			break
		}
		length += n
	}

	if length != size {
		return nil, fmt.Errorf("<msp> invalid size (%d vs %d)", length, size)
	}

	crcActual := uint16(0)
	for i := 0; i < size-2; i++ {
		crcActual = util.UpdateCRC16(crcActual, uint16(buf[i]))
	}
	crcExpected := uint16(buf[5+paramLen+2]) | uint16(buf[5+paramLen+1])<<8

	if crcExpected != crcActual {
		return nil, fmt.Errorf("<msp> invalid crc (%d vs %d)", crcActual, crcExpected)
	}

	ack := BLHeliAck(buf[5+paramLen])
	log.Debugf("<blheli> received cmd: 0x%x addr: %d ack: 0x%x paramLen: %d", cmd, addr, ack, paramLen)
	return &BLHeliResponse{
		CMD:    cmd,
		ACK:    ack,
		ADDR:   addr,
		PARAMS: buf[5 : 5+paramLen],
	}, nil
}

func (p *BLHeliProtocol) SendBlheli(cmd BLHeliCmd, addr uint16, params []byte) (*BLHeliResponse, error) {
	if len(params) > 256 {
		return nil, errors.New("params >= 256")
	}

	if len(params) == 0 {
		params = append(params, 0)
	}

	paramLen := uint8(0)
	if len(params) < 256 {
		paramLen = uint8(len(params))
	}
	buf := []byte{
		blheliEscapeLocal,
		byte(cmd),
		uint8((addr >> 8) & 0xFF), //ADDR HIGH
		uint8(addr & 0xFF),        //ADDR LOW
		paramLen,                  // PARAM_LEN
	}
	buf = append(buf, params...)
	buf = util.AppendCRC16(buf)

	log.Printf("<blheli> sent cmd: 0x%x addr: %d paramLen: %d", cmd, addr, paramLen)
	p.rw.Write(buf)

	res, err := p.readBlheli()
	if err != nil {
		return nil, err
	}
	if res.CMD != cmd {
		return nil, fmt.Errorf("<blheli> invalid response cmd (0x%x vs 0x%x)", res.CMD, cmd)
	}
	if res.ACK != BLHeliAckOk {
		return nil, fmt.Errorf("<blheli> invalid ack %s (0x%x)", ackToString[res.ACK], res.ACK)
	}
	return res, nil
}

func (p *BLHeliProtocol) ReadFlash(offset, length uint16) ([]byte, error) {
	buf := make([]byte, length)
	read := uint16(0)

	for read < length {
		size := length - read
		if size > 128 {
			size = 128
		}

		res, err := p.SendBlheli(BLHeliCmdDeviceRead, offset+read, []byte{uint8(size)})
		if err != nil {
			return nil, err
		}
		log.Printf("<blheli> readFlash %d (%d)", offset+read, len(res.PARAMS))

		copy(buf[read:], res.PARAMS)
		read += uint16(len(res.PARAMS))
	}

	return buf, nil
}

func (p *BLHeliProtocol) WriteFlash(offset uint16, buf []byte) error {
	length := uint16(len(buf))

	for offset < length {
		size := length - offset
		if size > 128 {
			size = 128
		}
		res, err := p.SendBlheli(BLHeliCmdDeviceWrite, offset, buf[offset:offset+size])
		if err != nil {
			return err
		}
		log.Printf("<blheli> writeFlash ack: %d offset: %d (%d)", res.ACK, offset, size)
		offset += size
	}

	return nil
}
