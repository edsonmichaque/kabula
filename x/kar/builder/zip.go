package builder

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func Zip(target string, dst io.WriteCloser, entries []Entry) error {
	writer := zip.NewWriter(dst)
	defer writer.Close()
	for _, e := range entries {
		fmt.Printf("adding %s\n", e.Name)

		c, err := os.Open(filepath.Join(target, e.Name))
		if err != nil {
			return err
		}

		stat, err := c.Stat()
		if err != nil {
			return err
		}

		header, err := zip.FileInfoHeader(stat)
		if err != nil {
			return err
		}

		hdr, err := writer.CreateHeader(header)
		if err != nil {
			return err
		}

		if _, err := io.Copy(hdr, c); err != nil {
			return err
		}

		c.Close()
	}

	return nil
}
