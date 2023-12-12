/*
Copyright © 2023 Morteza Khazamipour me@morteza.dev
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// backupCmd represents the backup command
var backupCmd = &cobra.Command{
	Use:   "backup",
	Short: "backup etcd cluster to localdisk or s3 compatible storage",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	rootCmd.AddCommand(backupCmd)
}
