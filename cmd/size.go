/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// sizeCmd represents the size command
var sizeCmd = &cobra.Command{
	Use:   "size",
	Short: "Get DB size on each endpoint",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		etcdep := viper.GetStringSlice("etcd.endpoints")
		c := newClient()
		for _, ep := range etcdep {
			resp, _ := c.Status(context.TODO(), ep)
			fmt.Printf("endpoint: %s size is: %v MB\n", ep, resp.DbSize/1000000)
		}
	},
}

func init() {
	rootCmd.AddCommand(sizeCmd)
}
