/*
Copyright © 2023 Morteza Khazamipour me@morteza.dev
*/
package cmd

import (
	"context"
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// defragCmd represents the defrag command
var defragCmd = &cobra.Command{
	Use:   "defrag",
	Short: "Defrag etcd endpoints provided in config or flags",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		etcdep := viper.GetStringSlice("etcd.endpoints")
		for _, ep := range etcdep {
			dfresp, err := newClient().Defragment(context.Background(), ep)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("endpoint: %s / Defragment completed: %v\n", ep, dfresp.Header.GetMemberId())
		}
	},
}

func init() {
	rootCmd.AddCommand(defragCmd)
}
