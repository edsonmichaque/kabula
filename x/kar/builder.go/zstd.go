package builder

import (
	"io"

	"github.com/klauspost/compress/zstd"
)

func ZStd(target string, dst io.WriteCloser, entries []Entry) error {
	zw, err := zstd.NewWriter(dst)
	if err != nil {
		return err
	}
	defer zw.Close()

	return Tar(target, zw, entries)
}
