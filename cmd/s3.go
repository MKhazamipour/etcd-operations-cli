/*
Copyright © 2023 Morteza Khazamipour me@morteza.dev
*/
package cmd

import (
	"context"
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var s3Cmd = &cobra.Command{
	Use:   "s3",
	Short: "Backup etcd to s3 compatible storage",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		now := time.Now().Format(time.DateTime)
		formattedTime := strings.NewReplacer(" ", "_", "-", "_", ":", "_").Replace(now)
		client := newClient()
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

		bucket := aws.String(viper.GetString("s3BucketName"))
		key := aws.String("etcd_backup_" + formattedTime)

		s3Config := &aws.Config{
			Credentials:      credentials.NewStaticCredentials(viper.GetString("s3AccessKey"), viper.GetString("s3SecretKey"), ""),
			Endpoint:         aws.String(viper.GetString("s3Endpoint")),
			Region:           aws.String(viper.GetString("s3Region")),
			DisableSSL:       aws.Bool(false),
			S3ForcePathStyle: aws.Bool(true),
		}
		newSession := session.New(s3Config)

		s3Client := s3.New(newSession)
		_, err = s3Client.PutObject(&s3.PutObjectInput{
			Body:   strings.NewReader(string(snapshot)),
			Bucket: bucket,
			Key:    key,
		})
		if err != nil {
			fmt.Printf("Failed to upload data to %s/%s, %s\n", *bucket, *key, err.Error())
			return
		}
		fmt.Printf("Successfully uploaded backup with key %s\n", *key)
	},
}

func init() {
	backupCmd.AddCommand(s3Cmd)
	s3Cmd.PersistentFlags().StringP("bucket-name", "b", "", "Bucket name to save etcd's snapshot")
	viper.BindPFlag("s3BucketName", s3Cmd.PersistentFlags().Lookup("bucket-name"))
	s3Cmd.PersistentFlags().StringP("s3-endpoint", "p", "", "S3 Region, can be MinIO endpoint")
	viper.BindPFlag("s3Endpoint", s3Cmd.PersistentFlags().Lookup("s3-endpoint"))
	s3Cmd.PersistentFlags().StringP("region", "r", "", "S3 Region")
	viper.BindPFlag("s3Region", s3Cmd.PersistentFlags().Lookup("region"))
	s3Cmd.PersistentFlags().StringP("s3-access-key", "n", "", "S3 Access Key")
	viper.BindPFlag("s3AccessKey", s3Cmd.PersistentFlags().Lookup("s3-access-key"))
	s3Cmd.PersistentFlags().StringP("s3-secret-key", "s", "", "S3 Secret Key")
	viper.BindPFlag("s3SecretKey", s3Cmd.PersistentFlags().Lookup("s3-secret-key"))
}
