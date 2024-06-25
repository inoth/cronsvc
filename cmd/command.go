package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "cronsvc",
		Short: "",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("cronsvc is starting..........................")
			runApp()
		},
	}
)
