package s3

import (
	"bytes"
	"context"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/take0fit/knowledge-out/internal/domain/repository"
	"io/ioutil"
)

type FileStorageRepositoryImpl struct {
	Client *s3.Client
}

func NewS3StorageService(client *s3.Client) repository.FileStorageRepository {
	return &FileStorageRepositoryImpl{Client: client}
}

func (repo *FileStorageRepositoryImpl) UploadFile(ctx context.Context, bucket, key string, data []byte) error {
	_, err := repo.Client.PutObject(ctx, &s3.PutObjectInput{
		Bucket: &bucket,
		Key:    &key,
		Body:   ioutil.NopCloser(bytes.NewReader(data)),
	})
	return err
}

func (repo *FileStorageRepositoryImpl) DownloadFile(ctx context.Context, bucket, key string) ([]byte, error) {
	output, err := repo.Client.GetObject(ctx, &s3.GetObjectInput{
		Bucket: &bucket,
		Key:    &key,
	})
	if err != nil {
		return nil, err
	}
	defer output.Body.Close()
	return ioutil.ReadAll(output.Body)
}
