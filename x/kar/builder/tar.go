package builder

import (
	"archive/tar"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

type Builder func(string, io.WriteCloser, []Entry) error

type Entry struct {
	Name string
	Size int64
}

func Tar(src string, dst io.WriteCloser, entries []Entry) error {
	tw := tar.NewWriter(dst)
	defer tw.Close()

	for _, entry := range entries {
		fmt.Printf("adding %s\n", entry.Name)

		reader, err := os.Open(filepath.Join(src, entry.Name))
		if err != nil {
			return err
		}

		relPath, err := filepath.Rel(src, reader.Name())
		if err != nil {
			return err
		}

		stat, err := reader.Stat()
		if err != nil {
			return err
		}

		fmt.Printf("adding %s\n", relPath)
		hdr := tar.Header{
			Name:     entry.Name,
			Size:     stat.Size(),
			Mode:     int64(stat.Mode()),
			ModTime:  stat.ModTime(),
			Typeflag: tar.TypeReg,
		}

		fmt.Printf("adding %#v\n", hdr)

		if err := tw.WriteHeader(&hdr); err != nil {
			return err
		}

		if _, err := io.Copy(tw, reader); err != nil {
			return err
		}

		reader.Close()
	}

	return nil
}
