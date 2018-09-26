package cmd

import (
	"context"
	"os"

	"github.com/gobuffalo/genny"
	"github.com/gobuffalo/release/genny/initgen"
	"github.com/markbates/gonit/genny/gonit"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

var rootOptions = struct {
	*gonit.Options
	dryRun bool
}{
	Options: &gonit.Options{
		Options: &initgen.Options{},
	},
}

// rootCmd represents the init command
var rootCmd = &cobra.Command{
	Use:   "goinit",
	Short: "setups a project to use go",
	RunE: func(cmd *cobra.Command, args []string) error {
		var run *genny.Runner = genny.WetRunner(context.Background())
		if rootOptions.dryRun {
			run = genny.DryRunner(context.Background())
		}

		opts := rootOptions.Options

		gg, err := gonit.New(opts)
		if err != nil {
			return errors.WithStack(err)
		}
		run.WithGroup(gg)

		return run.Run()
	},
}

func init() {
	rootCmd.Flags().BoolVarP(&rootOptions.dryRun, "dry-run", "d", false, "runs the generator dry")
	rootCmd.Flags().BoolVarP(&rootOptions.Force, "force", "f", false, "force files to overwrite existing ones")
	rootCmd.Flags().StringVarP(&rootOptions.MainFile, "main-file", "m", "", "adds a .goreleaser.yml file (only for binary applications)")
	rootCmd.Flags().StringVarP(&rootOptions.VersionFile, "version-file", "v", "version.go", "path to a version file to maintain")
}
