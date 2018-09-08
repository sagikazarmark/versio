package cmd

import (
	"errors"
	"fmt"

	"github.com/Masterminds/semver"
	"github.com/spf13/cobra"
)

type bumpOptions struct {
	currentVersion string
}

func NewCmdBump() *cobra.Command {
	options := &bumpOptions{}

	cmd := &cobra.Command{
		Use:   "bump (major|minor|patch)",
		Short: "Bump (and tag) a version",
		RunE: func(cmd *cobra.Command, args []string) error {
			return options.run(args[0])
		},
		ValidArgs: []string{"major", "minor", "patch"},
		Args:      cobra.ExactArgs(1),
	}

	cmd.Flags().StringVar(&options.currentVersion, "current", "", "current version")
	cmd.MarkFlagRequired("current") // Until current version can come from configuration

	return cmd
}

func (o *bumpOptions) run(bumpType string) error {
	version, err := semver.NewVersion(o.currentVersion)
	if err != nil {
		return err
	}

	var newVersion semver.Version

	switch bumpType {
	case "major":
		newVersion = version.IncMajor()

	case "minor":
		newVersion = version.IncMinor()

	case "patch":
		newVersion = version.IncPatch()

	default:
		return errors.New("unknown bump type")
	}

	fmt.Println(newVersion)

	return nil
}
