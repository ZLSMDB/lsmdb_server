package service

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"testing"
)

/*

+---------------------+----------------------+-----------------------+
| Metadata Length (4B)| Value Name Length (4B)|       Value Name      |
+---------------------+----------------------+-----------------------+
|    Value Length (4B)|                    Value Data                    |
+---------------------+--------------------------------------------------+

*/

// Metadata 结构表示我们的数据结构
type Metadata struct {
	MetadataLength  uint32
	ValueNameLength uint32
	ValueName       string
	ValueLength     uint32
	ValueData       []byte
}

// Encode 方法将 Metadata 编码为 []byte
func (m *Metadata) Encode() []byte {
	buf := new(bytes.Buffer)

	// 写入 Metadata Length
	binary.Write(buf, binary.LittleEndian, m.MetadataLength)

	// 写入 Value Name Length
	binary.Write(buf, binary.LittleEndian, m.ValueNameLength)

	// 写入 Value Name
	buf.Write([]byte(m.ValueName))

	// 写入 Value Length
	binary.Write(buf, binary.LittleEndian, m.ValueLength)

	// 写入 Value Data
	buf.Write(m.ValueData)

	return buf.Bytes()
}

// Decode 方法从 []byte 解码为 Metadata
func (m *Metadata) Decode(data []byte) {
	buf := bytes.NewReader(data)

	// 读取 Metadata Length
	binary.Read(buf, binary.LittleEndian, &m.MetadataLength)

	// 读取 Value Name Length
	binary.Read(buf, binary.LittleEndian, &m.ValueNameLength)

	// 读取 Value Name
	nameBytes := make([]byte, m.ValueNameLength)
	buf.Read(nameBytes)
	m.ValueName = string(nameBytes)

	// 读取 Value Length
	binary.Read(buf, binary.LittleEndian, &m.ValueLength)

	// 读取 Value Data
	m.ValueData = make([]byte, m.ValueLength)
	buf.Read(m.ValueData)
}

func TestMeta(t *testing.T) {
	// 示例数据
	metadata := Metadata{
		MetadataLength:  4,
		ValueNameLength: 5,
		ValueName:       "hello",
		ValueLength:     11,
		ValueData:       []byte("Hello, World"),
	}

	// 编码为 []byte
	encodedData := metadata.Encode()
	fmt.Printf("Encoded Data: %v\n", encodedData)

	// 解码为 Metadata
	var decodedMetadata Metadata
	decodedMetadata.Decode(encodedData)
	fmt.Printf("Decoded Metadata: %+v\n", decodedMetadata)
}
