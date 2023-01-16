package cmd

import (
	"github.com/cvetkovski98/zvax-qrcode/internal/config"
	"github.com/cvetkovski98/zvax-qrcode/internal/repository"
	"github.com/spf13/cobra"
)

var createBucketCommand = &cobra.Command{
	Use:   "create-bucket",
	Short: "Create MinIO bucket",
	Long:  `Create MinIO bucket`,
	RunE:  createBucket,
}

func init() {
	createBucketCommand.Flags().StringP("bucket", "b", "", "Bucket name")
}

func createBucket(cmd *cobra.Command, args []string) error {
	ctx := cmd.Context()
	cfg := config.GetConfig()
	store, err := repository.Create(ctx, cfg)
	if err != nil {
		return err
	}
	bucket := cfg.ObjectStore.BucketName
	if cmd.Flag("bucket").Value.String() != "" {
		bucket = cmd.Flag("bucket").Value.String()
	}
	if err := store.CreateBucket(ctx, bucket); err != nil {
		return err
	}
	return nil
}
