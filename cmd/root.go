package cmd

import (
	"fmt"
	"os"

	"github.com/mateo1647/kk/internal/options"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string
var namespace string
var context string
var kubeconfig string
var labels string

var rootCmd = &cobra.Command{
	Use:   "kk",
	Short: "make kubectl moar easier",
	Long:  `a CLI to make kubectl commands easier`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// generic search options handler
var searchOptions = options.NewSearchOptions()

func init() {
	// Global Flags
	rootCmd.PersistentFlags().StringVarP(
		&searchOptions.Namespace, "namespace", "n", "",
		"Namespace for search. (default: \"default\")")
	rootCmd.PersistentFlags().BoolVarP(
		&searchOptions.AllNamespaces, "all-namespaces", "A", false,
		"If present, list the requested object(s) across all namespaces.")
	rootCmd.PersistentFlags().StringVarP(
		&searchOptions.Selector, "selector", "l", "",
		"Selector (label query) to filter on. (e.g. -l key1=value1,key2=value2)")
	rootCmd.PersistentFlags().StringVar(
		&searchOptions.FieldSelector, "field-selector", "",
		"Selector (field query) to filter on. (e.g. --field-selector key1=value1,key2=value2)")
}

func initConfig() {
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

		// Search config in home directory with name ".kk" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".kk")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}

	if kubeconfig == "" {
		if kenv := os.Getenv("KUBECONFIG"); kenv != "" {
			kubeconfig = kenv
		} else if h := os.Getenv("HOME"); h != "" {
			kubeconfig = fmt.Sprintf("%v/.kube/config", h)
		} else {
			panic(fmt.Errorf("error setting default kubeconfig. $HOME not set"))
		}
	}
}
