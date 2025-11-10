package bytes

import (
	"bytes"
	"encoding/binary"
)

// 整数到字节
func Int8ToBytes(n int8) []byte {
	bytesBuffer := bytes.NewBuffer([]byte{})
	if err := binary.Write(bytesBuffer, binary.LittleEndian, n); err != nil {
		return nil
	}
	return bytesBuffer.Bytes()
}

// 字节到整数
func BytesToInt8(data []byte) int8 {
	var res int8
	bytesBuffer := bytes.NewBuffer(data)
	if err := binary.Read(bytesBuffer, binary.LittleEndian, &res); err != nil {
		return 0
	}
	return res
}

// 整数到字节
func Int16ToBytes(n int16) []byte {
	bytesBuffer := bytes.NewBuffer([]byte{})
	if err := binary.Write(bytesBuffer, binary.LittleEndian, n); err != nil {
		return nil
	}
	return bytesBuffer.Bytes()
}

// 字节到整数
func BytesToInt16(data []byte) int16 {
	var res int16
	bytesBuffer := bytes.NewBuffer(data)
	if err := binary.Read(bytesBuffer, binary.LittleEndian, &res); err != nil {
		return 0
	}
	return res
}

// 整数到字节
func Int32ToBytes(n int32) []byte {
	bytesBuffer := bytes.NewBuffer([]byte{})
	if err := binary.Write(bytesBuffer, binary.LittleEndian, n); err != nil {
		return nil
	}
	return bytesBuffer.Bytes()
}

// 字节到整数
func BytesToInt32(data []byte) int32 {
	var res int32
	bytesBuffer := bytes.NewBuffer(data)
	if err := binary.Read(bytesBuffer, binary.LittleEndian, &res); err != nil {
		return 0
	}
	return res
}

// 整数到字节
func Int64ToBytes(n int64) []byte {
	bytesBuffer := bytes.NewBuffer([]byte{})
	if err := binary.Write(bytesBuffer, binary.LittleEndian, n); err != nil {
		return nil
	}
	return bytesBuffer.Bytes()
}

// 字节到整数
func BytesToInt64(data []byte) int64 {
	var res int64
	bytesBuffer := bytes.NewBuffer(data)
	if err := binary.Read(bytesBuffer, binary.LittleEndian, &res); err != nil {
		return 0
	}
	return res
}

// 合并字节
func BytesCombine(pBytes ...[]byte) []byte {
	return bytes.Join(pBytes, []byte{})
}
