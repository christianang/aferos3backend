package aferos3backend

import (
	"os"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type S3Fs struct {
	s3Client *s3.S3
	bucket   string
}

func NewS3Fs(accessKeyID string, secretAccessKey string, region string, bucket string) Fs {
	s3Client := s3.New(session.New(&aws.Config{
		Credentials: credentials.NewStaticCredentials(accessKeyID, secretAccessKey, ""),
		Region:      aws.String(region),
	}))

	return &S3Fs{
		s3Client: s3Client,
		bucket:   bucket,
	}
}

func (s S3Fs) Create(name string) (File, error) {
	_, err := s.s3Client.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(name),
		Body:   aws.ReadSeekCloser(strings.NewReader("")),
	})
	if err != nil {
		panic(err)
	}

	return &S3File{}, nil
}

func (s S3Fs) Mkdir(name string, perm os.FileMode) error {
	panic("not implemented")
	return nil
}

func (s S3Fs) MkdirAll(path string, perm os.FileMode) error {
	panic("not implemented")
	return nil
}

func (s S3Fs) Open(name string) (File, error) {
	panic("not implemented")
	return nil, nil
}

func (s S3Fs) OpenFile(name string, flag int, perm os.FileMode) (File, error) {
	panic("not implemented")
	return nil, nil
}

func (s S3Fs) Remove(name string) error {
	panic("not implemented")
	return nil
}

func (s S3Fs) RemoveAll(path string) error {
	panic("not implemented")
	return nil
}

func (s S3Fs) Rename(oldname, newname string) error {
	panic("not implemented")
	return nil
}

func (s S3Fs) Stat(name string) (os.FileInfo, error) {
	panic("not implemented")
	return nil, nil
}

func (s S3Fs) Name() string {
	return "S3Fs"
}

func (s S3Fs) Chmod(name string, mode os.FileMode) error {
	panic("not implemented")
	return nil
}

func (s S3Fs) Chtimes(name string, atime time.Time, mtime time.Time) error {
	panic("not implemented")
	return nil
}
