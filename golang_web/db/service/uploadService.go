package service

import (
	"io"
	"os"
	"path"
)

type UploadServicer interface {
	UploadImage(string, io.Reader) (string, error)
}

type LocalUploadService struct {
	filePath    string // ./images
	netBasePath string // http://localhost:8080/images
}

func NewLocalUploadService(filePath, netBasePath string) *LocalUploadService {
	return &LocalUploadService{
		filePath:    filePath,
		netBasePath: netBasePath,
	}
}

func (l *LocalUploadService) UploadImage(objName string, reader io.Reader) (string, error) {
	dst := path.Join(l.filePath, objName)
	out, err := os.Create(dst)
	if err != nil {
		return "", err
	}
	defer out.Close()

	_, err = io.Copy(out, reader)

	return path.Join(l.netBasePath, objName), err
}
