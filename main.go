package main

import (
	"github.com/docker/cli/cli-plugins/manager"
	"github.com/docker/cli/cli-plugins/plugin"
	"github.com/docker/cli/cli/command"
	"github.com/spf13/cobra"
)

const MinimalAPIVersion = "1.12"

func main() {
	plugin.Run(func(_ command.Cli) *cobra.Command {
		cmd := &cobra.Command{
			Short: "Docker Reef",
			Long:  `A tool to install docker cli plugins.`,
			Use:   "reef",
			Run:   reef,
		}
		originalPreRun := cmd.PersistentPreRunE
		cmd.PersistentPreRunE = func(cmd *cobra.Command, args []string) error {
			if err := plugin.PersistentPreRunE(cmd, args); err != nil {
				return err
			}
			if originalPreRun != nil {
				return originalPreRun(cmd, args)
			}
			return nil
		}
		return cmd
	}, manager.Metadata{
		SchemaVersion: "0.1.0",
		Vendor:        "Ulysses Souza",
		Version:       "1.0.0",
	})
}

func reef(_ *cobra.Command, _ []string) {
}
