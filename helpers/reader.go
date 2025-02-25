package helpers

import (
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/0xmukesh/sound-synthesizer/constants"
	"github.com/0xmukesh/sound-synthesizer/types"
	"github.com/0xmukesh/sound-synthesizer/utils"
)

type WaveReader struct{}

func NewWaveReader() WaveReader {
	return WaveReader{}
}

func (r WaveReader) ParseFile(file string) (types.Wave, error) {
	f, err := os.Open(file)
	if err != nil {
		return types.Wave{}, err
	}
	defer f.Close()

	data, err := io.ReadAll(f)
	if err != nil {
		return types.Wave{}, err
	}

	header, err := r.parseHeader(data)
	if err != nil {
		return types.Wave{}, err
	}

	wavefmt, err := r.parseWaveFmt(data)
	if err != nil {
		return types.Wave{}, err
	}

	samples, err := r.parseData(data)
	if err != nil {
		return types.Wave{}, err
	}

	wave := types.Wave{
		WaveHeader: header,
		WaveFmt:    wavefmt,
		Samples:    samples,
	}

	return wave, nil
}

func (r WaveReader) parseHeader(data []byte) (types.WaveHeader, error) {
	header := types.WaveHeader{}

	chunkId := data[0:4]
	if string(chunkId) != constants.WaveChunkId {
		return header, errors.New("invalid file")
	}
	header.ChunkId = chunkId

	chunkSize := data[4:8]
	header.ChunkSize = utils.Bits32ToInt(chunkSize)

	format := data[8:12]
	if string(format) != constants.WaveFileFormat {
		return header, errors.New("invalid format")
	}

	return header, nil
}

func (r WaveReader) parseWaveFmt(data []byte) (types.WaveFmt, error) {
	wavefmt := types.WaveFmt{}

	subChunk1Id := data[12:16]
	if string(subChunk1Id) != constants.WaveSubChunk1Id {
		return wavefmt, fmt.Errorf("invalid sub chunk 1 id - %s", string(subChunk1Id))
	}

	wavefmt.SubChunk1Id = subChunk1Id
	wavefmt.SubChunk1Size = utils.Bits32ToInt(data[16:20])
	wavefmt.AudioFormat = utils.Bits16ToInt(data[20:22])
	wavefmt.NumOfChannels = utils.Bits16ToInt(data[22:24])
	wavefmt.SampleRate = utils.Bits32ToInt(data[24:28])
	wavefmt.ByteRate = utils.Bits32ToInt(data[28:32])
	wavefmt.BlockAlign = utils.Bits16ToInt(data[32:34])
	wavefmt.BitsPerSample = utils.Bits16ToInt(data[34:36])

	return wavefmt, nil
}

func (r WaveReader) parseData(data []byte) ([]types.Sample, error) {
	wavefmt, err := r.parseWaveFmt(data)
	if err != nil {
		return nil, err
	}

	subChunk2Id := data[36:40]
	if string(subChunk2Id) != constants.WaveSubChunk2Id {
		return nil, fmt.Errorf("invalid sub chunk 2 id - %s", string(subChunk2Id))
	}

	bytesPerSampleSize := wavefmt.BitsPerSample / 8
	rawData := data[44:]

	samples := []types.Sample{}

	for i := 0; i < len(rawData); i += bytesPerSampleSize {
		rawSample := rawData[i : i+bytesPerSampleSize]
		unscaledSample := utils.BitsToInt(rawSample, wavefmt.BitsPerSample)
		scaledSample := types.Sample(float64(unscaledSample) / float64(utils.MaxValue(wavefmt.BitsPerSample)))
		samples = append(samples, scaledSample)
	}

	return samples, nil
}
