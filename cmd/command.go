package answercmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/answerdev/answer/internal/base/conf"
	"github.com/answerdev/answer/internal/cli"
	"github.com/answerdev/answer/internal/install"
	"github.com/answerdev/answer/internal/migrations"
	"github.com/answerdev/answer/plugin"
	"github.com/segmentfault/pacman/log"
	"github.com/spf13/cobra"
)

var (
	// dataDirPath save all answer application data in this directory. like config file, upload file...
	dataDirPath string
	// dumpDataPath dump data path
	dumpDataPath string
	// plugins needed to build in answer application
	buildWithPlugins []string
	// build output path
	buildOutput string
	// This config is used to upgrade the database from a specific version manually.
	// If you want to upgrade the database to version 1.1.0, you can use `answer upgrade -f v1.1.0`.
	upgradeVersion string
)

func init() {
	rootCmd.Version = fmt.Sprintf("%s\nrevision: %s\nbuild time: %s", Version, Revision, Time)

	rootCmd.PersistentFlags().StringVarP(&dataDirPath, "data-path", "C", "/data/", "data path, eg: -C ./data/")

	dumpCmd.Flags().StringVarP(&dumpDataPath, "path", "p", "./", "dump data path, eg: -p ./dump/data/")

	buildCmd.Flags().StringSliceVarP(&buildWithPlugins, "with", "w", []string{}, "plugins needed to build")

	buildCmd.Flags().StringVarP(&buildOutput, "output", "o", "", "build output path")

	upgradeCmd.Flags().StringVarP(&upgradeVersion, "from", "f", "", "upgrade from specific version, eg: -f v1.1.0")

	for _, cmd := range []*cobra.Command{initCmd, checkCmd, runCmd, dumpCmd, upgradeCmd, buildCmd, pluginCmd} {
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
			cli.FormatAllPath(dataDirPath)
			fmt.Println("config file path: ", cli.GetConfigFilePath())
			fmt.Println("Answer is string..........................")
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

			configFileExist := cli.CheckConfigFile(cli.GetConfigFilePath())
			if configFileExist {
				fmt.Println("config file exists, try to read the config...")
				c, err := conf.ReadConfig(cli.GetConfigFilePath())
				if err != nil {
					fmt.Println("read config failed: ", err.Error())
					return
				}

				fmt.Println("config file read successfully, try to connect database...")
				if cli.CheckDBTableExist(c.Data.Database) {
					fmt.Println("connect to database successfully and table already exists, do nothing.")
					return
				}
			}

			// start installation server to install
			install.Run(cli.GetConfigFilePath())
		},
	}

	// upgradeCmd represents the upgrade command
	upgradeCmd = &cobra.Command{
		Use:   "upgrade",
		Short: "upgrade Answer version",
		Long:  `upgrade Answer version`,
		Run: func(_ *cobra.Command, _ []string) {
			log.SetLogger(log.NewStdLogger(os.Stdout))
			cli.FormatAllPath(dataDirPath)
			cli.InstallI18nBundle(true)
			c, err := conf.ReadConfig(cli.GetConfigFilePath())
			if err != nil {
				fmt.Println("read config failed: ", err.Error())
				return
			}
			if err = migrations.Migrate(c.Data.Database, c.Data.Cache, upgradeVersion); err != nil {
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
			cli.FormatAllPath(dataDirPath)
			c, err := conf.ReadConfig(cli.GetConfigFilePath())
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
			cli.FormatAllPath(dataDirPath)
			fmt.Println("Start checking the required environment...")
			if cli.CheckConfigFile(cli.GetConfigFilePath()) {
				fmt.Println("config file exists [✔]")
			} else {
				fmt.Println("config file not exists [x]")
			}

			if cli.CheckUploadDir() {
				fmt.Println("upload directory exists [✔]")
			} else {
				fmt.Println("upload directory not exists [x]")
			}

			c, err := conf.ReadConfig(cli.GetConfigFilePath())
			if err != nil {
				fmt.Println("read config failed: ", err.Error())
				return
			}

			if cli.CheckDBConnection(c.Data.Database) {
				fmt.Println("db connection successfully [✔]")
			} else {
				fmt.Println("db connection failed [x]")
			}
			fmt.Println("check environment all done")
		},
	}

	// buildCmd used to build another answer with plugins
	buildCmd = &cobra.Command{
		Use:   "build",
		Short: "used to build answer with plugins",
		Long:  `Build a new Answer with plugins that you need`,
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("try to build a new answer with plugins:\n%s\n", strings.Join(buildWithPlugins, "\n"))
			err := cli.BuildNewAnswer(buildOutput, buildWithPlugins, cli.OriginalAnswerInfo{
				Version:  Version,
				Revision: Revision,
				Time:     Time,
			})
			if err != nil {
				fmt.Printf("build failed %v", err)
			} else {
				fmt.Printf("build new answer successfully %s\n", buildOutput)
			}
		},
	}

	// pluginCmd prints all plugins packed in the binary
	pluginCmd = &cobra.Command{
		Use:   "plugin",
		Short: "prints all plugins packed in the binary",
		Long:  `prints all plugins packed in the binary`,
		Run: func(_ *cobra.Command, _ []string) {
			_ = plugin.CallBase(func(base plugin.Base) error {
				info := base.Info()
				fmt.Printf("%s[%s] made by %s\n", info.SlugName, info.Version, info.Author)
				return nil
			})
		},
	}
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
