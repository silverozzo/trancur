package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"trancur/app"
	"trancur/helper"
)

var rootCmd = &cobra.Command{
	Use: "main",
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		ctx, cancel := context.WithCancel(ctx)
		defer cancel()

		go func() {
			helper.AwaitSignal(ctx)
			cancel()
		}()

		app.Run(ctx)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
