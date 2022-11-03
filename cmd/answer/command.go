package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/answerdev/answer/internal/base/conf"
	"github.com/answerdev/answer/internal/cli"
	"github.com/answerdev/answer/internal/install"
	"github.com/answerdev/answer/internal/migrations"
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
	rootCmd.Version = fmt.Sprintf("%s\nrevision: %s\nbuild time: %s", Version, Revision, Time)

	initCmd.Flags().StringVarP(&dataDirPath, "data-path", "C", "/data/", "data path, eg: -C ./data/")

	rootCmd.PersistentFlags().StringVarP(&configFilePath, "config", "c", "", "config path, eg: -c config.yaml")

	dumpCmd.Flags().StringVarP(&dumpDataPath, "path", "p", "./", "dump data path, eg: -p ./dump/data/")

	for _, cmd := range []*cobra.Command{initCmd, checkCmd, runCmd, dumpCmd, upgradeCmd} {
		rootCmd.AddCommand(cmd)
	}
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
		Run: func(_ *cobra.Command, _ []string) {
			runApp()
		},
	}

	// initCmd represents the init command
	initCmd = &cobra.Command{
		Use:   "init",
		Short: "init answer application",
		Long:  `init answer application`,
		Run: func(_ *cobra.Command, _ []string) {
			// check config file and database. if config file exists and database is already created, init done
			cli.InstallAllInitialEnvironment(dataDirPath)
			// set default config file path
			if len(configFilePath) == 0 {
				configFilePath = filepath.Join(cli.ConfigFilePath, cli.DefaultConfigFileName)
			}

			configFileExist := cli.CheckConfigFile(configFilePath)
			if configFileExist {
				fmt.Println("config file exists, try to read the config...")
				c, err := conf.ReadConfig(configFilePath)
				if err != nil {
					fmt.Println("read config failed: ", err.Error())
					return
				}

				fmt.Println("config file read successfully, try to connect database...")
				if cli.CheckDB(c.Data.Database, true) {
					fmt.Println("connect to database successfully and table already exists, do nothing.")
					return
				}
			}

			// start installation server to install
			install.Run(configFilePath)
		},
	}

	// upgradeCmd represents the upgrade command
	upgradeCmd = &cobra.Command{
		Use:   "upgrade",
		Short: "upgrade Answer version",
		Long:  `upgrade Answer version`,
		Run: func(_ *cobra.Command, _ []string) {
			c, err := conf.ReadConfig(configFilePath)
			if err != nil {
				fmt.Println("read config failed: ", err.Error())
				return
			}
			if err = migrations.Migrate(c.Data.Database); err != nil {
				fmt.Println("migrate failed: ", err.Error())
				return
			}
			fmt.Println("upgrade done")
		},
	}

	// dumpCmd represents the dump command
	dumpCmd = &cobra.Command{
		Use:   "dump",
		Short: "back up data",
		Long:  `back up data`,
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Println("Answer is backing up data")
			c, err := conf.ReadConfig(configFilePath)
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
		Run: func(_ *cobra.Command, _ []string) {
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

			c, err := conf.ReadConfig(configFilePath)
			if err != nil {
				fmt.Println("read config failed: ", err.Error())
				return
			}

			if cli.CheckDB(c.Data.Database, false) {
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
