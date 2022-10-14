package main

import (
	"fmt"
	"os"

	"github.com/segmentfault/answer/internal/cli"
	"github.com/spf13/cobra"
)

var (
	// configFilePath is the config file path
	configFilePath string
	// dataDirPath save all answer application data in this directory. like config file, upload file...
	dataDirPath string
	// dumpDataPath dump data path
	dumpDataPath string
)

func init() {
	rootCmd.Version = Version

	initCmd.Flags().StringVarP(&dataDirPath, "data-path", "C", "/data/", "data path, eg: -C ./data/")

	runCmd.Flags().StringVarP(&configFilePath, "config", "c", "", "config path, eg: -c config.yaml")

	dumpCmd.Flags().StringVarP(&dumpDataPath, "path", "p", "./", "dump data path, eg: -p ./dump/data/")

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
			cli.InstallAllInitialEnvironment(dataDirPath)
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
			c, err := readConfig()
			if err != nil {
				fmt.Println("read config failed: ", err.Error())
				return
			}
			err = cli.DumpAllData(c.Data.Database, dumpDataPath)
			if err != nil {
				fmt.Println("dump failed: ", err.Error())
				return
			}
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
			if cli.CheckConfigFile(configFilePath) {
				fmt.Println("config file exists [✔]")
			} else {
				fmt.Println("config file not exists [x]")
			}

			if cli.CheckUploadDir() {
				fmt.Println("upload directory exists [✔]")
			} else {
				fmt.Println("upload directory not exists [x]")
			}

			c, err := readConfig()
			if err != nil {
				fmt.Println("read config failed: ", err.Error())
				return
			}

			if cli.CheckDB(c.Data.Database) {
				fmt.Println("db connection successfully [✔]")
			} else {
				fmt.Println("db connection failed [x]")
			}
			fmt.Println("check environment all done")
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
