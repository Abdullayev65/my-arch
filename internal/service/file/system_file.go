package file

import (
	"errors"
	"io"
	"mime/multipart"
	"my-arch/internal/model"
	"my-arch/internal/tools/stream"
	"os"
	"path"
	"strings"
	"time"
)

type SystemFile struct{}

func NewSystemFile() *SystemFile {
	return &SystemFile{}
}

func (s *SystemFile) Upload(file *multipart.FileHeader, folder string) (string, error) {
	if file == nil {
		return "", errors.New("Files.Upload: Files is null")
	}

	dst := path.Join("./files/", folder, s.makeSubPath(file.Filename))

	if _, err := os.Stat(s.getFolder(dst)); errors.Is(err, os.ErrNotExist) {
		err = os.MkdirAll(s.getFolder(dst), os.ModePerm)
		if err != nil {
			return "", err
		}
	}

	src, err := file.Open()
	if err != nil {
		return "", err
	}

	defer src.Close()

	out, err := os.Create(dst)
	if err != nil {
		return "", err
	}

	defer out.Close()

	_, err = io.Copy(out, src)

	if err != nil {
		return "err", err
	}

	return dst, nil
}

func (s *SystemFile) Delete(dst string) error {
	return os.Remove("." + dst)
}

func (s *SystemFile) MultipleUpload(files []*multipart.FileHeader, folder string) ([]string, error) {
	var links []string

	for _, file := range files {
		link, err := s.Upload(file, folder)
		if err != nil {
			return nil, err
		}

		links = append(links, link)
	}

	return links, nil
}

func (s *SystemFile) UploadFile(file *multipart.FileHeader, folder string) (*model.FileData, error) {
	path, err := s.Upload(file, folder)
	if err != nil {
		return nil, err
	}

	fd := new(model.FileData)
	fd.Path = path
	fd.Name = file.Filename
	fd.Size = file.Size

	return fd, nil
}

func (s *SystemFile) MultipleUploadFile(files []*multipart.FileHeader, folder string) (fds []*model.FileData, err error) {
	fds = make([]*model.FileData, 0, len(files))

	defer func() {
		if err == nil {
			return
		}
		stream.ForEach(fds, func(f *model.FileData) {
			s.Delete(f.Path)
		})
	}()

	for _, file := range files {
		fd, er1 := s.UploadFile(file, folder)
		if er1 != nil {
			err = er1
			return nil, err
		}

		fds = append(fds, fd)
	}

	return
}

func (s *SystemFile) makeSubPath(filename string) string {
	now := time.Now().Format("06/01/02/15.04.05.")
	return now + filename
}

func (s *SystemFile) getFolder(dst string) string {
	idx := strings.LastIndex(dst, "/")
	if idx == -1 {
		return "./"
	}
	return dst[:idx]
}
