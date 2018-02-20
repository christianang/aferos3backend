package aferos3backend_test

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"

	. "github.com/onsi/gomega"
)

func s3Client() *s3.S3 {
	return s3.New(session.New(&aws.Config{
		Credentials: credentials.NewStaticCredentials(config.AccessKeyID, config.SecretAccessKey, ""),
		Region:      aws.String(config.Region),
	}))
}

func cleanupBucket() {
	objectsList, err := s3Client().ListObjects(&s3.ListObjectsInput{
		Bucket: aws.String(config.Bucket),
	})
	Expect(err).NotTo(HaveOccurred())

	objectsToDelete := []*s3.ObjectIdentifier{}
	for _, object := range objectsList.Contents {
		objectsToDelete = append(objectsToDelete, &s3.ObjectIdentifier{
			Key: object.Key,
		})
	}

	result, err := s3Client().DeleteObjects(&s3.DeleteObjectsInput{
		Bucket: aws.String(config.Bucket),
		Delete: &s3.Delete{
			Objects: objectsToDelete,
			Quiet:   aws.Bool(false),
		},
	})
	Expect(err).NotTo(HaveOccurred())
	Expect(result.Errors).To(BeEmpty())
}

func fileExistsOnS3(filename string) (bool, error) {
	result, err := s3Client().ListObjects(&s3.ListObjectsInput{
		Bucket: aws.String(config.Bucket),
		Prefix: aws.String(filename),
	})
	Expect(err).NotTo(HaveOccurred())

	for _, content := range result.Contents {
		return aws.StringValue(content.Key) == filename, nil
	}

	return false, nil
}
