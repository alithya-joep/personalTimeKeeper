/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	myutils "github.com/alithya-joep/personalTimeKeeper/myUtils"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "time proects time",
	Long:  `show the proects and there time for a week.`,
	Run: func(cmd *cobra.Command, args []string) {
		// get refrence to the projects
		projects := myutils.Projects{}
		// load the projects from fie
		if err := projects.Load(projectfile); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		//projects.Print()
		projects.PrintTea()

	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
