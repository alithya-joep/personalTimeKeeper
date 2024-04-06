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
	ptt  add -p new project -t code review -c INT1577 -d 2024-04-07`,

	Run: func(cmd *cobra.Command, args []string) {
		projectname, _ := cmd.Flags().GetString("project")
		task, _ := cmd.Flags().GetString("task")
		comment, _ := cmd.Flags().GetString("comment")
		date, _ := cmd.Flags().GetString("date")
		// get refrence to the projects
		projects := myutils.Projects{}

		// load the projects from fie
		if err := projects.Load(projectfile); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		projects.Add(projectname, task, comment, date)
		// store the projects back to file
		err := projects.Store(projectfile)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		projects.Print()
		fmt.Println("add called")
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().StringP("project", "p", "", "project name")
	addCmd.Flags().StringP("task", "t", "", "task code")
	addCmd.Flags().StringP("comment", "c", "", "comment like 'ZZZ1234'")
	addCmd.Flags().StringP("date", "d", "", "date for start of week.")
}
