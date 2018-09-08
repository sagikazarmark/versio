package cmd

import "github.com/spf13/cobra"

// NewVersioCommand creates the `versio` command and its nested children.
func NewVersioCommand(version string) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "versio",
		Version: version,
		Short:   "Tag versions easier",
		Long:    `Versio helps you tagging versions and maintaining version numbers across the project.`,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	cmd.AddCommand(NewCmdBump())

	return cmd
}
