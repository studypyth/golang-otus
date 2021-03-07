package main

import (
	"errors"
	"io"
	"os"
)

var (
	ErrUnsupportedFile       = errors.New("unsupported file")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
)

func Copy(from, to string, offset, limit int64) error {
	fi, err := os.Stat(from)
	if err != nil {
		return ErrUnsupportedFile
	}
	size := fi.Size()
	// offset больше, чем размер файла - невалидная ситуация;
	if offset > size {
		return ErrOffsetExceedsFileSize
	}
	srcFile, err := os.OpenFile(from, os.O_RDONLY, 0)
	if err != nil {
		return err
	}
	if offset > 0 {
		_, err := srcFile.Seek(offset, io.SeekCurrent)
		if err != nil {
			return err
		}
	}
	dstFile, err := os.Create(to)
	if err != nil {
		return err
	}
	if limit == 0 || limit > size {
		limit = size
	}
	_, err = io.CopyN(dstFile, srcFile, limit)
	if err != nil && !errors.Is(err, io.EOF) {
		return err
	}
	defer func() {
		err1 := dstFile.Close()
		err2 := srcFile.Close()
		if err1 != nil && err2 != nil {
			err = err1
		}
	}()
	return nil
}
