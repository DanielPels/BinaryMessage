package BinaryMessage

import (
	"encoding/binary"
	"bytes"
	"math"
)

type BinaryMessagePoolInterface interface {
	GetBinaryMessage() *BinaryMessage
	ReleaseBinaryMessage(*BinaryMessage)
}

type BinaryMessagePool struct {
	pool chan *BinaryMessage
}

func NewBinaryMessagePool(cap int) *BinaryMessagePool {
	return &BinaryMessagePool{
		pool: make(chan *BinaryMessage, cap),
	}
}

func (b *BinaryMessagePool) GetBinaryMessage() *BinaryMessage {
	select {
	case m := <-b.pool:
		return m
	default:
		return NewBinaryMessage()
	}
}

func (b *BinaryMessagePool) ReleaseBinaryMessage(m *BinaryMessage) {
	m.Reset()
	select {
	case b.pool <- m:
	default:
	}
}

type BinaryMessageInterface interface {
	WriteInterface(interface{})
	WriteIntAsUint8(int)
	WriteIntAsUint16(int)
	WriteIntAsUint32(int)
	WriteFloatAsUint16(float64)
	WriteFloatAsInt16(float64)
	WriteBytes([]byte)
	GetBuffer() *bytes.Buffer
	GetBytes() []byte
	HadError() bool
	Reset()
}

func NewBinaryMessage() *BinaryMessage {
	return &BinaryMessage{
		buffer: new(bytes.Buffer),
	}
}

type BinaryMessage struct {
	buffer *bytes.Buffer
	error  error
}

func (m *BinaryMessage) WriteIntAsUint32(data int) {
	err := writeToBuffer(m.buffer, uint32(data))
	if err != nil {
		m.error = err
	}
}

func (m *BinaryMessage) WriteIntAsInt32(data int) {
	err := writeToBuffer(m.buffer, int32(data))
	if err != nil {
		m.error = err
	}
}

func (m *BinaryMessage) WriteInterface(data interface{}) {
	err := writeToBuffer(m.buffer, data)
	if err != nil {
		m.error = err
	}
}
func (m *BinaryMessage) WriteIntAsUint8(data int) {
	err := writeToBuffer(m.buffer, uint8(data))
	if err != nil {
		m.error = err
	}
}
func (m *BinaryMessage) WriteIntAsUint16(data int) {
	err := writeToBuffer(m.buffer, uint16(data))
	if err != nil {
		m.error = err
	}
}
func (m *BinaryMessage) WriteFloatAsUint16(data float64) {
	err := writeToBuffer(m.buffer, uint16(math.Floor(data)))
	if err != nil {
		m.error = err
	}
}

func (m *BinaryMessage) WriteFloatAsInt16(data float64) {
	err := writeToBuffer(m.buffer, int16(math.Floor(data)))
	if err != nil {
		m.error = err
	}
}

func (m *BinaryMessage) WriteBytes(data []byte) {
	_, err := m.buffer.Write(data)
	if err != nil {
		m.error = err
	}
}
func (m *BinaryMessage) GetBuffer() *bytes.Buffer {
	return m.buffer
}
func (m *BinaryMessage) GetBytes() []byte {
	return m.buffer.Bytes()
}
func (m *BinaryMessage) HadError() bool {
	if m.error != nil {
		return true
	}
	return false
}
func (m *BinaryMessage) Reset() {
	m.buffer.Reset()
	m.error = nil
}
func writeToBuffer(buf *bytes.Buffer, data interface{}) error {
	return binary.Write(buf, binary.LittleEndian, data)
}
