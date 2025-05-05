/*
Copyright © 2025 Eric Tian <erictianc@outlook.com>
*/
package cmd

import (
	"log/slog"

	"github.com/spf13/cobra"
)

// checkCmd represents the check command
var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Check current environment",
	Long:  `Use different checker to check current environment wisely.`,
	Run: func(cmd *cobra.Command, args []string) {
		slog.Debug("running check command:", "arguments", args)
		// tea.LogToFile("debug.log", "debug")
	},
}

func init() {
	rootCmd.AddCommand(checkCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// checkCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// checkCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
