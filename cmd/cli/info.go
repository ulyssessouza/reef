package main

import (
	"errors"
	"fmt"
	"net/url"

	"github.com/spf13/cobra"
)

func isURL(arg string) bool {
	_, err := url.ParseRequestURI(arg)
	return err == nil
}

var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Provides informations about a puglin",
	Long:  "Retrieves the .reef.yml from the repository and displays the informations.",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 || !isURL(args[0]) {
			return errors.New("requires an url")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("info", args[0])
	},
}

func init() {
	rootCmd.AddCommand(infoCmd)
}
