/*
Copyright © 2025 Eric Tian <erictianc@outlook.com>
*/
package cmd

import (
	"fmt"
	"log/slog"
	"os"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

var (
	debug   bool
	depot   string
	logfile *os.File
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "envy",
	Version: "v0.0.1",
	Short:   "A general purpose environment solver",
	Long: `
Envy is designed to setup development/running environment by one-key,
`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
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

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.envy.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	cobra.OnInitialize(initLogger)
	cobra.OnFinalize(func() {
		if logfile != nil {
			logfile.Close()
		}
	})
	rootCmd.PersistentFlags().BoolVar(&debug, "debug", false, "show debug logs")
	rootCmd.PersistentFlags().StringVarP(&depot, "depot", "d", ".envy", "directory for envy")
	// viper.BindPFlag("debug", rootCmd.PersistentFlags().Lookup("debug"))
}

func initLogger() {
	logDir := depot + "/logs/"
	logfileName := time.Now().Format("2006-01-02.15-04-05.000-Mon-Jan.log")
	err := os.MkdirAll(logDir, os.ModePerm)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating log directories.")
		os.Exit(1)
	}

	logfile, err = tea.LogToFile(logDir+logfileName, "")
	// logfile, err = os.Create(logDir + logfileName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating log file")
		os.Exit(1)
	}

	// logger := slog.New(slog.NewTextHandler(logfile, nil))
	// slog.SetDefault(logger)
	if debug {
		slog.SetLogLoggerLevel(slog.LevelDebug)
	}
}
