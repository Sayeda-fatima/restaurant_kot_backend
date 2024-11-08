package common

import (
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
)

type ImageUpload interface {
	UploadImage(file *multipart.FileHeader, path string) (string, error)
}

type imageUpload struct{

}

func NewImageCommon () ImageUpload{
	return &imageUpload{}
}

func (ic *imageUpload) UploadImage(file *multipart.FileHeader, path string) (string, error){

	// source 
	src, err := file.Open()
	if err!=nil{
		return "", err
	}
	defer src.Close()

	fileName := file.Filename
	fullPath := filepath.Join(path, fileName)

	// destination file
	dst, err := os.Create(fullPath)
	if err != nil{
		return "", err
	}
	defer dst.Close()

	// copy file to destination
	if _, err := io.Copy(dst, src); err!=nil{
		return "", err
	}
	return fullPath, nil
}