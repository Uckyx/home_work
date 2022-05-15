package main

import (
	"errors"
	"io"
	"os"

	"github.com/cheggaaa/pb/v3"
)

var (
	ErrUnsupportedFile       = errors.New("unsupported file")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
	ErrFileNotExist          = errors.New("file not exist")
)

func Copy(fromPath, toPath string, offset, limit int64) error {
	fromFileStat, err := os.Stat(fromPath)
	if err != nil {
		return ErrFileNotExist
	}

	if !fromFileStat.Mode().IsRegular() {
		return ErrUnsupportedFile
	}

	if offset > fromFileStat.Size() {
		return ErrOffsetExceedsFileSize
	}

	fromFile, err := os.Open(fromPath)
	if err != nil {
		return err
	}
	defer fromFile.Close()

	toFile, err := os.Create(toPath)
	if err != nil {
		return err
	}
	defer toFile.Close()

	_, err = fromFile.Seek(offset, 0)
	if err != nil {
		return err
	}

	if limit == 0 {
		limit = fromFileStat.Size()
	}

	bar := pb.Full.Start64(limit)
	if limit == 0 {
		_, err = io.Copy(toFile, bar.NewProxyReader(fromFile))
		if err != nil {
			return err
		}
	} else {
		_, err = io.CopyN(toFile, bar.NewProxyReader(fromFile), limit)
		if err != nil {
			return err
		}
	}

	bar.Finish()

	return nil
}
