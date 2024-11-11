package common

import (
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"
)

type ImageUpload interface {
	UploadImage(file *multipart.FileHeader, path string) (string, error)
}

type imageUpload struct{

}

func NewImageUpload () ImageUpload{
	return &imageUpload{}
}

func (ic *imageUpload) UploadImage(file *multipart.FileHeader, path string) (string, error){

	// source 
	src, err := file.Open()
	if err!=nil{
		return "", err
	}
	defer src.Close()

	timeNow := time.Now().Format("2006-0102T15_0405_070000")
	fileName := timeNow + "_" + file.Filename
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