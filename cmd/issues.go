package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// issuesCmd represents the issues command
var issuesCmd = &cobra.Command{
	Use:   "issues",
	Short: "Print issues in current project",
	Long: `Print the list of open issues in the
	current github project`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("issues called")
	},
}

func init() {
	rootCmd.AddCommand(issuesCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// issuesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// issuesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
