/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package fiddle

import (
	"github.com/jipilmuk/orb/utills"
	"github.com/spf13/cobra"
)

// prsntCmd represents the prsnt command.
var prsntCmd = &cobra.Command{
	Use:   "prsnt",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		utills.InitComponents("/_prsnt")
	},
}

func init() {
	FiddleCmd.AddCommand(prsntCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// prsntCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// prsntCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
