/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "tr",
	Short: "Translate characters",
	Long:  `tr is a command-line utility for translating characters i.e. deleting, squeezing or replacing characters from given text.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 2 || len(args) > 2 {
			fmt.Fprint(cmd.OutOrStdout(), HELPSTRING)
			return
		}

		p1, p2 := args[0], args[1]
		if isValid, err := isValidRangePattern(p1); !isValid || err != nil {
			fmt.Fprint(cmd.OutOrStdout(), HELPSTRING)
			return
		}
		if isValid, err := isValidRangePattern(p2); !isValid || err != nil {
			fmt.Fprint(cmd.OutOrStdout(), HELPSTRING)
			return
		}

		r1, r2 := createRangeFromPattern(p1), createRangeFromPattern(p2)

		rangeSubstitutions := createRangeSubstitutions(r1, r2)

		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			inp := scanner.Text()
			fmt.Printf("%s\n", substitute(inp, rangeSubstitutions))
		}
	},
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

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.tr.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
