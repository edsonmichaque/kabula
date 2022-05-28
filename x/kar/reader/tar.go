package reader

import (
	"archive/tar"
	"archive/zip"
	"bytes"
	"compress/gzip"
	"io"

	"github.com/klauspost/compress/zstd"
	"github.com/ulikunitz/xz"
)

func ZStd(f io.Reader) (io.Reader, error) {
	var dst bytes.Buffer

	xr, err := zstd.NewReader(f)
	if err != nil {
		return nil, err
	}

	if _, err := io.Copy(&dst, xr); err != nil {
		if err != io.EOF {
			return nil, err
		}
	}

	return &dst, nil
}

func XZ(f io.Reader) (io.Reader, error) {
	var dst bytes.Buffer

	xr, err := xz.NewReader(f)
	if err != nil {
		return nil, err
	}

	if _, err := io.Copy(&dst, xr); err != nil {
		if err != io.EOF {
			return nil, err
		}
	}

	return &dst, nil
}

func GZip(f io.Reader) (io.Reader, error) {
	var dst bytes.Buffer

	gr, err := gzip.NewReader(f)
	if err != nil {
		return nil, err
	}

	defer gr.Close()

	if _, err := io.Copy(&dst, gr); err != nil {
		if err != io.EOF {
			return nil, err
		}
	}

	return &dst, nil
}

type Decompress func(io.Reader) (io.Reader, error)

func Tar(dst io.Reader, fn Decompress) func(r io.Reader) (map[string][]byte, error) {
	return func(r io.Reader) (map[string][]byte, error) {
		tf, err := fn(dst)
		if err != nil {
			return nil, err
		}

		files := make(map[string][]byte)

		tr := tar.NewReader(tf)
		for {
			header, err := tr.Next()
			if err != nil {
				if err != io.EOF {
					return nil, err
				}
				break
			}

			if header.Typeflag != tar.TypeReg {
				continue
			}

			rawBuf := new(bytes.Buffer)
			if _, err := io.Copy(rawBuf, tr); err != nil {
				if err != io.EOF {
					return nil, err
				}
			}

			compressedBuf := new(bytes.Buffer)
			zr, err := zstd.NewWriter(compressedBuf)
			if err != nil {
				return nil, err
			}

			if _, err := io.Copy(zr, rawBuf); err != nil {
				if err != io.EOF {
					return nil, err
				}
			}

			if err := zr.Close(); err != nil {
				return nil, err
			}

			files[header.Name] = compressedBuf.Bytes()
		}

		return files, nil
	}
}

func Zip(dst io.Reader) (map[string][]byte, error) {
	data, err := io.ReadAll(dst)
	if err != nil {
		return nil, err
	}

	zr, err := zip.NewReader(bytes.NewReader(data), 0)
	if err != nil {
		return nil, err
	}

	files := make(map[string][]byte)

	for _, f := range zr.File {
		r, err := f.Open()
		if err != nil {
			return nil, err
		}

		rawBuf := new(bytes.Buffer)
		if _, err := io.Copy(rawBuf, r); err != nil {
			return nil, err
		}

		compressedBuf := new(bytes.Buffer)
		zr, err := zstd.NewWriter(compressedBuf)
		if err != nil {
			return nil, err
		}

		if _, err := io.Copy(zr, rawBuf); err != nil {
			if err != io.EOF {
				return nil, err
			}
		}

		if err := zr.Close(); err != nil {
			return nil, err
		}

		files[f.Name] = compressedBuf.Bytes()
	}

	return files, nil
}
