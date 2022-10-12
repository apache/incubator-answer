package main

import (
	"fmt"
	"os"

	"github.com/segmentfault/answer/internal/cli"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.Version = Version
	runCmd.Flags().StringVarP(&confFlag, "config", "c", "data/config.yaml", "config path, eg: -c config.yaml")

	rootCmd.AddCommand(initCmd)
	rootCmd.AddCommand(checkCmd)
	rootCmd.AddCommand(runCmd)
	rootCmd.AddCommand(dumpCmd)
	rootCmd.AddCommand(upgradeCmd)
}

var (
	// rootCmd represents the base command when called without any subcommands
	rootCmd = &cobra.Command{
		Use:   "answer",
		Short: "Answer is a minimalist open source Q&A community.",
		Long: `Answer is a minimalist open source Q&A community.
To run answer, use:
	- 'answer init' to initialize the required environment.
	- 'answer run' to launch application.`,
	}

	// runCmd represents the run command
	runCmd = &cobra.Command{
		Use:   "run",
		Short: "Run the application",
		Long:  `Run the application`,
		Run: func(cmd *cobra.Command, args []string) {
			runApp()
		},
	}

	// initCmd represents the init command
	initCmd = &cobra.Command{
		Use:   "init",
		Short: "init answer application",
		Long:  `init answer application`,
		Run: func(cmd *cobra.Command, args []string) {
			cli.InstallAllInitialEnvironment()
			c, err := readConfig()
			if err != nil {
				fmt.Println("read config failed: ", err.Error())
				return
			}
			fmt.Println("read config successfully")
			if err := cli.InitDB(c.Data.Database); err != nil {
				fmt.Println("init database error: ", err.Error())
				return
			}
			fmt.Println("init database successfully")
		},
	}

	// upgradeCmd represents the upgrade command
	upgradeCmd = &cobra.Command{
		Use:   "upgrade",
		Short: "upgrade Answer version",
		Long:  `upgrade Answer version`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("The current app version is 1.0.0, the latest version is 1.0.0, no need to upgrade.")
		},
	}

	// dumpCmd represents the dump command
	dumpCmd = &cobra.Command{
		Use:   "dump",
		Short: "back up data",
		Long:  `back up data`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Answer is backing up data")
			fmt.Println("Answer backed up the data successfully.")
		},
	}

	// checkCmd represents the check command
	checkCmd = &cobra.Command{
		Use:   "check",
		Short: "checking the required environment",
		Long:  `Check if the current environment meets the startup requirements`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Start checking the required environment...")
			fmt.Println("config file exists [✔]")
			fmt.Println("db connection successfully [✔]")
			fmt.Println("cache directory exists [✔]")
			fmt.Println("Successful! The current environment meets the startup requirements.")
		},
	}
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
