package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"syscall"

	. "github.com/alexpantyukhin/go-pattern-match"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const COMMANDS_CONFIG = "config"
const COMMANDS_CONFIG_INIT = "init"
const COMMANDS_CONFIG_ALIASES = "aliases"
const COMMANDS_CONFIG_ALIASES_ADD = "add"
const COMMANDS_CONFIG_ALIASES_REMOVE = "remove"

var cfgFile string

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	DisableFlagParsing: true,
	Use:                "gitm",
	Short:              "Stay on the scene like a git machine. ",
	Long:               `Git on down or up in style with the smoothest git aliases and the funkiest git flow.`,
	Run: func(cmd *cobra.Command, args []string) {

		firstArg := args[0]

		if firstArg == "config" {
			// do something
			configCommand(cmd, args)
			return
		}

		argumentWords := []string{
			"all",
		}

		commandWords := []string{
			"about",
			"begin",
			"break",
			"done",
			"down",
			"feel",
			"feels",
			"good",
			"into",
			"it up",
			"it",
			"me",
			"mess",
			"minute",
			"out",
			"say",
			"scene",
			"shout",
			"started",
			"stuff",
			"the",
			"there",
			"together",
			"try",
			"up",
			"woo",
			"yeah",
			"yell",
			"yourself",
		}

		allowedWords := append(argumentWords, commandWords...)

		var keywords = make(map[string]bool)

		lastIndex := 0

		for index, arg := range args {
			for _, allowedWord := range allowedWords {
				if arg == allowedWord {
					lastIndex = index
					keywords[arg] = true
					break
				}
			}
		}

		extraArgs := args[lastIndex+1:]

		_, shouldTry := Match(keywords).When(map[string]interface{}{
			"try": true,
		}, true).Result()

		_, command := Match(keywords).
			When(map[string]interface{}{
				"about": true,
			}, "branch").

			// checkout
			When(map[string]interface{}{
				"into": true,
				"it":   true,
			}, "checkout").
			When(map[string]interface{}{
				"into": true,
			}, "checkout").

			// commit
			When(map[string]interface{}{
				"say": true,
			}, "commit").
			When(map[string]interface{}{
				"shout": true,
			}, "commit").
			When(map[string]interface{}{
				"yell": true,
			}, "commit").

			// blame
			When(map[string]interface{}{
				"yourself": true,
			}, "blame").
			When(map[string]interface{}{
				"me": true,
			}, "blame").

			// stash
			When(map[string]interface{}{
				"minute": true,
			}, "stash").
			When(map[string]interface{}{
				"scene": true,
			}, "stash").

			// init
			When(map[string]interface{}{
				"started": true,
			}, "init").
			When(map[string]interface{}{
				"begin": true,
			}, "init").

			// remote
			When(map[string]interface{}{
				"there": true,
			}, "remote").

			// rm
			When(map[string]interface{}{
				"out": true,
			}, "rm").
			When(map[string]interface{}{
				"lost": true,
			}, "rm").

			// merge
			When(map[string]interface{}{
				"together": true,
			}, "merge").

			// status
			When(map[string]interface{}{
				"woo": true,
			}, "status").
			When(map[string]interface{}{
				"yeah": true,
			}, "status").

			// diff
			When(map[string]interface{}{
				"stuff": true,
			}, "diff").
			When(map[string]interface{}{
				"mess": true,
			}, "diff").

			// rebase
			When(map[string]interface{}{
				"break": true,
				"down":  true,
			}, "rebase").

			// pull
			When(map[string]interface{}{
				"down": true,
			}, "pull").

			// push
			When(map[string]interface{}{
				"up": true,
			}, "push").

			// diff
			When(map[string]interface{}{
				"the": true,
			}, "add").
			When(map[string]interface{}{
				"that": true,
			}, "add").
			Result()

		commandString := fmt.Sprintf("%v", command)

		execArgs := append([]string{"git", commandString}, extraArgs...)

		binary, lookErr := exec.LookPath("git")
		if lookErr != nil {
			panic(lookErr)
		}

		if shouldTry == true {
			fmt.Println("---git-machine sandbox mode---")
			fmt.Println("Would have executed: ")
			fmt.Println(execArgs)
		} else {
			execError := syscall.Exec(binary, execArgs, os.Environ())
			if execError != nil {
				panic(execError)
			}
		}
	},
}

func configCommand(cmd *cobra.Command, args []string) {

	configCommand := args[1]

	if configCommand == COMMANDS_CONFIG_INIT {
		viper.SafeWriteConfig()
	}

	// add aliases
	if configCommand == COMMANDS_CONFIG_ALIASES {

		aliasCommand := args[2]

		if aliasCommand == COMMANDS_CONFIG_ALIASES_ADD {
			aliasFolderPath := viper.GetString("AliasFolderPath")

			// does alias folder exist
			_, err := ioutil.ReadDir(aliasFolderPath)
			if err != nil {
				os.Mkdir(aliasFolderPath, 0770)
			}

			// get gitm path
			gitmPath, err := exec.LookPath("gitm")
			if err != nil {
				log.Fatal("gitm not found on PATH.")
				os.Exit(1)
			}

			// create symlinks
			symlinks := viper.GetStringSlice("Aliases")

			for _, symlink := range symlinks {
				os.Symlink(gitmPath, fmt.Sprintf("%v/%v", aliasFolderPath, symlink))
				_, symlinkError := exec.LookPath(symlink)
				if symlinkError != nil {
					log.Fatal("alias not found on PATH. Make sure you have configured the correct folder.")
					os.Exit(1)
				}
			}

		}

		// remove aliases
		if aliasCommand == COMMANDS_CONFIG_ALIASES_REMOVE {
			aliasFolderPath := viper.GetString("AliasFolderPath")

			// does alias folder exist
			_, err := ioutil.ReadDir(aliasFolderPath)
			if err != nil {
				log.Fatal("git machine alias folder does not exist. Check your configured path")
				os.Exit(1)
			}

			// get gitm path
			gitmPath, err := exec.LookPath("gitm")
			if err != nil {
				log.Fatal("gitm not found on PATH.")
				os.Exit(1)
			}

			// quietly remove symlinks. use verbose flag to decide to show
			symlinks := viper.GetStringSlice("Aliases")

			for _, symlink := range symlinks {
				symlinkPath := fmt.Sprintf("%v/%v", aliasFolderPath, symlink)

				symlinkInfo, _ := os.Lstat(symlinkPath)
				if symlinkInfo.Mode()&os.ModeSymlink != 0 {
					symlinkLink, err := os.Readlink(symlinkPath)
					if err != nil || symlinkLink == gitmPath {
						os.Remove(symlinkPath)
					}
				}
			}

		}
	}
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.git-machine)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// RootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func initConfig() {
	viper.SetDefault("Aliases", []string{"just", "lets", "want", "need"})

	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		viper.SetDefault("AliasFolderPath", fmt.Sprintf("%v/gitm-aliases", home))

		viper.AddConfigPath(home)
		viper.SetConfigName(".git-machine")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
