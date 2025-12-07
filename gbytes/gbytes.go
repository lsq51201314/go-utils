package gbytes

import (
	"bytes"
	"encoding/binary"
	"math"
	"strings"
	"unicode/utf16"

	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

// 整数到字节
func Uint8ToByte(num uint8) byte {
	return byte(num)
}

// 字节到整数
func ByteToUint8(b byte) uint8 {
	return uint8(b)
}

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
func Uint16ToBytes(num uint16) []byte {
	buf := new(bytes.Buffer)
	if err := binary.Write(buf, binary.LittleEndian, num); err != nil {
		return []byte{}
	}
	return buf.Bytes()
}

// 字节到整数
func BytesToUint16(data []byte) uint16 {
	return binary.LittleEndian.Uint16(data)
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

func Uint32ToBytes(num uint32) []byte {
	buf := new(bytes.Buffer)
	if err := binary.Write(buf, binary.LittleEndian, num); err != nil {
		return []byte{}
	}
	return buf.Bytes()
}

func BytesToUint32(data []byte) uint32 {
	return binary.LittleEndian.Uint32(data)
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

// 小数到字节
func Float32ToBytes(f float32) []byte {
	var buf [4]byte
	binary.LittleEndian.PutUint32(buf[:], math.Float32bits(f))
	return buf[:]
}

// 字节到小数
func BytesToFloat32(b []byte) float32 {
	return math.Float32frombits(binary.LittleEndian.Uint32(b))
}

// 字节到文本
func WStrBytesToStr(wsbyte []byte) string {
	if len(wsbyte)%2 != 0 {
		wsbyte = wsbyte[:len(wsbyte)-1]
	}
	endIndex := len(wsbyte)
	for i := 0; i < len(wsbyte)-1; i += 2 {
		if wsbyte[i] == 0 && wsbyte[i+1] == 0 {
			endIndex = i
			break
		}
	}
	wsbyte = wsbyte[:endIndex]
	if len(wsbyte) == 0 {
		return ""
	}
	u16s := make([]uint16, len(wsbyte)/2)
	for i := range u16s {
		u16s[i] = binary.LittleEndian.Uint16(wsbyte[2*i:])
	}
	runes := utf16.Decode(u16s)
	for len(runes) > 0 && runes[len(runes)-1] == 0 {
		runes = runes[:len(runes)-1]
	}
	return string(runes)
}

// 文本到字节
func StrToWStrBytes(str string) []byte {
	runes := utf16.Encode([]rune(str))
	dat := make([]byte, 0, len(runes)*2)
	for _, r := range runes {
		dat = append(dat, byte(r), byte(r>>8))
	}
	return dat
}

// 字节到文本
func BytesToStr(sbyte []byte) string {
	trimmed := bytes.TrimRight(sbyte, "\x00")
	decoder := simplifiedchinese.GBK.NewDecoder()
	result, _, err := transform.Bytes(decoder, trimmed)
	if err != nil {
		return ""
	}
	return string(result)
}

// 文本到字节
func StrToBytes(s string) []byte {
	encoder := simplifiedchinese.GBK.NewEncoder()
	var buf bytes.Buffer
	writer := transform.NewWriter(&buf, encoder)
	if _, err := writer.Write([]byte(s)); err != nil {
		return nil
	}
	if err := writer.Close(); err != nil {
		return nil
	}
	return buf.Bytes()
}

// 字节到文本
func BytesToGBK(sbyte []byte) string {
	decoder := simplifiedchinese.GBK.NewDecoder()
	result, _ := decoder.String(string(sbyte))
	if idx := strings.IndexByte(result, 0); idx != -1 {
		result = result[:idx]
	}
	return result
}

// 文本到字节
func GBKToBytes(s string) []byte {
	encoder := simplifiedchinese.GBK.NewEncoder()
	gbkBytes, err := encoder.Bytes([]byte(s))
	if err != nil {
		return nil
	}
	return append(gbkBytes, 0)
}

// 字节补齐
func PadBytes(data []byte, length int) []byte {
	if len(data) > length {
		return data[:length]
	}
	if len(data) == length {
		return data
	}
	result := make([]byte, length)
	copy(result, data)
	return result
}

// 合并字节
func BytesCombine(pBytes ...[]byte) []byte {
	return bytes.Join(pBytes, []byte{})
}
