package aferos3backend

import "os"

type S3File struct {
}

func (S3File) Close() error {
	panic("not implemented")
	return nil
}

func (S3File) Read(p []byte) (n int, err error) {
	panic("not implemented")
	return 0, nil
}

func (S3File) ReadAt(p []byte, off int64) (n int, err error) {
	panic("not implemented")
	return 0, nil
}

func (S3File) Seek(offset int64, whence int) (int64, error) {
	panic("not implemented")
	return 0, nil
}

func (S3File) Write(p []byte) (n int, err error) {
	panic("not implemented")
	return 0, nil
}

func (S3File) WriteAt(p []byte, off int64) (n int, err error) {
	panic("not implemented")
	return 0, nil
}

func (S3File) Name() string {
	panic("not implemented")
	return ""
}

func (S3File) Readdir(count int) ([]os.FileInfo, error) {
	panic("not implemented")
	return nil, nil
}

func (S3File) Readdirnames(n int) ([]string, error) {
	panic("not implemented")
	return nil, nil
}

func (S3File) Stat() (os.FileInfo, error) {
	panic("not implemented")
	return nil, nil
}

func (S3File) Sync() error {
	panic("not implemented")
	return nil
}

func (S3File) Truncate(size int64) error {
	panic("not implemented")
	return nil
}

func (S3File) WriteString(s string) (ret int, err error) {
	panic("not implemented")
	return 0, nil
}
