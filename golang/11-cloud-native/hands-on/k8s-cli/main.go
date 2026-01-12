package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "k8s-cli",
		Short: "A minimal Kubernetes CLI clone",
		Long:  `A hands-on project demonstrating Cobra for building cloud-native tools.`,
	}

	// 1. Add "get" command
	var getCmd = &cobra.Command{
		Use:   "get",
		Short: "Display one or many resources",
	}

	// 2. Add "pods" subcommand to "get"
	var podsCmd = &cobra.Command{
		Use:   "pods",
		Short: "List pods",
		Run: func(cmd *cobra.Command, args []string) {
			// In a real app, we'd use k8s client-go here
			fmt.Println("NAME             STATUS    RESTARTS   AGE")
			fmt.Println("nginx-deployment 3/3       0          14m")
			fmt.Println("api-gateway      1/1       0          5h")
		},
	}

	// 3. Flags
	var verbose bool
	podsCmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "Enable verbose output")

	getCmd.AddCommand(podsCmd)
	rootCmd.AddCommand(getCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
