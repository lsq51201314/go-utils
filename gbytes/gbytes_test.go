package gbytes

import (
	"fmt"
	"testing"
)

func TestGbytes(t *testing.T) {
	fmt.Println(ByteToUint8(Uint8ToByte(1)))
	fmt.Println(BytesToInt8(Int8ToBytes(1)))
	fmt.Println(BytesToUint16(Uint16ToBytes(1)))
	fmt.Println(BytesToInt16(Int16ToBytes(1)))
	fmt.Println(BytesToUint32(Uint32ToBytes(1)))
	fmt.Println(BytesToInt32(Int32ToBytes(1)))
	fmt.Println(BytesToInt64(Int64ToBytes(1)))
	fmt.Println(BytesToFloat32(Float32ToBytes(1.2)))
	fmt.Println(WStrBytesToStr(StrToWStrBytes("你好啊")))
	fmt.Println(BytesToStr(StrToBytes("你好啊")))
	fmt.Println(BytesToGBK(GBKToBytes("你好啊")))
	fmt.Println(PadBytes([]byte{}, 10))
	fmt.Println(BytesCombine([]byte{1}, []byte{2}))
}
