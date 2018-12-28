package MFTCommon

import (
	"bytes"
	"github.com/minio/minio-go"
	"github.com/sirupsen/logrus"
)

const s3Endpoint = "***REMOVED***:9000"
const s3AccessKeyID = "***REMOVED***"
const s3SecretAccessKey = "***REMOVED***"
const s3useSSL = false

const MainBucket = "files"

type Storage struct {
	client *minio.Client
	log    *logrus.Logger
}

func StorageConnect(Log *logrus.Logger) *Storage {

	// Initialize minio client object.
	Log.Info("Connecting to S3")
	client, err := minio.New(s3Endpoint, s3AccessKeyID, s3SecretAccessKey, s3useSSL)
	if err != nil {
		Log.WithError(err).Panic("Connecting failed")
		return nil // unreached
	}
	storage := &Storage{client: client, log: Log}

	// Create bucket
	Log.Info("Creating Bucket ", MainBucket)
	err = storage.MakeBucket(MainBucket)
	if err != nil {
		Log.WithError(err).Panic("Bucket creation failed")
		return nil // unreached
	}

	return storage
}

func (storage Storage) GetClient() *minio.Client {
	return storage.client
}

func (storage Storage) MakeBucket(bucketName string) error {
	exists, err := storage.client.BucketExists(bucketName)
	if err != nil {
		return err
	}
	if exists {
		storage.log.Debug("Bucket '", bucketName, "' already exists ")
		return nil
	}

	err = storage.client.MakeBucket(bucketName, "us-east-1")
	if err != nil {
		storage.log.WithError(err).Error("Could not create Minio Bucket")
		return err
	}

	storage.log.Debug("Successfully created Bucket")
	return nil
}

func (storage Storage) StoreBytes(byte []byte, remotePath string) error {
	_, err := storage.client.PutObject(MainBucket, remotePath, bytes.NewReader(byte), int64(len(byte)), minio.PutObjectOptions{})
	if err != nil {
		storage.log.WithError(err).Error("Could not store: %v\n", err)
	}
	return err
}

func (storage Storage) StoreFile(localPath string, remotePath string) error {

	_, err := storage.client.FPutObject(MainBucket, remotePath, localPath, minio.PutObjectOptions{})
	return err
}

func (storage Storage) FileExists(remotePath string) (minio.ObjectInfo, error) {
	return storage.client.StatObject(MainBucket, remotePath, minio.StatObjectOptions{})
}

func (storage Storage) GetFile(remotePath string) (*minio.Object, error) {

	return storage.client.GetObject(MainBucket, remotePath, minio.GetObjectOptions{})
}
