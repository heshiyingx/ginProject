package coding

import (
	"encoding/binary"
	"fmt"
	"io"
)

// MsgHeader aa
const MsgHeader = "12345678"

// Encode 编码
func Encode(bytesBffer io.Writer, content string) error {
	if err := binary.Write(bytesBffer, binary.BigEndian, []byte(MsgHeader)); err != nil {
		return err
	}

	clen := int32(len([]byte(content)))
	if err := binary.Write(bytesBffer, binary.BigEndian, clen); err != nil {
		return err
	}

	if err := binary.Write(bytesBffer, binary.BigEndian, []byte(content)); err != nil {
		return err
	}
	return nil
}

// Decode asd
func Decode1(bytesBuffer io.Reader) ([]byte, error) {
	MagicBuf := make([]byte, len(MsgHeader))
	if _, err := io.ReadFull(bytesBuffer, MagicBuf); err != nil {
		return nil, err
	}

	if string(MagicBuf) != MsgHeader {
		return nil, fmt.Errorf("msg_header error")
	}

	lengthBuf := make([]byte, 4)
	if _, err := io.ReadFull(bytesBuffer, lengthBuf); err != nil {
		return nil, err
	}

	length := binary.BigEndian.Uint32(lengthBuf)
	bodyBuf := make([]byte, length)
	if _, err := io.ReadFull(bytesBuffer, bodyBuf); err != nil {
		return nil, err
	}
	return bodyBuf, nil
}
