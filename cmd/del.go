package cmd

import (
	"fmt"
	"os"

	myutils "github.com/alithya-joep/personalTimeKeeper/myUtils"
	"github.com/spf13/cobra"
)

// delCmd represents the del command
var delCmd = &cobra.Command{
	Use:   "del",
	Short: "delete project line",
	Long: `remove line by index
	`,
	Run: func(cmd *cobra.Command, args []string) {
		row, _ := cmd.Flags().GetInt("delete")

		// get refrence to the projects
		projects := myutils.Projects{}
		// load the projects from fie
		if err := projects.Load(projectfile); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		projects.Delete(row)
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
	rootCmd.AddCommand(delCmd)
	delCmd.Flags().IntP("delete", "r", 0, "delete a row")
}
