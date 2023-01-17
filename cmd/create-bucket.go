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
	if cmd.Flag("bucket").Value.String() != "" {
		cfg.ObjectStore.BucketName = cmd.Flag("bucket").Value.String()
	}
	store, err := repository.Create(ctx, cfg)
	if err != nil {
		return err
	}
	if err := store.CreateBucket(ctx); err != nil {
		return err
	}
	return nil
}
