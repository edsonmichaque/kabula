package builder

import (
	"io"

	"github.com/ulikunitz/xz"
)

func TXz(target string, dst io.WriteCloser, entries []Entry) error {
	xw, err := xz.NewWriter(dst)
	if err != nil {
		return err
	}

	defer xw.Close()

	return Tar(target, xw, entries)
}
