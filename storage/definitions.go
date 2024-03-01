package storage

import "mime/multipart"

type FileDetails struct {
	Disk      string
	Extension string
	FileName  string
	Size      int64
	Path      string
	URL       string
	MimeType  string
}

type S3Config struct {
	Bucket          string
	AccessKeyId     string
	AccessKeySecret string
	Region          string
	AccountId       *string
	CDNUrl          *string
}

type Storage interface {
	Delete(files ...string) error
	Exist(path string) bool
	GetBytes(file string) ([]byte, error)
	// MakeDirectory(directory string) error
	MimeType(file string) (string, error)
	Put(path string, source *multipart.FileHeader) (FileDetails, error)
	PutFile(path string, source *multipart.FileHeader) (FileDetails, error)
	Size(file string) (int64, error)
}
