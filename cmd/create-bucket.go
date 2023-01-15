package cmd

import (
	"github.com/cvetkovski98/zvax-qrcode/internal/config"
	"github.com/cvetkovski98/zvax-qrcode/internal/repository"
	"github.com/cvetkovski98/zvax-qrcode/pkg/minio"
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
	cfg := config.GetConfig()
	minioClient, err := minio.NewMinioClient(&cfg.MinIO)
	if err != nil {
		return err
	}
	store := repository.NewMinioObjectStore(minioClient)
	bucket := cfg.MinIO.BucketName
	if cmd.Flag("bucket").Value.String() != "" {
		bucket = cmd.Flag("bucket").Value.String()
	}
	if err = store.CreateBucket(cmd.Context(), bucket); err != nil {
		return err
	}
	return nil
}
