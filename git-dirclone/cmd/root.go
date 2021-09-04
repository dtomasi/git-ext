package cmd

import (
	"github.com/spf13/cobra"
	"github.com/whilp/git-urls"
	"os"
	"os/exec"
	"os/user"
	"path"
	"path/filepath"
	"strings"
)

var (
	// rootCmd represents the base command when called without any subcommands
	rootCmd = newRootCmd()
)

func newRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "git-dirclone",
		Short: "git extension ",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {

			rootDir, err := cmd.PersistentFlags().GetString("root")
			if err != nil {
				return err
			}

			rootDir, err = expandPathWithTilde(rootDir)
			if err != nil {
				return err
			}

			urlObj, err := giturls.Parse(args[0])
			if err != nil {
				return err
			}

			gitCmd := exec.Command("git", "clone", args[0], path.Join(rootDir, urlObj.Host, strings.TrimSuffix(urlObj.Path, ".git")))
			gitCmd.Stdout = os.Stdout
			gitCmd.Stderr = os.Stderr
			return gitCmd.Run()
		},
	}

	cmd.PersistentFlags().StringP("root", "r", os.Getenv("GIT_DIRCLONE_ROOT_DIR"), "root directory. default is environment variable GIT_DIRCLONE_ROOT_DIR")

	return cmd
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func expandPathWithTilde(rootDir string) (string, error) {

	usr, err := user.Current()
	if err != nil {
		return "", err
	}

	dir := usr.HomeDir

	if rootDir == "~" {
		// In case of "~", which won't be caught by the "else if"
		rootDir = dir
	} else if strings.HasPrefix(rootDir, "~/") {
		// Use strings.HasPrefix so we don't match paths like
		// "/something/~/something/"
		rootDir = filepath.Join(dir, rootDir[2:])
	}

	return rootDir, nil
}