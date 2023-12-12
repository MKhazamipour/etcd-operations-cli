/*
Copyright © 2023 Morteza Khazamipour me@morteza.dev
*/
package cmd

import (
	"context"
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// local represents the s3 command
var localBackup = &cobra.Command{
	Use:   "local",
	Short: "Backup etcd to localdisk",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		now := time.Now().Format(time.DateTime)
		formattedTime := strings.NewReplacer(" ", "_", "-", "_", ":", "_").Replace(now)

		client := newClient()
		snapshotFilePath := viper.GetString("backupLocation")
		ctx := context.TODO()

		resp, err := client.Snapshot(ctx)
		if err != nil {
			fmt.Println(err)
		}
		defer resp.Close()

		snapshot, err := io.ReadAll(resp)
		if err != nil {
			fmt.Println(err)
		}

		snapshotFile, err := os.Create(snapshotFilePath + "_etcd-backup_" + formattedTime)
		if err != nil {
			fmt.Println(err)
		}
		defer snapshotFile.Close()

		_, err = snapshotFile.Write(snapshot)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("Snapshot saved to: %s\n", snapshotFile.Name())
	},
}

func init() {
	backupCmd.AddCommand(localBackup)
	localBackup.PersistentFlags().StringP("backup-location", "l", "", "Location to save the etcd backup on disk, /tmp/backup1.db")
	viper.BindPFlag("backupLocation", localBackup.PersistentFlags().Lookup("backup-location"))
}
