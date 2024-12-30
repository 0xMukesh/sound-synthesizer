package utils

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math"
)

func BitsToInt(b []byte, size int) int {
	switch size {
	case 16:
		return Bits16ToInt(b)
	case 32:
		return Bits32ToInt(b)
	default:
		panic("invalid size. only 16 and 32 bits are accepted")
	}
}

func Bits16ToInt(b []byte) int {
	if len(b) != 2 {
		panic(fmt.Errorf("invalid size. expected 2, got %d", len(b)))
	}

	var payload int16
	buf := bytes.NewReader(b)
	if err := binary.Read(buf, binary.LittleEndian, &payload); err != nil {
		panic(err.Error())
	}

	return int(payload)
}

func Bits32ToInt(b []byte) int {
	if len(b) != 4 {
		panic(fmt.Errorf("invalid size. expected 4, got %d", len(b)))
	}

	var payload int32
	buf := bytes.NewReader(b)
	if err := binary.Read(buf, binary.LittleEndian, &payload); err != nil {
		panic(err.Error())
	}

	return int(payload)
}

func IntToBits(i int, size int) []byte {
	switch size {
	case 16:
		return Int16ToBits(i)
	case 32:
		return Int32ToBits(i)
	default:
		panic("invalid size. only 16 and 32 bits are accepted")
	}
}

func Int16ToBits(i int) []byte {
	b := make([]byte, 2)
	binary.LittleEndian.PutUint16(b, uint16(i))
	return b
}

func Int32ToBits(i int) []byte {
	b := make([]byte, 4)
	binary.LittleEndian.PutUint32(b, uint32(i))
	return b
}

func FloatToBits(f float64, size int) []byte {
	bits := math.Float64bits(f)
	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, bits)

	switch size {
	case 2:
		return b[:2]
	case 4:
		return b[:4]
	}

	return b
}

func BitsToFloat(b []byte) float64 {
	switch len(b) {
	case 4:
		bits32 := binary.LittleEndian.Uint32(b)
		return float64(math.Float32frombits(bits32))
	case 8:
		bits64 := binary.LittleEndian.Uint64(b)
		return math.Float64frombits(bits64)
	default:
		panic(fmt.Errorf("invalid size: %d, must be 32 or 64 bits", len(b)*8))
	}
}

func MaxValue(numOfBits int) int {
	var max int

	switch numOfBits {
	case 8:
		max = math.MaxInt8
	case 16:
		max = math.MaxInt16
	case 32:
		max = math.MaxInt32
	case 64:
		max = math.MaxInt64
	default:
		panic(fmt.Errorf("invalid size - %d, must be 8, 16, 32 or 64-bits only", numOfBits))
	}

	return max
}

func PanPositionToChanMultipliers(p float64) (float64, float64) {
	if !(p >= -1 && p <= 1) {
		panic("pan position outside [-1, 1] range")
	}

	leftChanMultiplier := (p / 2) - 0.5
	rightChanMultiplier := (p / 2) + 0.5

	return leftChanMultiplier, rightChanMultiplier
}
