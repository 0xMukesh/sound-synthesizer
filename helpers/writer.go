package helpers

import (
	"os"

	"github.com/0xmukesh/sound-synthesizer/constants"
	"github.com/0xmukesh/sound-synthesizer/types"
	"github.com/0xmukesh/sound-synthesizer/utils"
)

type WaveWriter struct{}

func NewWaveWriter() WaveWriter {
	return WaveWriter{}
}

func (w WaveWriter) WriteWaveFile(file string, samples []types.Sample, wavefmt types.WaveFmt) error {
	f, err := os.Create(file)
	if err != nil {
		return err
	}
	defer f.Close()

	var data []byte

	headerBits := utils.CreateHeaderBits(samples, wavefmt)
	data = append(data, headerBits...)

	wfmtInBits := utils.WaveFmtToBits(wavefmt)
	data = append(data, wfmtInBits...)

	data = append(data, []byte(constants.WaveSubChunk2Id)...)
	data = append(data, utils.Int32ToBits(len(samples)*wavefmt.NumOfChannels*wavefmt.BitsPerSample/8)...)

	samplesBits, err := utils.SamplesToBits(samples, wavefmt)
	if err != nil {
		return err
	}
	data = append(data, samplesBits...)

	if _, err := f.Write(data); err != nil {
		return err
	}

	return nil
}
