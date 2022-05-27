package kar

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
)

func openArchive(target string) (*bytes.Buffer, error) {
	if _, err := os.Stat(target); err != nil {
		return nil, err
	}

	f, err := os.Open(target)
	if err != nil {
		return nil, err
	}

	defer f.Close()

	buf := new(bytes.Buffer)
	if _, err := io.Copy(buf, f); err != nil {
		return nil, err
	}

	return buf, nil
}

func closeArchive(c io.Closer) {
	if err := c.Close(); err != nil {
		log.Println("cannot close")
	}
}

func Info(target string) (*Options, error) {
	f, err := openArchive(target)
	if err != nil {
		return nil, err
	}

	a, err := detectArchiveFormat(bytes.NewReader(f.Bytes()[:]), magicNumbers)
	if err != nil {
		return nil, err
	}

	m, err := tarReader(f, gzipReader)
	if err != nil {
		return nil, err
	}

	fmt.Println(m)

	return &Options{Container: *a}, nil
}

var (
	ErrInvalidContainer = errors.New("unknown file format")
)

type mem struct {
	files map[string][]byte
}

func readHead(f io.Reader) ([]byte, error) {
	head := make([]byte, 512)

	if _, err := io.ReadFull(f, head[:]); err != nil {
		if err != io.ErrUnexpectedEOF {
			return nil, ErrInvalidContainer
		}
	}

	return head[:], nil
}

func detectArchiveFormat(f io.Reader, magicNumbers map[archiveFormat][]byte) (*archiveFormat, error) {
	head, err := readHead(f)
	if err != nil {
		return nil, err
	}

	for format, p := range magicNumbers {
		if bytes.HasPrefix(head[:], p) {
			return &format, nil
		}
	}

	return nil, ErrInvalidContainer
}

var magicNumbers = map[archiveFormat][]byte{
	Zip:  []byte("PK\x03\x04"),
	GZip: []byte("\x1F\x8B\x08"),
	ZStd: []byte("\x28\xB5\x2F\xFD"),
	XZ:   []byte("\xFD\x37\x7A\x58\x5A\x00"),
}

type info struct {
	magic  []byte
	offset int
}
