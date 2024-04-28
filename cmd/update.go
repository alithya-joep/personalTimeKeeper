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

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "update a rows hours",
	Long: `select a row and update a value in the daily hours. For example:

	ptt -r 3 -d mon -v 4`,
	Run: func(cmd *cobra.Command, args []string) {
		row, _ := cmd.Flags().GetInt("row")
		newValue, _ := cmd.Flags().GetUint8("newValue")
		day, _ := cmd.Flags().GetString("day")

		fmt.Println("Row:", row)
		fmt.Println("New Value:", newValue)
		fmt.Println("Day:", day)

		// get refrence to the projects
		projects := myutils.Projects{}

		// load the projects from fie
		if err := projects.Load(projectfile); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		projects.Update(row, day, newValue)
		// store the projects back to file
		err := projects.Store(projectfile)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		projects.Print()

	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
	updateCmd.Flags().IntP("row", "r", 0, "select row to update")
	updateCmd.Flags().Uint8P("newValue", "v", 0, "new value for cell")
	updateCmd.Flags().StringP("day", "d", "", "select day like thu")
}
