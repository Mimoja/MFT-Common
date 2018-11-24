package MFTCommon

import (
	"bytes"
	"github.com/minio/minio-go"
	"github.com/sirupsen/logrus"
)

const s3Endpoint = "localhost:9000"
const s3AccessKeyID = "***REMOVED***"
const s3SecretAccessKey = "***REMOVED***"
const s3useSSL = false

const DownloadedBucket = "downloader"
const FlashImagesBucket = "flashimages"
const MEImagesBucket = "meimages"
const MicrocodeBucket = "microcode"
const EFIBucket = "efiimages"
const CryptoBucket = "crypto"

type Storage struct {
	client *minio.Client
	log    *logrus.Logger
}

func StorageConnect(Log *logrus.Logger) *Storage {
	buckets := []string{DownloadedBucket, FlashImagesBucket, MEImages, EFIBucket, MicrocodeBucket, CryptoBucket}

	// Initialize minio client object.
	Log.Info("Connecting to S3")
	client, err := minio.New(s3Endpoint, s3AccessKeyID, s3SecretAccessKey, s3useSSL)
	if err != nil {
		Log.WithError(err).Panic("Connecting failed")
		return nil // unreached
	}
	storage := &Storage{client: client, log: Log}

	// Create starting buckets
	for _, bucket := range buckets {
		Log.Info("Creating Bucket ", bucket)
		err = storage.MakeBucket(bucket)
		if err != nil {
			Log.WithError(err).Panic("Bucket creation failed")
			return nil // unreached
		}
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

	logrus.Debug("Successfully created", bucketName)
	return nil
}

func (storage Storage) StoreBytes(bucketName string, byte []byte, remotePath string) error {
	_, err := storage.client.PutObject(bucketName, remotePath, bytes.NewReader(byte), int64(len(byte)), minio.PutObjectOptions{})
	if err != nil {
		storage.log.WithError(err).Error("Could not store: %v\n", err)
	}
	return err
}

func (storage Storage) StoreFile(bucketName string, localPath string, remotePath string) error {

	_, err := storage.client.FPutObject(bucketName, remotePath, localPath, minio.PutObjectOptions{})
	return err
}

func (storage Storage) FileExists(bucketName string, remotePath string) (minio.ObjectInfo, error) {
	return storage.client.StatObject(bucketName, remotePath, minio.StatObjectOptions{})
}

func (storage Storage) GetFile(bucketName string, remotePath string) (*minio.Object, error) {

	return storage.client.GetObject(bucketName, remotePath, minio.GetObjectOptions{})
}
