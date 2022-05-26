package archive

import (
	"archive/tar"
	"archive/zip"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"

	spec "github.com/edsonmichaque/kabula/kabula-spec"
)

type ArchiveArgs struct {
	Content   ContentFormat
	Container ArchiveFormat
}

var DefaultArgs = ArchiveArgs{
	Content:   FormatJSON,
	Container: KindGZip,
}

func Build(target string, args ArchiveArgs) error {
	ext := "json"
	if args.Content == FormatXML {
		ext = "xml"
	}

	fullpath := filepath.Join(target, fmt.Sprintf("manifest.%s", ext))
	f, err := os.Open(fullpath)
	if err != nil {
		return err
	}

	defer func() {
		_ = f.Close()
	}()

	data, err := io.ReadAll(f)
	if err != nil {
		return err
	}

	var m spec.Manifest
	if args.Content == FormatXML {
		if err := xml.Unmarshal(data, &m); err != nil {
			return err
		}
	} else {
		if err := json.Unmarshal(data, &m); err != nil {
			return err
		}
	}

	dst, err := os.Create(fmt.Sprintf("%s_%s.%s", m.Name, m.Version, spec.PkgExt+string(ext[0])))
	if err != nil {
		return err
	}

	defer dst.Close()

	dir, err := os.Open(target)
	if err != nil {
		return err
	}

	dirEntries, err := dir.Readdir(0)
	if err != nil {
		return err
	}

	if args.Container == KindTar {
		if err := BuildTar(target, dst, dirEntries); err != nil {
			return err
		}
	}

	if args.Container == KindZip {
		if err := BuildZip(target, dst, dirEntries); err != nil {
			return err
		}
	}

	return nil
}

func BuildTar(target string, dst io.WriteCloser, l []fs.FileInfo) error {
	writer := tar.NewWriter(dst)
	defer writer.Close()
	for _, e := range l {
		if e.IsDir() {
			continue
		}

		c, err := os.Open(filepath.Join(target, e.Name()))
		if err != nil {
			return err
		}

		entry, err := filepath.Rel(c.Name(), target)
		if err != nil {
			return err
		}

		if err := writer.WriteHeader(&tar.Header{
			Name:     entry,
			Size:     e.Size(),
			Mode:     int64(e.Mode()),
			ModTime:  e.ModTime(),
			Typeflag: tar.TypeReg,
		}); err != nil {
			return err
		}

		if _, err := io.Copy(writer, c); err != nil {
			return err
		}

		c.Close()
	}

	return nil
}

func BuildZip(target string, dst io.WriteCloser, l []fs.FileInfo) error {
	writer := zip.NewWriter(dst)
	defer writer.Close()
	for _, e := range l {
		if e.IsDir() {
			continue
		}

		c, err := os.Open(filepath.Join(target, e.Name()))
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
