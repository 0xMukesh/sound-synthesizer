package cmd

import (
	"context"

	"github.com/0xmukesh/sound-synthesizer/commands"
	"github.com/spf13/cobra"
)

func Execute() error {
	rootCmd := &cobra.Command{
		Version: "0.0.1",
		Use:     "ss",
		Long:    "`ss` is a simple sound sythesizer",
		Example: "ss",
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
	}

	genCmd := commands.GenerateCmd{}
	amplifyCmd := commands.AmplifyCmd{}
	stereoPanCmd := commands.StereoPanCmd{}

	rootCmd.AddCommand(genCmd.Command())
	rootCmd.AddCommand(amplifyCmd.Command())
	rootCmd.AddCommand(stereoPanCmd.Command())

	return rootCmd.ExecuteContext(context.Background())
}
