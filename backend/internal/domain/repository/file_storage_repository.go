package repository

import "context"

type FileStorageRepository interface {
	UploadFile(ctx context.Context, bucket, key string, data []byte) error
	DownloadFile(ctx context.Context, bucket, key string) ([]byte, error)
}
