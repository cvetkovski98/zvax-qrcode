package repository

import (
	"context"
	"log"
	"net/url"
	"strings"
	"time"

	"cloud.google.com/go/storage"
	qrcode "github.com/cvetkovski98/zvax-qrcode/internal"
)

var projectID = "zvax-project"

type gcpObjectStore struct {
	client *storage.Client
	bucket string
}

// CreateBucket implements qrcode.ObjectStore
func (gcp *gcpObjectStore) CreateBucket(ctx context.Context) error {
	if _, err := gcp.client.Bucket(gcp.bucket).Attrs(ctx); err == nil {
		log.Printf("we already own bucket %s", gcp.bucket)
		return nil
	}
	bucketAttrs := &storage.BucketAttrs{
		Location: "europe-west6",
		UniformBucketLevelAccess: storage.UniformBucketLevelAccess{
			Enabled:    false,
			LockedTime: time.Now().Add(90 * 25 * time.Hour), // 90 days
		},
	}
	if err := gcp.client.Bucket(gcp.bucket).Create(ctx, projectID, bucketAttrs); err != nil {
		log.Printf("failed to create bucket %s", gcp.bucket)
		return err
	}
	log.Printf("successfully created bucket %s", gcp.bucket)
	return nil
}

// GetResourceLocation implements qrcode.ObjectStore
func (gcp *gcpObjectStore) GetResourceLocation(ctx context.Context, objectName string) (*url.URL, error) {
	object := gcp.client.Bucket(gcp.bucket).Object(objectName)
	attrs, err := object.Attrs(ctx)
	if err != nil {
		return nil, err
	}
	mediaLink := strings.Replace(attrs.MediaLink, "/download", "", 1)
	return url.Parse(mediaLink)
}

// RemoveQR implements qrcode.ObjectStore
func (gcp *gcpObjectStore) RemoveQR(ctx context.Context, email string) error {
	objectName := getObjectName(email)
	return gcp.client.Bucket(gcp.bucket).Object(objectName).Delete(ctx)
}

// UploadQR implements qrcode.ObjectStore
func (gcp *gcpObjectStore) UploadQR(ctx context.Context, email string, content []byte) (string, error) {
	objectName := getObjectName(email)
	object := gcp.client.Bucket(gcp.bucket).Object(objectName)
	writer := object.NewWriter(ctx)
	if _, err := writer.Write(content); err != nil {
		return "", err
	}
	if err := writer.Close(); err != nil {
		return "", err
	}

	if err := object.ACL().Set(ctx, storage.AllUsers, storage.RoleReader); err != nil {
		return "", err
	}
	return objectName, nil
}

func NewGCPObjectStore(client *storage.Client, bucket string) qrcode.ObjectStore {
	return &gcpObjectStore{
		client: client,
		bucket: bucket,
	}
}
