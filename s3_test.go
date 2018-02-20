package aferos3backend_test

import (
	"github.com/christianang/aferos3backend"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Afero S3Fs Backend", func() {
	var (
		s3Fs aferos3backend.Fs
	)

	BeforeEach(func() {
		s3Fs = aferos3backend.NewS3Fs(
			config.AccessKeyID,
			config.SecretAccessKey,
			config.Region,
			config.Bucket,
		)
	})

	AfterEach(func() {
		cleanupBucket()
	})

	Describe("Name", func() {
		It("returns S3Fs", func() {
			Expect(s3Fs.Name()).To(Equal("S3Fs"))
		})
	})

	Describe("Create", func() {
		FIt("creates a file on S3", func() {
			_, err := s3Fs.Create("some-file")
			Expect(err).NotTo(HaveOccurred())

			Expect(fileExistsOnS3("some-file")).To(BeTrue())
		})
	})
})
