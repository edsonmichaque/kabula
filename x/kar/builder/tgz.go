package builder

import (
	"compress/gzip"
	"io"
)

func TGz(target string, dst io.WriteCloser, entries []Entry) error {
	gw := gzip.NewWriter(dst)
	defer gw.Close()

	return Tar(target, gw, entries)
}
