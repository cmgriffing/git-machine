package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
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
	Use:   "gitm",
	Short: "Stay on the scene like a git machine. ",
	Long:  `Git on down or up in style with the smoothest git aliases and the funkiest git flow.`,
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
			"down",
			"up",
			"it up",
			"into",
			"together",
			"the",
			"feel",
			"feels",
			"good",
			"woo",
			"yeah",
			"mess",
			"stuff",
			"out",
			"started",
			"done",
			"break",
			"there",
			"say",
			"yell",
			"shout",
			"yourself",
			"me",
			"minute",
			"scene",
			"begin",
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

		_, command := Match(keywords).

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

			// checkout
			When(map[string]interface{}{
				"into": true,
			}, "checkout").

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

		extras := ""

		if len(extraArgs) > 0 {
			extras = fmt.Sprintf(" %v", strings.Join(extraArgs, " "))
		}

		specifics := fmt.Sprintf("git %v%v", command, extras)

		fmt.Println("EXECUTING:")
		fmt.Println(specifics)

		binary, lookErr := exec.LookPath("git")
		if lookErr != nil {
			panic(lookErr)
		}

		execError := syscall.Exec(binary, strings.Split(specifics, " "), os.Environ())
		if execError != nil {
			panic(execError)
		}
	},
}

func configCommand(cmd *cobra.Command, args []string) {

	configCommand := args[1]

	if configCommand == COMMANDS_CONFIG_INIT {
		fmt.Println("Saving file")
		viper.SafeWriteConfig()
	}

	// add aliases
	if configCommand == COMMANDS_CONFIG_ALIASES {

		aliasCommand := args[2]

		if aliasCommand == COMMANDS_CONFIG_ALIASES_ADD {
			fmt.Println("Adding aliased gitm commands")

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
				fmt.Println(gitmPath)
				_, symlinkError := exec.LookPath(symlink)
				if symlinkError != nil {
					log.Fatal("alias not found on PATH. Make sure you have configured the correct folder.")
					os.Exit(1)
				}
			}

		}

		// remove aliases
		if aliasCommand == COMMANDS_CONFIG_ALIASES_REMOVE {
			fmt.Println("Removing aliased gitm commands")

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
