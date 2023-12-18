/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package patterns

import (
	"github.com/jipilmuk/orb/utills"
	"github.com/spf13/cobra"
)

// webVitalsPtrnsCmd represents the webVitalPtrns command.
var webVitalsPtrnsCmd = &cobra.Command{
	Use:   "webVitalsPtrns",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		utills.InitFiddleForPatterns("/_pattern/web-vitals-patterns/")
	},
}

func init() {
	patternCmd.AddCommand(webVitalsPtrnsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// webVitalPtrnsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// webVitalPtrnsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
