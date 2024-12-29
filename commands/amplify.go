package commands

import (
	"fmt"
	"os"

	"github.com/0xmukesh/sound-synthesizer/helpers"
	"github.com/0xmukesh/sound-synthesizer/types"
	"github.com/spf13/cobra"
)

type AmplifyCmd struct{}

func NewAmplifyCmd() AmplifyCmd {
	return AmplifyCmd{}
}

func (c AmplifyCmd) Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "amplify",
		Short:   "changes amplitude of an input wave file and saves it to another file",
		Example: "ss amplify",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := c.Handler(cmd, args); err != nil {
				fmt.Println(err.Error())
				os.Exit(1)
			}

			return nil
		},
	}

	cmd.Flags().String("input", "", "input file (required)")
	cmd.Flags().String("output", "", "output file (required)")
	cmd.Flags().Float64("scale_factor", 1.0, "scale factor")

	cmd.MarkFlagRequired("input")
	cmd.MarkFlagRequired("output")

	return cmd
}

func (c AmplifyCmd) Handler(cmd *cobra.Command, args []string) error {
	input, _ := cmd.Flags().GetString("input")
	output, _ := cmd.Flags().GetString("output")
	scaleFactor, _ := cmd.Flags().GetFloat64("scale_factor")

	waveReader := helpers.NewWaveReader()
	waveWriter := helpers.NewWaveWriter()

	wave, err := waveReader.ParseFile(input)
	if err != nil {
		return err
	}

	var updatedSamples []types.Sample

	for _, sample := range wave.Samples {
		updatedSample := types.Sample(float64(sample) * scaleFactor)
		updatedSamples = append(updatedSamples, updatedSample)
	}

	if err := waveWriter.WriteWaveFile(output, updatedSamples, wave.WaveFmt); err != nil {
		return err
	}

	return nil
}
