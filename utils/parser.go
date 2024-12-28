package utils

import (
	"fmt"
	"math"

	"github.com/0xmukesh/sound-synthesizer/constants"
	"github.com/0xmukesh/sound-synthesizer/types"
)

func WaveFmtToBits(wfmt types.WaveFmt) []byte {
	var b []byte

	b = append(b, wfmt.SubChunk1Id...)
	b = append(b, Int32ToBits(wfmt.SubChunk1Size)...)
	b = append(b, Int16ToBits(wfmt.AudioFormat)...)
	b = append(b, Int16ToBits(wfmt.NumOfChannels)...)
	b = append(b, Int32ToBits(wfmt.SampleRate)...)
	b = append(b, Int32ToBits(wfmt.ByteRate)...)
	b = append(b, Int16ToBits(wfmt.BlockAlign)...)
	b = append(b, Int16ToBits(wfmt.BitsPerSample)...)

	return b
}

func SamplesToBits(samples []types.Sample, wfmt types.WaveFmt) ([]byte, error) {
	var b []byte

	for _, s := range samples {
		var multiplier int

		switch wfmt.BitsPerSample {
		case 8:
			multiplier = math.MaxInt8
		case 16:
			multiplier = math.MaxInt16
		case 32:
			multiplier = math.MaxInt32
		case 64:
			multiplier = math.MaxInt64
		default:
			return nil, fmt.Errorf("invalid size - %d, must be 8, 16, 32 or 64-bits only", wfmt.BitsPerSample)
		}

		bits := IntToBits(int(float64(s)*float64(multiplier)), wfmt.BitsPerSample)
		b = append(b, bits...)
	}

	return b, nil
}

func CreateHeaderBits(samples []types.Sample, wfmt types.WaveFmt) []byte {
	var b []byte

	chunkSizeInBits := Int32ToBits(36 + (len(samples)*wfmt.NumOfChannels*wfmt.BitsPerSample)/8)

	b = append(b, []byte(constants.WaveChunkId)...)
	b = append(b, chunkSizeInBits...)
	b = append(b, []byte(constants.WaveFileFormat)...)

	return b
}
