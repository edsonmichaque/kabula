package kab

import (
	"archive/tar"
	"archive/zip"
	"compress/gzip"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/edsonmichaque/kabula/x/spec"
	"github.com/klauspost/compress/zstd"
)

type Options struct {
	Content   ContentFormat
	Container ContainerFormat
}

var DefaultOptions = Options{
	Content:   FormatJSON,
	Container: Gzip,
}

type Builder func(string, io.WriteCloser, []fs.FileInfo) error

const (
	extXML       = "xml"
	extJSON      = "json"
	manifestPath = "manifest"
)

func Build(target string, args Options) error {
	contentFormat := extJSON
	if args.Content == FormatXML {
		contentFormat = extXML
	}

	fullpath := filepath.Join(target, fmt.Sprintf("%s.%s", manifestPath, contentFormat))
	f, err := os.Open(fullpath)
	if err != nil {
		return err
	}

	defer f.Close()

	data, err := io.ReadAll(f)
	if err != nil {
		return err
	}

	var manifestFile spec.Manifest

	if args.Content == FormatXML {
		if err := xml.Unmarshal(data, &manifestFile); err != nil {
			return err
		}
	}

	if args.Content == FormatJSON {
		if err := json.Unmarshal(data, &manifestFile); err != nil {
			return err
		}
	}

	dstPath := fmt.Sprintf("%s_%s.%s", manifestFile.Name, manifestFile.Version, spec.Ext+string(contentFormat[0]))
	dst, err := os.Create(dstPath)
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

	var build Builder
	if args.Container == Tar {
		build = buildTar
	}

	if args.Container == Gzip {
		build = buildTgz
	}

	if args.Container == Zip {
		build = buildZip
	}

	if args.Container == Zstd {
		build = buildZstd
	}

	if err := build(target, dst, dirEntries); err != nil {
		return err
	}

	return nil
}

func buildTar(target string, dst io.WriteCloser, entries []fs.FileInfo) error {
	tw := tar.NewWriter(dst)
	defer tw.Close()

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		reader, err := os.Open(filepath.Join(target, entry.Name()))
		if err != nil {
			return err
		}

		relPath, err := filepath.Rel(reader.Name(), target)
		if err != nil {
			return err
		}

		if err := tw.WriteHeader(&tar.Header{
			Name:     relPath,
			Size:     entry.Size(),
			Mode:     int64(entry.Mode()),
			ModTime:  entry.ModTime(),
			Typeflag: tar.TypeReg,
		}); err != nil {
			return err
		}

		if _, err := io.Copy(tw, reader); err != nil {
			return err
		}

		reader.Close()
	}

	return nil
}

func buildTgz(target string, dst io.WriteCloser, entries []fs.FileInfo) error {
	gw := gzip.NewWriter(dst)
	defer gw.Close()

	return buildTar(target, gw, entries)
}

func buildZstd(target string, dst io.WriteCloser, entries []fs.FileInfo) error {
	zw, err := zstd.NewWriter(dst)
	if err != nil {
		return err
	}
	defer zw.Close()

	return buildTar(target, zw, entries)
}

func buildZip(target string, dst io.WriteCloser, entries []fs.FileInfo) error {
	writer := zip.NewWriter(dst)
	defer writer.Close()
	for _, e := range entries {
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
