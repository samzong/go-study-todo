/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var (
	cfgFile     string
	packageBase string
	userLicense string
	Region      string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "todo",
	Short: "A brief description of your application",
	Long:  "A longer description that spans multiple lines and likely contains examples and usage of using your application.",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringP("config", "c", "", "config file")
	rootCmd.Flags().StringVarP(&Region, "region", "r", "us-west-2", "AWS region")
	rootCmd.MarkFlagsOneRequired("region")
}
