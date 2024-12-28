package main

import (
	"fmt"
	"math"
	"time"

	"github.com/0xmukesh/sound-synthesizer/constants"
	"github.com/0xmukesh/sound-synthesizer/helpers"
	"github.com/0xmukesh/sound-synthesizer/types"
)

const (
	duration   = 5
	sampleRate = 44100
	freq       = 100
)

func main() {
	// number of samples
	ns := duration * sampleRate
	// angle increment per sample
	angle := (math.Pi * 2.0) / float64(ns)
	// exponential decay; f(x) = (0.0001/x)^(1/ns)
	startAmplitude := 1.0
	endAmplitude := 1.0e-4
	decayFactor := math.Pow(endAmplitude/startAmplitude, 1.0/float64(ns))

	start := time.Now()

	wavefmt := types.WaveFmt{
		SubChunk1Id:   []byte(constants.WaveSubChunk1Id),
		SubChunk1Size: 16,
		AudioFormat:   1,
		NumOfChannels: 2,
		SampleRate:    sampleRate,
		ByteRate:      (sampleRate) * 2 * 16 / 8,
		BlockAlign:    2 * 16 / 8,
		BitsPerSample: 16,
	}

	var samples []types.Sample

	for i := 0; i < ns; i++ {
		sample := types.Sample(math.Sin(angle*freq*float64(i)) * startAmplitude)
		startAmplitude *= decayFactor

		samples = append(samples, sample)
	}

	waveWriter := helpers.NewWaveWriter()
	if err := waveWriter.WriteWaveFile("test.wav", samples, wavefmt); err != nil {
		panic(err.Error())
	}

	fmt.Printf("done - %dms\n", time.Since(start).Milliseconds())
}
