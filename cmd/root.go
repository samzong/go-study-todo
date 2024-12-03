/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
// Package declaration for the cmd package
package cmd

import (
	// Importing the os package for handling operating system functionalities
	"os"

	// Importing the cobra package for creating CLI applications
	"github.com/spf13/cobra"
)

// Declaring variables for configuration file, package base, user license, and AWS region
var (
	// cfgFile stores the path to the configuration file
	cfgFile string
	// packageBase stores the base path of the package
	packageBase string
	// userLicense stores the user license information
	userLicense string
	// Region stores the AWS region, defaulting to "us-west-2"
	Region string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	// Use specifies the command name
	Use: "todo",
	// Short provides a brief description of the application
	Short: "A brief description of your application",
	// Long provides a longer description of the application, including examples and usage
	Long: "A longer description that spans multiple lines and likely contains examples and usage of using your application.",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This function is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	// Execute the root command and handle any errors
	err := rootCmd.Execute()
	if err != nil {
		// Exit the program with status code 1 if there is an error
		os.Exit(1)
	}
}

// init function initializes the flags and configuration settings
func init() {
	// Adding a string flag for the configuration file with shorthand "c"
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.todo.yaml)")
	// Adding a string flag for the AWS region with shorthand "r" and default value "us-west-2"
	rootCmd.Flags().StringVarP(&Region, "region", "r", "us-west-2", "AWS region <required>")
	// Marking the region flag as required
	rootCmd.MarkFlagsOneRequired("region")
	rootCmd.TraverseChildren = false
}
