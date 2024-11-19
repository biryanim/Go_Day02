package rotate

import (
	"archive/tar"
	"compress/gzip"
	"io"
	"os"
)

func CreateArchive(filename string, buf io.Writer) error {
	gw := gzip.NewWriter(buf)
	defer gw.Close()
	tw := tar.NewWriter(gw)
	defer tw.Close()
	err := AddToArchive(tw, filename)
	if err != nil {
		return err
	}
	return nil
}

func AddToArchive(tw *tar.Writer, filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	info, err := file.Stat()
	if err != nil {
		return err
	}

	header, err := tar.FileInfoHeader(info, info.Name())
	if err != nil {
		return err
	}
	header.Name = filename

	err = tw.WriteHeader(header)
	if err != nil {
		return err
	}
	_, err = io.Copy(tw, file)
	if err != nil {
		return err
	}
	return nil
}
