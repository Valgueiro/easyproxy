package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init(){
	rootCmd.AddCommand(versionCmd)
}

func run(cobra *cobra.Command, args []string){
	fmt.Println("Version - 0.0.1-alpha")
}

var versionCmd = &cobra.Command{
	Use: "version",
	Short: "Print relevant versions",
	Run: run,
}
