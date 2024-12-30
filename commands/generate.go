package commands

import (
	"fmt"
	"math"
	"os"
	"time"

	"github.com/0xmukesh/sound-synthesizer/constants"
	"github.com/0xmukesh/sound-synthesizer/helpers"
	"github.com/0xmukesh/sound-synthesizer/types"
	"github.com/spf13/cobra"
)

type GenerateCmd struct{}

func (c GenerateCmd) Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "generate",
		Short:   "generates a pure sine wave with a constant frequency and saves it a .wave file",
		Example: "ss generate",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := c.Handler(cmd); err != nil {
				fmt.Println(err.Error())
				os.Exit(1)
			}

			return nil
		},
	}

	cmd.Flags().Float64("freq", 440, "freq of the sine wave")
	cmd.Flags().Int("duration", 5, "duration of the audio")
	cmd.Flags().Int("sample_rate", 44100, "sample rate which is to be used while sampling")
	cmd.Flags().Int("num_channels", 2, "number of audio channels")
	cmd.Flags().Int("bits_per_sample", 16, "number of bits per sample")
	cmd.Flags().String("filename", "rec.wav", "name of the file where the audio would be saved")

	return cmd
}

func (c GenerateCmd) Handler(cmd *cobra.Command) error {
	freq, _ := cmd.Flags().GetFloat64("freq")
	duration, _ := cmd.Flags().GetInt("duration")
	sampleRate, _ := cmd.Flags().GetInt("sample_rate")
	numOfChannels, _ := cmd.Flags().GetInt("num_channels")
	bitsPerSample, _ := cmd.Flags().GetInt("bits_per_sample")
	fileName, _ := cmd.Flags().GetString("filename")

	ns := duration * sampleRate
	angle := (math.Pi * 2.0) / float64(ns)

	startAmplitude := 1.0
	endAmplitude := 1.0e-2
	decayFactor := math.Pow(endAmplitude/startAmplitude, 1.0/float64(ns))

	start := time.Now()

	wavefmt := types.WaveFmt{
		SubChunk1Id:   []byte(constants.WaveSubChunk1Id),
		SubChunk1Size: 16,
		AudioFormat:   1,
		NumOfChannels: numOfChannels,
		SampleRate:    sampleRate,
		ByteRate:      (sampleRate) * (numOfChannels) * (bitsPerSample) / 8,
		BlockAlign:    (numOfChannels) * (bitsPerSample) / 8,
		BitsPerSample: (bitsPerSample),
	}

	var samples []types.Sample

	for i := 0; i < ns; i++ {
		sample := types.Sample(math.Sin(angle*freq*float64(i)) * startAmplitude)
		startAmplitude *= decayFactor

		samples = append(samples, sample)
	}

	waveWriter := helpers.NewWaveWriter()
	if err := waveWriter.WriteWaveFile(fileName, samples, wavefmt); err != nil {
		return err
	}

	fmt.Printf("done - %dms\n", time.Since(start).Milliseconds())
	return nil
}
