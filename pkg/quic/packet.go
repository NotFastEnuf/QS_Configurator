package quic

import (
	"io"
)

type QuicCommand uint8

const (
	QuicCmdInvalid QuicCommand = iota
	QuicCmdGet
	QuicCmdSet
	QuicCmdLog
	QuicCmdCalImu
	QuicCmdBlackbox
	QuicCmdResetBlackbox
	QuicCmdMax
)

const (
	QuicBlackboxReset = iota
	QuicBlackboxList
	QuicBlackboxGet
)

const (
	QuicFlagNone uint8 = iota
	QuicFlagError
	QuicFlagStreaming
	QuicFlagUpdate
)

type QuicValue uint8

const (
	QuicValInvalid QuicValue = iota
	QuicValInfo
	QuicValProfile
	QuicValDefaultProfile
	QuicValBlackboxRate
	QuicValPidRatePresets
	QuicValVtxSettings
	QuicValOSDFont
)

type QuicPacket struct {
	cmd  QuicCommand
	flag uint8
	len  uint16

	Payload io.ReadCloser
}

const quicHeaderLen = uint16(4)