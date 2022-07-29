/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

func RunE (cmd *cobra.Command, args []string) error {
	leftArg := args[0]
	rightArg := args[1]
	rest := args[2:]
	socatargs := append(rest, leftArg, rightArg)
	oscmd := exec.Command("socat", socatargs...)
	return oscmd.Start()
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "porter_socat_bin",
	Short: "Run Socat as a background executable",
	RunE: RunE,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.porter_socat_bin.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("debug", "v", false, "Enable verbose output")
}

// YUCK! Extract arguments separated by -- necessary to accept unknown flags
// And pass them to socat. See https://github.com/spf13/cobra/issues/739
func extractUnknownArgs(flags *pflag.FlagSet, args []string) []string {
	unknownArgs := []string{}

	for i := 0; i < len(args); i++ {
		a := args[i]
		var f *pflag.Flag
		if a[0] == '-' {
			if len(a) > 1 && a[1] == '-' {
				f = flags.Lookup(strings.SplitN(a[2:], "=", 2)[0])
			} else if len(a) > 1 {
				for _, s := range a[1:] {
					f = flags.ShorthandLookup(string(s))
					if f == nil {
						break
					}
				}
			} else {
				continue
			}
		} else {
			continue
		}
		if f != nil {
			if f.NoOptDefVal == "" && i+1 < len(args) && f.Value.String() == args[i+1] {
				i++
			}
			continue
		}
		unknownArgs = append(unknownArgs, a)
	}
	return unknownArgs
}
