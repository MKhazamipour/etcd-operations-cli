/*
Copyright © 2023 Morteza Khazamipour
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "etcd-operations-cli",
	Short: "Backup and defragmenting tool for etcd",
	Long:  ``,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/etcd-configs.yaml)")
	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.PersistentFlags().StringSliceP("etcd-endpoints", "e", []string{}, "etcd endpoints separated by comma, 127.0.0.1:2379,127.0.0.2:2379")
	viper.BindPFlag("etcd.endpoints", rootCmd.PersistentFlags().Lookup("etcd-endpoints"))
	rootCmd.PersistentFlags().StringP("ca-cert", "a", "", "etcd ca-cert file")
	viper.BindPFlag("caPath.cacert", rootCmd.PersistentFlags().Lookup("ca-cert"))
	rootCmd.PersistentFlags().StringP("key", "k", "", "etcd key file")
	viper.BindPFlag("caPath.key", rootCmd.PersistentFlags().Lookup("key"))
	rootCmd.PersistentFlags().StringP("cert", "c", "", "etcd cert file")
	viper.BindPFlag("caPath.cert", rootCmd.PersistentFlags().Lookup("cert"))
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".etcd-operations-cli" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName("etcd-configs")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
