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

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add new project",
	Short: "add a new project to the list of projects",
	Long: `add  new project task for the week. For example:
	personalTimeKeeper add "new project"`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("add called")
		projectname, _ := cmd.Flags().GetString("project")
		task, _ := cmd.Flags().GetString("task")
		comment, _ := cmd.Flags().GetString("comment")
		// get refrence to the projects
		projects := myutils.Projects{}

		// load the projects from fie
		if err := projects.Load(projectfile); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		projects.Add(projectname, task, comment)
		// store the projects back to file
		err := projects.Store(projectfile)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	addCmd.Flags().StringP("project", "p", "", "project name")
	addCmd.Flags().StringP("task", "t", "", "task code")
	addCmd.Flags().StringP("comment", "c", "", "comment like 'ZZZ1234'")
}
