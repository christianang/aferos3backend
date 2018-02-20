package aferos3backend_test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestS3Fs(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "S3Fs")
}

var (
	config Config
)

type Config struct {
	AccessKeyID     string
	SecretAccessKey string
	Region          string
	Bucket          string
}

var _ = BeforeSuite(func() {
	var err error
	config, err = readConfig()
	Expect(err).NotTo(HaveOccurred())
})

func readConfig() (Config, error) {
	configPath := os.Getenv("CONFIG")

	configContents, err := ioutil.ReadFile(configPath)
	if err != nil {
		return Config{}, fmt.Errorf("failed to read config file: %s", err)
	}

	var config Config
	err = json.Unmarshal(configContents, &config)
	if err != nil {
		return Config{}, fmt.Errorf("failed to unmarshal config file: %s", err)
	}

	return config, err
}
