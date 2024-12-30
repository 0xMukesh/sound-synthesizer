package commands

import (
	"fmt"
	"os"

	"github.com/0xmukesh/sound-synthesizer/helpers"
	"github.com/0xmukesh/sound-synthesizer/types"
	"github.com/0xmukesh/sound-synthesizer/utils"
	"github.com/spf13/cobra"
)

type StereoPanCmd struct{}

func (c StereoPanCmd) Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "stereopan",
		Short:   "takes in a mono audio file and does panning and returns a stereo audio file",
		Example: "ss stereopan",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := c.Handler(cmd); err != nil {
				fmt.Println(err.Error())
				os.Exit(1)
			}

			return nil
		},
	}

	cmd.Flags().String("input", "", "input file (required)")
	cmd.Flags().String("output", "", "output file (required)")
	cmd.Flags().Float64("panning_position", 0.0, "panning position")

	cmd.MarkFlagRequired("input")
	cmd.MarkFlagRequired("output")

	return cmd
}

func (c StereoPanCmd) Handler(cmd *cobra.Command) error {
	input, _ := cmd.Flags().GetString("input")
	output, _ := cmd.Flags().GetString("output")
	panningPosition, _ := cmd.Flags().GetFloat64("panning_position")

	waveReader := helpers.NewWaveReader()
	waveWriter := helpers.NewWaveWriter()

	wave, err := waveReader.ParseFile(input)
	if err != nil {
		return err
	}

	leftChanMultiplier, rightChanMultiplier := utils.PanPositionToChanMultipliers(panningPosition)

	var updatedSamples []types.Sample

	for _, sample := range wave.Samples {
		updatedSamples = append(updatedSamples, types.Sample(sample.ToFloat()*leftChanMultiplier))
		updatedSamples = append(updatedSamples, types.Sample(sample.ToFloat()*rightChanMultiplier))
	}

	wave.WaveFmt.NumOfChannels = 2

	if err := waveWriter.WriteWaveFile(output, updatedSamples, wave.WaveFmt); err != nil {
		return err
	}

	return nil
}
