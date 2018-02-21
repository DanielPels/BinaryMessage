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
	WriteInterface(interface{}) error
	WriteIntAsUint8(int) error
	WriteIntAsUint16(int) error
	WriteFloatAsUint16(float64) error
	WriteBytes([]byte) error
	GetBuffer() *bytes.Buffer
	GetBytes() []byte
	Reset()
}

func NewBinaryMessage() *BinaryMessage {
	return &BinaryMessage{
		buffer: new(bytes.Buffer),
	}
}

type BinaryMessage struct {
	buffer *bytes.Buffer
}

func (m *BinaryMessage) WriteInterface(data interface{}) error {
	return writeToBuffer(m.buffer, data)
}
func (m *BinaryMessage) WriteIntAsUint8(data int) error {
	return writeToBuffer(m.buffer, uint8(data))
}
func (m *BinaryMessage) WriteIntAsUint16(data int) error {
	return writeToBuffer(m.buffer, uint16(data))
}
func (m *BinaryMessage) WriteFloatAsUint16(data float64) error {
	return writeToBuffer(m.buffer, uint16(math.Floor(data)))
}
func (m *BinaryMessage) WriteBytes(data []byte) error {
	_, err := m.buffer.Write(data)
	return err
}
func (m *BinaryMessage) GetBuffer() *bytes.Buffer {
	return m.buffer
}
func (m *BinaryMessage) GetBytes() []byte {
	return m.buffer.Bytes()
}
func (m *BinaryMessage) Reset() {
	m.buffer.Reset()
}
func writeToBuffer(buf *bytes.Buffer, data interface{}) error {
	return binary.Write(buf, binary.LittleEndian, data)
}
